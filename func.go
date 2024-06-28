package VAS_util_go

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"net"
)

// Extracts the IP from a string containing both port and IP
func ExtractIP(remote string) string {
	host, _, err := net.SplitHostPort(remote)
	if err != nil {
		return ""
	}
	return host
}

func GeneratePrivatePEM(privateKey *rsa.PrivateKey) []byte {
	keyPem := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	buf := new(bytes.Buffer)

	pem.Encode(buf, keyPem)
	return buf.Bytes()
}

func GeneratePublicPEM(publicKey *rsa.PublicKey) []byte {
	keyPem := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(publicKey),
	}

	buf := new(bytes.Buffer)

	pem.Encode(buf, keyPem)
	return buf.Bytes()
}
