package transparency

import (
	ct "github.com/google/certificate-transparency-go"
	ctTls "github.com/google/certificate-transparency-go/tls"
	ctX509 "github.com/google/certificate-transparency-go/x509"
)

// https://gist.github.com/dvas0004/c6037de1ef6bc66e6d52b3d562ad690c
func ParseLeafInput(leafInput string) (string, error) {
	var payload ct.MerkleTreeLeaf
	ctTls.Unmarshal([]byte(leafInput), &payload)

	switch eType := payload.TimestampedEntry.EntryType; eType {
	case 0:
		cert, _ := ctX509.ParseCertificate(payload.TimestampedEntry.X509Entry.Data)
		for _, domain := range cert.DNSNames {
			return domain, nil
		}

	case 1:
		cert, _ := ctX509.ParseTBSCertificate(payload.TimestampedEntry.PrecertEntry.TBSCertificate)
		for _, domain := range cert.DNSNames {
			return domain, nil
		}
	default:
		return "", ErrUnknowEntryType
	}
	return "", ErrUnknowEntryType
}
