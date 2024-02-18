package scribe

import (
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"net"
	"net/url"
	"strings"
	"time"
)

// Execute get certificate directly
func Execute(target, serverName string) (result string, err error) {
	if !strings.Contains(target, "://") {
		// Default to use https
		target = "https://" + target
	}
	u, err := url.Parse(target)
	if err != nil {
		return
	}

	p := u.Port()
	if p == "" {
		p = "443"
	}
	target = net.JoinHostPort(u.Hostname(), p)

	if serverName == "" {
		serverName = u.Hostname()
	}
	// log.Println("dest: ", dest, " sni: ", serverName)

	g, err := New(strings.ToLower(u.Scheme), CertGetterOption{
		Target:     target,
		ServerName: serverName,
	})
	if err != nil {
		return
	}
	cert, err := g.GetCert(DialTimeout, nil)
	if err != nil {
		return
	}

	q := u.Query()
	switch strings.ToLower(q.Get("fmt")) {
	case "sha256", "sha-256", "sha", "openssl", "fingerprint": // openssl fingerprint
		result = fingerprintSHA256(cert[0])
	default: // pem
		pemCerts := make([]string, 0)
		for _, cert := range cert {
			pemCert := &pem.Block{
				Type:  "CERTIFICATE",
				Bytes: cert.Raw,
			}
			pemCerts = append(pemCerts, string(pem.EncodeToMemory(pemCert)))
		}

		result = strings.Join(pemCerts, "")
	}

	return
}

type Scribe interface {
	// GetCert returns target's certificate. If not provide a conn, it will create one by itself.
	GetCert(time.Duration, net.Conn) ([]*x509.Certificate, error)
}

// New returns a new Scribe
func New(protocol string, c CertGetterOption) (Scribe, error) {
	switch protocol {
	case "", "https", "tls":
		return NewTLSCertGetter(c), nil
	case "quic", "h3", "http3":
		return NewQuicCertGetter(c), nil
	default:
		return nil, errors.New("unknown protocol")
	}
}

const DialTimeout = time.Duration(5) * time.Second

type CertGetterOption struct {
	Target     string
	ServerName string
}

func (c *CertGetterOption) tlsConfig() *tls.Config {
	return &tls.Config{
		ServerName:         c.ServerName,
		InsecureSkipVerify: true,
	}
}

func fingerprintSHA256(cert *x509.Certificate) string {
	hash := sha256.Sum256(cert.Raw)
	hashStr := hex.EncodeToString(hash[:])
	hashStr = strings.ToUpper(hashStr)

	hashSli := strings.Split(hashStr, "")

	var result []string
	for i := 0; i < len(hashSli)-1; i += 2 {
		appendSlice := strings.Join(hashSli[i:i+2], "")
		result = append(result, appendSlice)
	}

	return strings.Join(result, ":")
}
