package provider

import (
	"context"
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"math/big"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/blindfold"
	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestBlindfoldInlineIdentityHistory(t *testing.T) {
	rsaKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal(err)
	}
	ecKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	material := func(key any) *blindfold.Material {
		t.Helper()
		signer := key.(crypto.Signer)
		template := &x509.Certificate{SerialNumber: big.NewInt(1), Subject: pkix.Name{CommonName: "inline.example.test"}, NotBefore: time.Now().Add(-time.Hour), NotAfter: time.Now().Add(time.Hour), KeyUsage: x509.KeyUsageDigitalSignature}
		der, e := x509.CreateCertificate(rand.Reader, template, template, signer.Public(), key)
		if e != nil {
			t.Fatal(e)
		}
		private, e := x509.MarshalPKCS8PrivateKey(key)
		if e != nil {
			t.Fatal(e)
		}
		result, e := blindfold.Normalize(blindfold.Input{CertificatePEM: pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: der}), PrivateKey: pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: private})})
		if e != nil {
			t.Fatal(e)
		}
		t.Cleanup(result.Close)
		return result
	}
	alpha, beta := material(rsaKey), material(ecKey)
	public, _ := json.Marshal(map[string]any{"data": map[string]any{"tenant": "example-tenant", "key_version": 1, "modulus_base64": base64.StdEncoding.EncodeToString(rsaKey.N.Bytes()), "public_exponent_base64": "AQAB"}})
	policy := []byte(`{"data":{"tenant":"example-tenant","policy_id":"101"}}`)
	encryptionContext, err := blindfold.ParseContext(public, policy)
	if err != nil {
		t.Fatal(err)
	}
	location := "string:///synthetic-retained-ciphertext"
	provenance, _ := json.Marshal(blindfold.Provenance{Version: 1, Entries: map[string]blindfold.Entry{"alpha": alpha.Entry(encryptionContext.Digest, location)}})
	remote := map[string]any{"metadata": map[string]any{"annotations": map[string]any{blindfold.Annotation: string(provenance)}}, "spec": map[string]any{"certificates": []any{map[string]any{"certificate_url": alpha.CertificateURL, "private_key": map[string]any{"blindfold_secret_info": map[string]any{"location": location}}}}}}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case strings.Contains(r.URL.Path, "get_public_key"):
			_, _ = w.Write(public)
		case strings.Contains(r.URL.Path, "get_policy_document"):
			_, _ = w.Write(policy)
		default:
			_ = json.NewEncoder(w).Encode(remote)
		}
	}))
	defer server.Close()
	nativeTypes := map[string]tftypes.Type{}
	for key := range blindfoldSchema().Attributes {
		nativeTypes[key] = tftypes.String
	}
	nativeType := tftypes.Object{AttributeTypes: nativeTypes}
	nodeType := tftypes.Object{AttributeTypes: map[string]tftypes.Type{"blindfold": nativeType, "certificate_url": tftypes.String}}
	rootType := tftypes.Object{AttributeTypes: map[string]tftypes.Type{"name": tftypes.String, "namespace": tftypes.String, "certificates": tftypes.List{ElementType: nodeType}}}
	value := func(ids []string, materials []*blindfold.Material, previous bool) tftypes.Value {
		nodes := make([]tftypes.Value, len(ids))
		for i, id := range ids {
			inputs := map[string]tftypes.Value{}
			for key := range nativeTypes {
				inputs[key] = tftypes.NewValue(tftypes.String, nil)
			}
			setString(inputs, "id", id)
			setString(inputs, "certificate_pem", string(materials[i].ChainPEM))
			setString(inputs, "private_key_wo", string(materials[i].KeyPEM))
			setString(inputs, "material_version", "1")
			cert := tftypes.NewValue(tftypes.String, nil)
			if previous {
				inputs["private_key_wo"] = tftypes.NewValue(tftypes.String, nil)
				setString(inputs, "algorithm", materials[i].Algorithm)
				setString(inputs, "chain_identity", materials[i].ChainIdentity)
				setString(inputs, "spki_identity", materials[i].SPKIIdentity)
				setString(inputs, "context_digest", encryptionContext.Digest)
				setString(inputs, "encrypted_location", location)
				cert = tftypes.NewValue(tftypes.String, materials[i].CertificateURL)
			}
			nodes[i] = tftypes.NewValue(nodeType, map[string]tftypes.Value{"blindfold": tftypes.NewValue(nativeType, inputs), "certificate_url": cert})
		}
		return tftypes.NewValue(rootType, map[string]tftypes.Value{"name": tftypes.NewValue(tftypes.String, "example"), "namespace": tftypes.NewValue(tftypes.String, "example"), "certificates": tftypes.NewValue(rootType.AttributeTypes["certificates"], nodes)})
	}
	adapter := &blindfoldResource{client: client.NewClient(server.URL, "synthetic-token"), item: "/api/config/namespaces/%s/example/%s"}
	state := value([]string{"alpha"}, []*blindfold.Material{alpha}, true)
	t.Run("new ID inserted before another algorithm", func(t *testing.T) {
		config := value([]string{"beta", "alpha"}, []*blindfold.Material{beta, alpha}, false)
		nodes, e := adapter.nodes(context.Background(), config, config, state, false)
		if e != nil {
			t.Fatal(e)
		}
		defer closeNodes(nodes)
		if len(nodes) != 2 || !nodes[0].rotate || nodes[0].location != "" || nodes[1].rotate || nodes[1].location != location {
			t.Fatal("history crossed stable IDs")
		}
	})
	t.Run("existing ID algorithm change rejected", func(t *testing.T) {
		config := value([]string{"alpha"}, []*blindfold.Material{beta}, false)
		nodes, e := adapter.nodes(context.Background(), config, config, state, false)
		closeNodes(nodes)
		if e == nil || !strings.Contains(e.Error(), "algorithm changes") {
			t.Fatal("existing ID algorithm change accepted", e)
		}
	})
}
