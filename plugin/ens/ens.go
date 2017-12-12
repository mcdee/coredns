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

func (e ENS) IsAuthoritative(domain string) bool {
	// We consider ourselves authoritative if the domain has an SOA record in ENS
	rr, err := e.Query(domain, domain, dns.TypeNS, false)
	return err == nil && len(rr) > 0
}

func (e ENS) NumRecords(domain string, name string) (uint16, error) {
	// Trim trailing '.' if present before hashing
	domain = strings.TrimSuffix(domain, ".")
	domainHash := ens.NameHash(domain)
	nameHash := ens.LabelHash(name)

	resolverAddress, err := ens.Resolver(e.Registry, domain)
	if err != nil {
		return 0, err
	}

	resolverContract, err := ens.DnsResolverContractByAddress(e.Client, resolverAddress)
	if err != nil {
		return 0, err
	}

	return resolverContract.NameEntries(nil, domainHash, nameHash)
}

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

// ServeDNS implements the plugin.Handler interface.
func (e ENS) OldServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	state := request.Request{W: w, Req: r}

	a := new(dns.Msg)
	a.SetReply(r)
	a.Compress = true
	a.Authoritative = true

	// Obtain resolver for the domain
	data, err := query(e, state.Name(), state.QType(), ".")
	if err != nil || len(data) == 0 {
		breakpoint := strings.Index(state.Name(), ".")
		if breakpoint != -1 {
			// Break out in to key/value
			dnsKey := state.Name()[0:breakpoint]
			dnsDomain := state.Name()[breakpoint+1:]
			data, err = query(e, dnsDomain, state.QType(), dnsKey)
		}
	}
	if err != nil {
		// Didn't find any records; pass
		return plugin.NextOrFailure(e.Name(), e.Next, ctx, w, r)
	}

	a.Answer = make([]dns.RR, 0)
	offset := 0
	var result dns.RR
	for offset < len(data) {
		result, offset, err = dns.UnpackRR(data, offset)
		// TODO report err if present?
		if err == nil {
			a.Answer = append(a.Answer, result)
		}
	}

	state.SizeAndDo(a)
	fmt.Println("Response is", a)
	w.WriteMsg(a)

	return 0, nil
}

func query(e ENS, domain string, qtype uint16, key string) ([]byte, error) {
	// Trim trailing '.' if present before hashing
	if strings.HasSuffix(domain, ".") {
		domain = domain[0 : len(domain)-1]
	}
	node := ens.NameHash(domain)

	resolverAddress, err := ens.Resolver(e.Registry, domain)
	if err != nil {
		return []byte{}, err
	}

	resolverContract, err := ens.DnsResolverContractByAddress(e.Client, resolverAddress)
	if err != nil {
		return []byte{}, err
	}

	data, err := resolverContract.DnsRecord(nil, node, ens.LabelHash(key), qtype)
	return data, err
}

// Name implements the Handler interface.
func (e ENS) Name() string { return "ens" }
