package scribe

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"net"
	"time"
	
	quic "github.com/quic-go/quic-go"
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

func GetQuicRawCert(target, serverName string) (cert []*x509.Certificate, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err != nil {
		return
	}
	
	tCfg := &tls.Config{
		ServerName:         serverName,
		InsecureSkipVerify: true,
		NextProtos:         []string{"h3"},
	}
	
	qCfg := &quic.Config{
		Versions: []quic.VersionNumber{quic.Version2, quic.Version1},
	}
	
	conn, err := quic.DialAddr(ctx, target, tCfg, qCfg)
	if err != nil {
		return
	}
	defer conn.CloseWithError(0x00, "NO_ERROR")
	
	cert = conn.ConnectionState().TLS.PeerCertificates
	
	return
}
