package blindfold

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/binary"
	"math/big"
	"testing"
)

// This decoder deliberately does not call the production encoder or its helpers.
func recoverEnvelope(t *testing.T, location string, key *rsa.PrivateKey) ([]byte, error) {
	t.Helper()
	b, err := base64.StdEncoding.DecodeString(location[len("string:///"):])
	if err != nil {
		t.Fatal(err)
	}
	off := 0
	lp := func() []byte {
		n := int(binary.BigEndian.Uint32(b[off:]))
		off += 4
		v := b[off : off+n]
		off += n
		return v
	}
	if string(lp()) != "example-tenant" {
		t.Fatal("tenant")
	}
	if binary.BigEndian.Uint32(b[off:]) != 1 {
		t.Fatal("version")
	}
	off += 4
	pid := binary.BigEndian.Uint64(b[off:])
	off += 8
	if b[off] != 2 {
		t.Fatal("algorithm")
	}
	off++
	e := new(big.Int).SetBytes(lp())
	n := new(big.Int).SetBytes(lp())
	wrapped := new(big.Int).SetBytes(lp())
	factor := new(big.Int).SetUint64(pid)
	factor.Lsh(factor, 1)
	factor.Add(factor, new(big.Int).SetUint64(1<<31|1))
	e.Mul(e, factor)
	phi := new(big.Int).Mul(new(big.Int).Sub(key.Primes[0], big.NewInt(1)), new(big.Int).Sub(key.Primes[1], big.NewInt(1)))
	d := new(big.Int).ModInverse(e, phi)
	if d == nil {
		t.Fatal("noninvertible fixture")
	}
	block := new(big.Int).Exp(wrapped, d, n).FillBytes(make([]byte, (n.BitLen()+7)/8-2))
	if !bytes.Equal(block[:4], []byte{0xde, 0xad, 0xbe, 0xef}) {
		t.Fatal("magic")
	}
	a, err := aes.NewCipher(block[16:48])
	if err != nil {
		t.Fatal(err)
	}
	g, err := cipher.NewGCM(a)
	if err != nil {
		t.Fatal(err)
	}
	return g.Open(nil, block[4:16], b[off:], nil)
}

func TestProtocolIndependentRecovery(t *testing.T) {
	// Generate a group for which both policy boundaries have an inverse.
	for _, pid := range []uint64{101, ^uint64(0)} {
		var key *rsa.PrivateKey
		for {
			var err error
			key, err = rsa.GenerateKey(rand.Reader, 2048)
			if err != nil {
				t.Fatal(err)
			}
			factor := new(big.Int).SetUint64(pid)
			factor.Lsh(factor, 1)
			factor.Add(factor, big.NewInt(1<<31|1))
			phi := new(big.Int).Mul(new(big.Int).Sub(key.Primes[0], big.NewInt(1)), new(big.Int).Sub(key.Primes[1], big.NewInt(1)))
			if new(big.Int).ModInverse(factor, phi) != nil {
				break
			}
		}
		// An exponent wider than uint64 catches accidental machine-integer conversions.
		e := new(big.Int).Lsh(big.NewInt(1), 80)
		e.Add(e, big.NewInt(1))
		phi := new(big.Int).Mul(new(big.Int).Sub(key.Primes[0], big.NewInt(1)), new(big.Int).Sub(key.Primes[1], big.NewInt(1)))
		for new(big.Int).ModInverse(e, phi) == nil {
			e.Add(e, big.NewInt(2))
		}
		c := Context{Tenant: "example-tenant", KeyVersion: 1, PolicyID: pid, Modulus: key.N, Exponent: e}
		secret := bytes.Repeat([]byte("synthetic-secret"), 150)
		a, err := Encrypt(secret, c)
		if err != nil {
			t.Fatal(err)
		}
		b, err := Encrypt(secret, c)
		if err != nil {
			t.Fatal(err)
		}
		if a == b {
			t.Fatal("encryption must be randomized")
		}
		got, err := recoverEnvelope(t, a, key)
		if err != nil || !bytes.Equal(got, secret) {
			t.Fatal("recovery", err)
		}
		raw, _ := base64.StdEncoding.DecodeString(a[10:])
		raw[len(raw)-1] ^= 1
		if _, err := recoverEnvelope(t, "string:///"+base64.StdEncoding.EncodeToString(raw), key); err == nil {
			t.Fatal("tamper accepted")
		}
	}
}

func TestProtocolLimits(t *testing.T) {
	k, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal(err)
	}
	c := Context{Tenant: "example-tenant", KeyVersion: 1, Modulus: k.N, Exponent: big.NewInt(65537)}
	if _, err := Encrypt(make([]byte, MaxInput+1), c); err == nil {
		t.Fatal("input limit")
	}
	if _, err := Encrypt(make([]byte, MaxEncoded), c); err == nil {
		t.Fatal("encoded limit")
	}
	c.Exponent = big.NewInt(2)
	if _, err := Encrypt([]byte("test"), c); err == nil {
		t.Fatal("invalid exponent")
	}
}
