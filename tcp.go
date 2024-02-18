package scribe

import (
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
		conn, err = net.DialTimeout("tcp", t.Target, timeout)
		if err != nil {
			return
		}
	}

	tlsConn := tls.Client(conn, t.tlsConfig())
	defer tlsConn.Close()

	cert = tlsConn.ConnectionState().PeerCertificates

	return
}
