package ens

import (
	"strings"

	"golang.org/x/crypto/sha3"
)

// DNSWireFormatDomainHash hashes a domain name in wire format
func DNSWireFormatDomainHash(domain string) (hash [32]byte) {
	sha := sha3.NewLegacyKeccak256()
	sha.Write(DNSWireFormat(domain))
	sha.Sum(hash[:0])
	return
}

// DNSWireFormat turns a domain name in to wire format
func DNSWireFormat(domain string) []byte {
	// Remove leading and trailing dots
	domain = strings.TrimLeft(domain, ".")
	domain = strings.TrimRight(domain, ".")
	domain = strings.ToLower(domain)

	if domain == "" {
		return []byte{0x00}
	}

	bytes := make([]byte, len(domain)+2)
	pieces := strings.Split(domain, ".")
	offset := 0
	for _, piece := range pieces {
		bytes[offset] = byte(len(piece))
		offset++
		copy(bytes[offset:offset+len(piece)], []byte(piece))
		offset += len(piece)
	}
	return bytes
}
