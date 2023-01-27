package crypto

import (
	"crypto"
	"crypto/rand"
	"crypto/tls"
	"errors"

	"github.com/ava-labs/avalanchego/utils/hashing"
)

var errInvalidTLSKey = errors.New("invalid TLS key")

// TLSSigner is signs ips with a TLS key.
type TLSSigner struct {
	privateKey crypto.Signer
}

// NewTLSSigner returns a new instance of TLSSigner.
func NewTLSSigner(cert *tls.Certificate) (TLSSigner, error) {
	privateKey, ok := cert.PrivateKey.(crypto.Signer)
	if !ok {
		return TLSSigner{}, errInvalidTLSKey
	}

	return TLSSigner{
		privateKey: privateKey,
	}, nil
}

func (t TLSSigner) Sign(bytes []byte) ([]byte, error) {
	tlsSig, err := t.privateKey.Sign(rand.Reader,
		hashing.ComputeHash256(bytes), crypto.SHA256)
	if err != nil {
		return nil, err
	}

	return tlsSig, err
}
