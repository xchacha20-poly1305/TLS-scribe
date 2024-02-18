//go:build !with_quic

package scribe

import (
	"crypto/x509"
	"errors"
	"net"
	"time"
)

func (q *QuicCertGetter) GetCert(_ time.Duration, _ net.Conn) ([]*x509.Certificate, error) {
	return nil, errors.New("not include quic")
}
