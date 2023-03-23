package auth

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"strings"

	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
)

// HasRole checks if the client identity has the specified role
func HasRole(ctx cid.ClientIdentity, role string) bool {
	mspid, err := ctx.GetMSPID()
	if err != nil {
		return false
	}
	attrs, err := ctx.GetAttributes()
	if err != nil {
		return false
	}
	for _, attr := range attrs {
		if strings.HasPrefix(attr.Name, fmt.Sprintf("hf.%s.role", mspid)) && attr.Value == role {
			return true
		}
	}
	return false
}

// GetCertAttribute retrieves the value of the specified attribute from the client certificate
func GetCertAttribute(cert string, attribute string) (string, error) {
	block, _ := pem.Decode([]byte(cert))
	if block == nil {
		return "", errors.New("failed to parse certificate PEM")
	}
	x509Cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return "", err
	}
	for _, attr := range x509Cert.Subject.pkixAttributeTypeAndValue {
		if attr.Type.String() == attribute {
			return attr.Value.String(), nil
		}
	}
	return "", errors.New("certificate attribute not found")
}
