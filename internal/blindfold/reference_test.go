package blindfold

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"testing"
)

func TestPinnedReferenceRecovery(t *testing.T) {
	if err := ValidateContract(); err != nil {
		t.Fatal(err)
	}
	b, _ := contract.ReadFile("contract/blindfold-contract-v1.txt")
	var bundle struct {
		Files map[string]string `json:"files"`
	}
	if err := json.Unmarshal(b, &bundle); err != nil {
		t.Fatal(err)
	}
	var ref struct {
		Location string `json:"location"`
	}
	if err := json.Unmarshal([]byte(bundle.Files["fixtures/pinned-reference.json"]), &ref); err != nil {
		t.Fatal(err)
	}
	fixtures, err := contract.ReadFile("contract/synthetic.json")
	if err != nil {
		t.Fatal(err)
	}
	var files map[string]string
	if err := json.Unmarshal(fixtures, &files); err != nil {
		t.Fatal(err)
	}
	raw, _ := base64.StdEncoding.DecodeString(files["rsa-key.pem"])
	block, _ := pem.Decode(raw)
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		t.Fatal(err)
	}
	got, err := recoverEnvelope(t, ref.Location, key.(*rsa.PrivateKey))
	if err != nil {
		t.Fatal(err)
	}
	want := make([]byte, 2048)
	for i := range want {
		want[i] = byte(i)
	}
	if !bytes.Equal(want, got) {
		t.Fatal("reference recovery mismatch")
	}
}
