//go:build !with_quic

package scribe

import (
	"crypto/x509"
	"errors"
)

func GetQuicRawCert(_, _ string) (cert []*x509.Certificate, err error) {
	return nil, errors.New("did not include QUIC")
}
