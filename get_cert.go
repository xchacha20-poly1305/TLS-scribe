package scribe

import (
	"crypto/tls"
	"crypto/x509"
	"net"
	"time"
)

func GetRawCert(target, serverName string) (cert []*x509.Certificate, err error) {
	dialer := &net.Dialer{
		Timeout: 5 * time.Second,
	}

	tCfg := &tls.Config{
		// NextProtos:               []string{},
		ServerName:         serverName,
		InsecureSkipVerify: true,
	}

	tlsConn, err := tls.DialWithDialer(dialer, "tcp", target, tCfg)
	if err != nil {
		return
	}
	defer tlsConn.Close()

	cert = tlsConn.ConnectionState().PeerCertificates

	return
}
