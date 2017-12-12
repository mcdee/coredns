package plugin

import (
	"strings"

	"github.com/miekg/dns"
)

// eligibleForWildcard sees if a name is eligible for a wildcard.  To be so it
// must have no resource records of any type specifically against its name
func eligibleForWildcard(server Server, domain string, name string) bool {
	if strings.HasPrefix(domain, "*.") {
		// Already a wildcard
		return false
	}
	records, err := server.NumRecords(domain, name)
	if err != nil {
		// TODO now what?
		return false
	}
	return records == 0
}

// replaceWithWildcard replaces the left most label with '*'.
func replaceWithAsteriskLabel(qname string) (wildcard string) {
	i, shot := dns.NextLabel(qname, 0)
	if shot {
		return ""
	}

	return "*." + qname[i:]
}
