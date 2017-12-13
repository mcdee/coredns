// Package ens implements a plugin that returns information held in the Ethereum Name Service.
package ens

import (
	"fmt"
	"strings"

	"github.com/coredns/coredns/plugin"
	"github.com/coredns/coredns/request"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/wealdtech/ethereal/ens"
	"github.com/wealdtech/ethereal/ens/registrycontract"

	"github.com/miekg/dns"
	"golang.org/x/net/context"
)

// ENS is a plugin that returns information held in the Ethereum Name Service.
type ENS struct {
	Next     plugin.Handler
	Client   *ethclient.Client
	Registry *registrycontract.RegistryContract
}

// IsAuthoritative checks if the ENS plugin is authoritative for a given domain
func (e ENS) IsAuthoritative(domain string) bool {
	// We consider ourselves authoritative if the domain has an SOA record in ENS
	rr, err := e.Query(domain, domain, dns.TypeNS, false)
	return err == nil && len(rr) > 0
}

// HasRecords checks if there are any records for a specific domain and name.
// This is used for wildcard eligibility
func (e ENS) HasRecords(domain string, name string) (bool, error) {
	// Trim trailing '.' if present before hashing
	domain = strings.TrimSuffix(domain, ".")
	domainHash := ens.NameHash(domain)
	nameHash := ens.LabelHash(name)

	resolverAddress, err := ens.Resolver(e.Registry, domain)
	if err != nil {
		return false, err
	}

	resolverContract, err := ens.DnsResolverContractByAddress(e.Client, resolverAddress)
	if err != nil {
		return false, err
	}

	return resolverContract.HasDnsRecords(nil, domainHash, nameHash)
}

// Query queries a given domain/name/resource combination
func (e ENS) Query(domain string, name string, qtype uint16, do bool) ([]dns.RR, error) {
	fmt.Printf("Request of %v for %v/%v\n", qtype, name, domain)
	// Trim trailing '.' if present before hashing
	domain = strings.TrimSuffix(domain, ".")
	domainHash := ens.NameHash(domain)
	nameHash := ens.LabelHash(name)

	resolverAddress, err := ens.Resolver(e.Registry, domain)
	if err != nil {
		return []dns.RR{}, err
	}

	resolverContract, err := ens.DnsResolverContractByAddress(e.Client, resolverAddress)
	if err != nil {
		return []dns.RR{}, err
	}

	data, err := resolverContract.DnsRecord(nil, domainHash, nameHash, qtype)
	if err != nil {
		return []dns.RR{}, err
	}

	results := make([]dns.RR, 0)
	offset := 0
	for offset < len(data) {
		var result dns.RR
		result, offset, err = dns.UnpackRR(data, offset)
		// TODO report err if present?
		if err == nil {
			results = append(results, result)
		}
	}

	return results, err
}

// ServeDNS implements the plugin.Handler interface.
func (e ENS) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	state := request.Request{W: w, Req: r}

	a := new(dns.Msg)
	a.SetReply(r)
	a.Compress = true
	a.Authoritative = true
	var result plugin.Result
	a.Answer, a.Ns, a.Extra, result = plugin.Lookup(e, state)
	switch result {
	case plugin.Success:
	case plugin.NoData:
	case plugin.NameError:
		a.Rcode = dns.RcodeNameError
	case plugin.ServerFailure:
		return dns.RcodeServerFailure, nil
	}

	state.SizeAndDo(a)
	w.WriteMsg(a)

	return 0, nil
}

// Name implements the Handler interface.
func (e ENS) Name() string { return "ens" }
