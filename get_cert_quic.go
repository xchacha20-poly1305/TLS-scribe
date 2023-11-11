//go:build with_quic

package scribe

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"time"

	quic "github.com/sagernet/quic-go"
)

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
