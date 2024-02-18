package scribe

import (
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"net"
	"net/url"
	"strings"
)

func Execute(target, serverName string) (result string, err error) {
	var cert []*x509.Certificate
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
	dest := net.JoinHostPort(u.Hostname(), p)

	if serverName == "" {
		serverName = u.Hostname()
	}
	// log.Println("dest: ", dest, " sni: ", serverName)

	q := u.Query()

	switch strings.ToLower(u.Scheme) {
	case "", "tls", "https":
		// log.Println("https")
		cert, err = GetRawCert(dest, serverName)
		if err != nil {
			return
		}

	case "quic", "h3", "http3":
		// log.Println("quic")
		cert, err = GetQuicRawCert(dest, serverName)
		if err != nil {
			return
		}

	default:
		return "", errors.New("Unkonw protocol: " + u.Scheme)
	}

	// 导出结果
	switch strings.ToLower(q.Get("fmt")) { // 获取查询字符串中指定的格式
	case "sha256", "sha-256", "sha", "openssl", "fingerprint": // openssl fingerprint
		// cert = fingerprintSHA256(c[len(c)-1])
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
