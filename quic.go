package scribe

type QuicCertGetter struct {
	CertGetterOption
}

var _ Scribe = (*QuicCertGetter)(nil)

func NewQuicCertGetter(c CertGetterOption) *QuicCertGetter {
	return &QuicCertGetter{
		CertGetterOption: c,
	}
}
