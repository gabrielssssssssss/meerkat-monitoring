package transparency

import (
	"encoding/base64"
	"strings"

	ct "github.com/google/certificate-transparency-go"
	ctTls "github.com/google/certificate-transparency-go/tls"
	ctX509 "github.com/google/certificate-transparency-go/x509"
)

// https://gist.github.com/dvas0004/c6037de1ef6bc66e6d52b3d562ad690c
func ParseLeafInput(leafInput string) (string, error) {
	decodedLeaf, err := base64.StdEncoding.DecodeString(leafInput)
	if err != nil {
		return "", err
	}

	var payload ct.MerkleTreeLeaf
	_, err = ctTls.Unmarshal([]byte(decodedLeaf), &payload)
	if err != nil {
		return "", err
	}

	switch eType := payload.TimestampedEntry.EntryType; eType {
	case 0:
		cert, _ := ctX509.ParseCertificate(payload.TimestampedEntry.X509Entry.Data)
		for _, domain := range cert.DNSNames {
			return strings.Replace(domain, "*.", "", 1), nil
		}

	case 1:
		cert, _ := ctX509.ParseTBSCertificate(payload.TimestampedEntry.PrecertEntry.TBSCertificate)
		for _, domain := range cert.DNSNames {
			return strings.Replace(domain, "*.", "", 1), nil
		}
	default:
		return "", ErrUnknownEntryType
	}
	return "", ErrUnknownEntryType
}
