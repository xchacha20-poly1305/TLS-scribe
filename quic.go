package scribe

// QuicCertGetter used to get QUIC certificate.
type QuicCertGetter struct {
	CertGetterOption
}

var _ Scribe = (*QuicCertGetter)(nil)

func NewQuicCertGetter(c CertGetterOption) *QuicCertGetter {
	return &QuicCertGetter{
		CertGetterOption: c,
	}
}
