package scribe

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"net"
	"time"
)

// TLSCertGetter used to get tcp-TLS certificate.
type TLSCertGetter struct {
	CertGetterOption
}

var _ Scribe = (*TLSCertGetter)(nil)

func NewTLSCertGetter(c CertGetterOption) *TLSCertGetter {
	return &TLSCertGetter{
		CertGetterOption: c,
	}
}

func (t *TLSCertGetter) GetCert(timeout time.Duration, conn net.Conn) (cert []*x509.Certificate, err error) {
	if conn == nil {
		conn, err = net.DialTimeout("tcp", net.JoinHostPort(t.Target.String(), t.Port), timeout)
		if err != nil {
			return
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if err != nil {
		return
	}

	tlsConn := tls.Client(conn, t.tlsConfig())
	err = tlsConn.HandshakeContext(ctx)
	if err != nil {
		return
	}
	defer tlsConn.Close()

	cert = tlsConn.ConnectionState().PeerCertificates

	return
}
