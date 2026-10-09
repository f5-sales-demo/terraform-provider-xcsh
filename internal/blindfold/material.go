package blindfold

import (
	"bytes"
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"io"
	"os"
	"time"

	"github.com/youmark/pkcs8"
	pkcs12 "software.sslmate.com/src/go-pkcs12"
)

type Input struct{ CertificatePEM, PrivateKey, PKCS12, Passphrase []byte }
type Material struct {
	KeyPEM, ChainPEM                                                               []byte
	CertificateURL, Fingerprint, ChainIdentity, SPKIIdentity, ExpiresAt, Algorithm string
}

func (m *Material) Close() { clear(m.KeyPEM); m.KeyPEM = nil }
func ReadFile(name string) ([]byte, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, errors.New("cannot read Blindfold input")
	}
	defer func() { _ = f.Close() }()
	s, err := f.Stat()
	if err != nil || !s.Mode().IsRegular() || s.Size() > MaxInput {
		return nil, errors.New("blindfold input must be a file no larger than 2 MiB")
	}
	b, err := io.ReadAll(io.LimitReader(f, MaxInput+1))
	if err != nil || len(b) > MaxInput {
		clear(b)
		return nil, errors.New("cannot read bounded Blindfold input")
	}
	return b, nil
}
func parseKey(b, password []byte) (crypto.Signer, error) {
	b = bytes.TrimSpace(b)
	if !bytes.HasPrefix(b, []byte("-----BEGIN ")) {
		return nil, errors.New("invalid private key PEM prefix")
	}
	block, rest := pem.Decode(b)
	if block == nil || len(bytes.TrimSpace(rest)) != 0 {
		return nil, errors.New("PEM input must contain exactly one private key")
	}
	der := block.Bytes
	defer clear(der)
	var key any
	var err error
	switch block.Type {
	case "ENCRYPTED PRIVATE KEY":
		key, err = pkcs8.ParsePKCS8PrivateKey(der, password)
	case "PRIVATE KEY":
		key, err = x509.ParsePKCS8PrivateKey(der)
	case "RSA PRIVATE KEY", "EC PRIVATE KEY":
		//nolint:staticcheck // Parse legacy protected PEM locally for xcsh input compatibility; never encrypt with it.
		if x509.IsEncryptedPEMBlock(block) {
			//nolint:staticcheck // Decrypt-only compatibility for existing protected PEM files.
			der, err = x509.DecryptPEMBlock(block, password)
			if err != nil {
				return nil, errors.New("invalid PEM key or passphrase")
			}
			defer clear(der)
		}
		if block.Type == "RSA PRIVATE KEY" {
			key, err = x509.ParsePKCS1PrivateKey(der)
		} else {
			key, err = x509.ParseECPrivateKey(der)
		}
	default:
		return nil, errors.New("unsupported private key PEM block")
	}
	if err != nil {
		return nil, errors.New("invalid PEM key or passphrase")
	}
	s, ok := key.(crypto.Signer)
	if !ok {
		return nil, errors.New("unsupported private key")
	}
	return s, nil
}
func parseChain(b []byte) ([]*x509.Certificate, error) {
	var out []*x509.Certificate
	for len(bytes.TrimSpace(b)) > 0 {
		b = bytes.TrimSpace(b)
		if !bytes.HasPrefix(b, []byte("-----BEGIN CERTIFICATE-----")) {
			return nil, errors.New("invalid certificate PEM chain")
		}
		block, rest := pem.Decode(b)
		if block == nil || block.Type != "CERTIFICATE" || len(block.Headers) > 0 {
			return nil, errors.New("invalid certificate PEM chain")
		}
		c, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			return nil, errors.New("invalid certificate")
		}
		out = append(out, c)
		b = rest
	}
	if len(out) == 0 {
		return nil, errors.New("empty certificate chain")
	}
	return out, nil
}
func Normalize(in Input) (*Material, error) {
	for _, b := range [][]byte{in.CertificatePEM, in.PrivateKey, in.PKCS12, in.Passphrase} {
		if len(b) > MaxInput {
			return nil, errors.New("blindfold input exceeds 2 MiB")
		}
	}
	var key crypto.Signer
	var certs []*x509.Certificate
	if len(in.PKCS12) > 0 {
		if len(in.CertificatePEM) > 0 || len(in.PrivateKey) > 0 {
			return nil, errors.New("bundle conflicts with PEM inputs")
		}
		if bytes.Contains(in.Passphrase, []byte{0}) {
			return nil, errors.New("invalid PKCS12 passphrase")
		}
		k, leaf, chain, err := pkcs12.DecodeChain(in.PKCS12, string(in.Passphrase))
		if err != nil {
			return nil, errors.New("invalid PKCS12 bundle or passphrase; require exactly one key")
		}
		var ok bool
		key, ok = k.(crypto.Signer)
		if !ok || leaf == nil {
			return nil, errors.New("PKCS12 bundle requires a certificate and one supported key")
		}
		certs = append([]*x509.Certificate{leaf}, chain...)
	} else {
		var err error
		key, err = parseKey(in.PrivateKey, in.Passphrase)
		if err != nil {
			return nil, err
		}
		certs, err = parseChain(in.CertificatePEM)
		if err != nil {
			return nil, err
		}
	}
	algo := ""
	switch k := key.(type) {
	case *rsa.PrivateKey:
		if k.N.BitLen() < 2048 || k.N.BitLen() > 8192 || k.Validate() != nil {
			return nil, errors.New("invalid RSA private key; require 2048-8192 bits")
		}
		algo = "RSA"
	case *ecdsa.PrivateKey:
		if (k.Curve != elliptic.P256() && k.Curve != elliptic.P384()) || k.D.Sign() <= 0 || k.D.Cmp(k.Params().N) >= 0 {
			return nil, errors.New("supported EC curves are P-256 and P-384")
		}
		//nolint:staticcheck // Validate ECDSA private scalar against its supplied public point; this is not ECDH.
		x, y := k.Curve.ScalarBaseMult(k.D.Bytes())
		if x.Cmp(k.X) != 0 || y.Cmp(k.Y) != 0 {
			return nil, errors.New("invalid EC private key")
		}
		algo = "EC"
	default:
		return nil, errors.New("unsupported private key")
	}
	spki, err := x509.MarshalPKIXPublicKey(key.Public())
	if err != nil {
		return nil, errors.New("invalid public key")
	}
	// Find the matching leaf and sort the complete supplied chain by signatures.
	leaf := -1
	seen := map[string]bool{}
	now := time.Now()
	for i, c := range certs {
		if seen[Hash(c.Raw)] {
			return nil, errors.New("duplicate certificate")
		}
		seen[Hash(c.Raw)] = true
		if now.Before(c.NotBefore) || !now.Before(c.NotAfter) {
			return nil, errors.New("certificate chain is expired or not yet valid")
		}
		if bytes.Equal(c.RawSubjectPublicKeyInfo, spki) {
			if leaf >= 0 {
				return nil, errors.New("ambiguous matching certificates")
			}
			leaf = i
		}
	}
	if leaf < 0 {
		return nil, errors.New("certificate does not match private key")
	}
	certs[0], certs[leaf] = certs[leaf], certs[0]
	for i := 0; i < len(certs)-1; i++ {
		found := -1
		for j := i + 1; j < len(certs); j++ {
			if bytes.Equal(certs[i].RawIssuer, certs[j].RawSubject) && certs[j].CheckSignature(certs[i].SignatureAlgorithm, certs[i].RawTBSCertificate, certs[i].Signature) == nil {
				if found >= 0 {
					return nil, errors.New("ambiguous certificate issuer")
				}
				found = j
			}
		}
		if found < 0 {
			return nil, errors.New("invalid certificate chain signature or unrelated certificate")
		}
		certs[i+1], certs[found] = certs[found], certs[i+1]
	}
	last := certs[len(certs)-1]
	if bytes.Equal(last.RawIssuer, last.RawSubject) && last.CheckSignature(last.SignatureAlgorithm, last.RawTBSCertificate, last.Signature) != nil {
		return nil, errors.New("invalid root signature")
	}
	der, err := x509.MarshalPKCS8PrivateKey(key)
	if err != nil {
		return nil, errors.New("cannot normalize private key")
	}
	defer clear(der)
	m := &Material{KeyPEM: pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: der}), Fingerprint: Hash(certs[0].Raw), SPKIIdentity: Hash(spki), ExpiresAt: certs[0].NotAfter.UTC().Format(time.RFC3339), Algorithm: algo}
	for _, c := range certs {
		m.ChainPEM = append(m.ChainPEM, pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: c.Raw})...)
	}
	m.ChainIdentity = Hash(m.ChainPEM)
	m.CertificateURL = "string:///" + base64.StdEncoding.EncodeToString(m.ChainPEM)
	if len(m.CertificateURL) > MaxEncoded || len(m.KeyPEM) > MaxInput {
		m.Close()
		return nil, errors.New("encoded certificate exceeds size limit")
	}
	return m, nil
}

// Entry is reconciliation metadata, not attestation. It contains only public identities.
type Entry struct {
	Chain      string `json:"chain"`
	SPKI       string `json:"spki"`
	Context    string `json:"context"`
	Ciphertext string `json:"ciphertext"`
	Algorithm  string `json:"algorithm"`
}
type Provenance struct {
	Version int              `json:"version"`
	Entries map[string]Entry `json:"entries"`
}

func (m *Material) Entry(context, location string) Entry {
	return Entry{Chain: m.ChainIdentity, SPKI: m.SPKIIdentity, Context: context, Ciphertext: Hash([]byte(location)), Algorithm: m.Algorithm}
}
func Matches(entry Entry, m *Material, context, certURL, location string) bool {
	return entry == m.Entry(context, location) && certURL == m.CertificateURL && len(location) > 10 && len(location) <= MaxEncoded && len(entry.Ciphertext) == 64
}
