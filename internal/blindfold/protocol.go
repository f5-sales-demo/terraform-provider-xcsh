// Package blindfold implements the native F5 binary envelope. It never invokes an executable.
package blindfold

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"math/big"
	"regexp"
	"strconv"
	"unicode/utf8"

	"gopkg.in/yaml.v3"
)

const MaxInput = 2 * 1024 * 1024
const MaxEncoded = 131072
const Annotation = "f5-sales-demo.com/blindfold"

var label = regexp.MustCompile(`^[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?$`)
var decimal = regexp.MustCompile(`^(0|[1-9][0-9]*)$`)
var aliases = map[string]string{"keyVersion": "key_version", "modulusBase64": "modulus_base64", "publicExponentBase64": "public_exponent_base64", "policyId": "policy_id", "policyInfo": "policy_info", "clientName": "client_name", "clientNameMatcher": "client_name_matcher", "clientSelector": "client_selector", "exactValues": "exact_values", "regexValues": "regex_values"}

type Context struct {
	Tenant            string
	KeyVersion        uint32
	PolicyID          uint64
	Modulus, Exponent *big.Int
	Digest            string
}

// ParseDocument rejects duplicate keys (including equal aliases), aliases, multiple documents,
// non-string keys and floating point values. Integer tokens retain their exact decimal value.
func ParseDocument(b []byte) (map[string]any, error) {
	if len(b) > MaxInput || !utf8.Valid(b) {
		return nil, errors.New("invalid Blindfold document size or encoding")
	}
	var root yaml.Node
	d := yaml.NewDecoder(bytes.NewReader(b))
	if err := d.Decode(&root); err != nil {
		return nil, errors.New("malformed Blindfold document")
	}
	var extra yaml.Node
	if err := d.Decode(&extra); err != io.EOF {
		return nil, errors.New("multiple Blindfold documents")
	}
	if len(root.Content) != 1 {
		return nil, errors.New("empty Blindfold document")
	}
	v, err := documentNode(root.Content[0])
	if err != nil {
		return nil, err
	}
	m, ok := v.(map[string]any)
	if !ok {
		return nil, errors.New("blindfold document must be an object")
	}
	return m, nil
}
func documentNode(n *yaml.Node) (any, error) {
	switch n.Kind {
	case yaml.MappingNode:
		m := map[string]any{}
		for i := 0; i < len(n.Content); i += 2 {
			k := n.Content[i]
			if k.Tag != "!!str" {
				return nil, errors.New("non-string Blindfold field")
			}
			name := k.Value
			if a, ok := aliases[name]; ok {
				name = a
			}
			if _, ok := m[name]; ok {
				return nil, errors.New("duplicate or conflicting Blindfold field")
			}
			v, err := documentNode(n.Content[i+1])
			if err != nil {
				return nil, err
			}
			m[name] = v
		}
		return m, nil
	case yaml.SequenceNode:
		v := make([]any, len(n.Content))
		for i, c := range n.Content {
			item, err := documentNode(c)
			if err != nil {
				return nil, err
			}
			v[i] = item
		}
		return v, nil
	case yaml.ScalarNode:
		switch n.Tag {
		case "!!str":
			return n.Value, nil
		case "!!int":
			if !decimal.MatchString(n.Value) {
				return nil, errors.New("Blindfold integer must be unsigned decimal")
			}
			return json.Number(n.Value), nil
		case "!!bool":
			return n.Value == "true", nil
		case "!!null":
			return nil, nil
		}
	}
	return nil, errors.New("unsupported Blindfold document value or alias")
}
func ParseContext(public, policy []byte) (Context, error) {
	p, err := ParseDocument(public)
	if err != nil {
		return Context{}, err
	}
	q, err := ParseDocument(policy)
	if err != nil {
		return Context{}, err
	}
	pub, ok := p["data"].(map[string]any)
	if !ok {
		return Context{}, errors.New("missing public key data")
	}
	pol, ok := q["data"].(map[string]any)
	if !ok {
		return Context{}, errors.New("missing policy data")
	}
	for _, k := range []string{"tenant", "key_version", "modulus_base64", "public_exponent_base64", "policy_id"} {
		if _, ok := p[k]; ok {
			return Context{}, errors.New("ambiguous public key")
		}
		if _, ok := q[k]; ok {
			return Context{}, errors.New("ambiguous policy")
		}
	}
	tenant, ok := pub["tenant"].(string)
	if !ok || !label.MatchString(tenant) || pol["tenant"] != tenant {
		return Context{}, errors.New("Blindfold material tenant mismatch")
	}
	v, ok := pub["key_version"].(json.Number)
	if !ok {
		return Context{}, errors.New("invalid key version")
	}
	version, err := strconv.ParseUint(string(v), 10, 32)
	if err != nil || version == 0 {
		return Context{}, errors.New("invalid key version")
	}
	id, ok := pol["policy_id"].(string)
	if !ok || !decimal.MatchString(id) {
		return Context{}, errors.New("invalid policy ID")
	}
	pid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return Context{}, errors.New("invalid policy ID")
	}
	integer := func(v any) (*big.Int, error) {
		s, ok := v.(string)
		if !ok {
			return nil, errors.New("invalid public integer")
		}
		b, err := base64.StdEncoding.Strict().DecodeString(s)
		if err != nil || len(b) == 0 || base64.StdEncoding.EncodeToString(b) != s || b[0] == 0 {
			return nil, errors.New("invalid public integer")
		}
		return new(big.Int).SetBytes(b), nil
	}
	n, err := integer(pub["modulus_base64"])
	if err != nil {
		return Context{}, err
	}
	e, err := integer(pub["public_exponent_base64"])
	if err != nil {
		return Context{}, err
	}
	canonical, err := CanonicalJSON([]any{p, q})
	if err != nil {
		return Context{}, errors.New("invalid context")
	}
	c := Context{Tenant: tenant, KeyVersion: uint32(version), PolicyID: pid, Modulus: n, Exponent: e, Digest: Hash(canonical)}
	return c, c.Validate()
}
func (c Context) Validate() error {
	if !label.MatchString(c.Tenant) || c.KeyVersion == 0 || c.Modulus == nil || c.Exponent == nil || c.Modulus.BitLen() < 2048 || c.Modulus.BitLen() > 8192 || c.Modulus.Bit(0) != 1 || c.Exponent.BitLen() < 2 || c.Exponent.Bit(0) != 1 || c.Exponent.Cmp(c.Modulus) >= 0 {
		return errors.New("invalid Blindfold tenant RSA material")
	}
	return nil
}
func Hash(b []byte) string { s := sha256.Sum256(b); return hex.EncodeToString(s[:]) }
func Encrypt(secret []byte, c Context) (string, error) {
	if err := c.Validate(); err != nil {
		return "", err
	}
	if len(secret) > MaxInput {
		return "", errors.New("Blindfold input exceeds 2 MiB")
	}
	width := (c.Modulus.BitLen() + 7) / 8
	// Check the final encoding before allocating or encrypting a large secret.
	size := 4 + len(c.Tenant) + 4 + 8 + 1 + 4 + len(c.Exponent.Bytes()) + 4 + width + 4 + width + len(secret) + 16
	if 10+base64.StdEncoding.EncodedLen(size) > MaxEncoded {
		return "", errors.New("encoded Blindfold output exceeds size limit")
	}
	block := make([]byte, width-2)
	defer clear(block)
	if _, err := rand.Read(block); err != nil {
		return "", errors.New("random generation failed")
	}
	copy(block, []byte{0xde, 0xad, 0xbe, 0xef})
	a, err := aes.NewCipher(block[16:48])
	if err != nil {
		return "", errors.New("encryption failed")
	}
	g, err := cipher.NewGCM(a)
	if err != nil {
		return "", errors.New("encryption failed")
	}
	encrypted := g.Seal(nil, block[4:16], secret, nil)
	factor := new(big.Int).SetUint64(c.PolicyID)
	factor.Lsh(factor, 1)
	factor.Add(factor, big.NewInt(1<<31|1))
	effective := new(big.Int).Mul(c.Exponent, factor)
	m := new(big.Int).SetBytes(block)
	defer func() { clear(m.Bits()); m.SetInt64(0) }()
	wrapped := new(big.Int).Exp(m, effective, c.Modulus).FillBytes(make([]byte, width))
	out := bytes.NewBuffer(make([]byte, 0, size))
	lp := func(b []byte) { _ = binary.Write(out, binary.BigEndian, uint32(len(b))); out.Write(b) }
	lp([]byte(c.Tenant))
	_ = binary.Write(out, binary.BigEndian, c.KeyVersion)
	_ = binary.Write(out, binary.BigEndian, c.PolicyID)
	out.WriteByte(2)
	lp(c.Exponent.Bytes())
	lp(c.Modulus.Bytes())
	lp(wrapped)
	out.Write(encrypted)
	return "string:///" + base64.StdEncoding.EncodeToString(out.Bytes()), nil
}

// CanonicalJSON uses sorted Go map keys and does not HTML-escape policy strings.
func CanonicalJSON(value any) ([]byte, error) {
	var b bytes.Buffer
	e := json.NewEncoder(&b)
	e.SetEscapeHTML(false)
	if err := e.Encode(value); err != nil {
		return nil, err
	}
	return bytes.TrimSuffix(b.Bytes(), []byte("\n")), nil
}
