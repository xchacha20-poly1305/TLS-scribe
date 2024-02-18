package scribe

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"net"
	"time"

	quic "github.com/sagernet/quic-go"
)

// TODO: use custom conn
func (q *QuicCertGetter) GetCert(timeout time.Duration, _ net.Conn) (cert []*x509.Certificate, err error) {
	hostPort := net.JoinHostPort(q.Target.String(), q.Port)

	/*
		addr, err := net.ResolveUDPAddr("udp", hostPort)
		if err != nil {
			return
		}

		if conn == nil {
			conn, err = net.DialUDP("udp", nil, addr)
			if err != nil {
				return
			}
		}
		udpConn, isUDPConn := conn.(*net.UDPConn)
		if !isUDPConn {
			return nil, errors.New("not UDP conn")
		}
	*/

	tCfg := q.tlsConfig()

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if err != nil {
		return
	}

	qCfg := &quic.Config{
		Versions: []quic.VersionNumber{quic.Version2, quic.Version1},
	}

	qConn, err := quic.DialAddr(ctx, hostPort, tCfg, qCfg)
	if err != nil {
		return
	}
	defer qConn.CloseWithError(0x00, "NO_ERROR")

	cert = qConn.ConnectionState().TLS.PeerCertificates

	return
}

func (q *QuicCertGetter) tlsConfig() *tls.Config {
	c := q.CertGetterOption.tlsConfig()
	c.NextProtos = []string{"h3"}
	return c
}
