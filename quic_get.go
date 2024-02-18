//go:build with_quic

package scribe

import (
	"context"
	"crypto/x509"
	"errors"
	"net"
	"time"

	quic "github.com/sagernet/quic-go"
)

func (q *QuicCertGetter) GetCert(timeout time.Duration, conn net.Conn) (cert []*x509.Certificate, err error) {
	if conn == nil {
		conn, err = net.DialTimeout("udp", q.Target, timeout)
		if err == nil {
			return
		}
	}
	udpConn, isUDPConn := conn.(*net.UDPConn)
	if !isUDPConn {
		return nil, errors.New("not UDP conn")
	}
	var packetConn net.PacketConn = udpConn

	addr, err := net.ResolveUDPAddr("udp", q.Target)
	if err != nil {
		return
	}

	tCfg := q.tlsConfig()
	tCfg.NextProtos = []string{"h3"}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if err != nil {
		return
	}

	qCfg := &quic.Config{
		Versions: []quic.VersionNumber{quic.Version2, quic.Version1},
	}

	qConn, err := quic.Dial(ctx, packetConn, addr, tCfg, qCfg)
	if err != nil {
		return
	}
	defer qConn.CloseWithError(0x00, "NO_ERROR")

	cert = qConn.ConnectionState().TLS.PeerCertificates

	return
}
