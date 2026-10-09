package blindfold

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"testing"
	"time"

	"github.com/youmark/pkcs8"
	pkcs12 "software.sslmate.com/src/go-pkcs12"
)

func pair(t *testing.T) (*ecdsa.PrivateKey, *x509.Certificate, []byte, []byte) {
	t.Helper()
	k, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	c := &x509.Certificate{SerialNumber: big.NewInt(42), Subject: pkix.Name{CommonName: "blindfold.example.test"}, NotBefore: time.Now().Add(-time.Hour), NotAfter: time.Now().Add(time.Hour), KeyUsage: x509.KeyUsageDigitalSignature}
	b, err := x509.CreateCertificate(rand.Reader, c, c, &k.PublicKey, k)
	if err != nil {
		t.Fatal(err)
	}
	c, err = x509.ParseCertificate(b)
	if err != nil {
		t.Fatal(err)
	}
	der, err := x509.MarshalPKCS8PrivateKey(k)
	if err != nil {
		t.Fatal(err)
	}
	return k, c, pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: b}), pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: der})
}
func TestMaterialNormalization(t *testing.T) {
	k, c, cert, key := pair(t)
	a, err := Normalize(Input{CertificatePEM: cert, PrivateKey: key})
	if err != nil {
		t.Fatal(err)
	}
	defer a.Close()
	sec, err := x509.MarshalECPrivateKey(k)
	if err != nil {
		t.Fatal(err)
	}
	b, err := Normalize(Input{CertificatePEM: append([]byte("\n"), cert...), PrivateKey: pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: sec})})
	if err != nil {
		t.Fatal(err)
	}
	defer b.Close()
	if a.ChainIdentity != b.ChainIdentity || !bytes.Equal(a.KeyPEM, b.KeyPEM) {
		t.Fatal("serialization triggered change")
	}
	protected, err := pkcs8.MarshalPrivateKey(k, []byte("synthetic-password"), nil)
	if err != nil {
		t.Fatal(err)
	}
	p, err := Normalize(Input{CertificatePEM: cert, PrivateKey: pem.EncodeToMemory(&pem.Block{Type: "ENCRYPTED PRIVATE KEY", Bytes: protected}), Passphrase: []byte("synthetic-password")})
	if err != nil {
		t.Fatal(err)
	}
	defer p.Close()
	if !bytes.Equal(a.KeyPEM, p.KeyPEM) {
		t.Fatal("protected normalization")
	}
	bundle, err := pkcs12.Modern.Encode(k, c, nil, "synthetic-password")
	if err != nil {
		t.Fatal(err)
	}
	v, err := Normalize(Input{PKCS12: bundle, Passphrase: []byte("synthetic-password")})
	if err != nil {
		t.Fatal(err)
	}
	defer v.Close()
	if v.ChainIdentity != a.ChainIdentity || !bytes.Equal(v.KeyPEM, a.KeyPEM) {
		t.Fatal("P12 normalization")
	}
	if _, err := Normalize(Input{CertificatePEM: cert, PrivateKey: append(key, key...)}); err == nil {
		t.Fatal("multiple keys accepted")
	}
	if _, err := Normalize(Input{PKCS12: bundle, Passphrase: []byte("wrong")}); err == nil {
		t.Fatal("wrong password accepted")
	}
	_, _, _, wrong := pair(t)
	if _, err := Normalize(Input{CertificatePEM: cert, PrivateKey: wrong}); err == nil {
		t.Fatal("mismatched key accepted")
	}
}
func TestStrictDocuments(t *testing.T) {
	for _, text := range []string{`{"data":{"tenant":"a","tenant":"a"}}`, `{"data":{"keyVersion":1,"key_version":1}}`, "data: &a {tenant: a}\nother: *a", "data: {}\n---\ndata: {}", "data: {key_version: 1.2}", "data: {key_version: 0x12}"} {
		if _, err := ParseDocument([]byte(text)); err == nil {
			t.Fatalf("accepted %s", text)
		}
	}
	v, err := ParseDocument([]byte(`{"data":{"keyVersion":4294967295,"policyId":"18446744073709551615"}}`))
	if err != nil || v == nil {
		t.Fatal(err)
	}
}
