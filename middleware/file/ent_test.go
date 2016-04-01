package file

import (
	"sort"
	"strings"
	"testing"

	"github.com/miekg/coredns/middleware"
	coretest "github.com/miekg/coredns/middleware/testing"

	"github.com/miekg/dns"
	"golang.org/x/net/context"
)

var entTestCases = []coretest.Case{
	{
		Qname: "b.c.miek.nl.", Qtype: dns.TypeA,
		Ns: []dns.RR{
			coretest.SOA("miek.nl.	1800	IN	SOA	linode.atoom.net. miek.miek.nl. 1282630057 14400 3600 604800 14400"),
		},
	},
}

func TestLookupENT(t *testing.T) {
	zone, err := Parse(strings.NewReader(dbMiekENTNL), testzone, "stdin")
	if err != nil {
		t.Fatalf("expect no error when reading zone, got %q", err)
	}

	fm := File{Next: coretest.ErrorHandler(), Zones: Zones{Z: map[string]*Zone{testzone: zone}, Names: []string{testzone}}}
	ctx := context.TODO()

	for _, tc := range entTestCases {
		m := tc.Msg()

		rec := middleware.NewResponseRecorder(&middleware.TestResponseWriter{})
		_, err := fm.ServeDNS(ctx, rec, m)
		if err != nil {
			t.Errorf("expected no error, got %v\n", err)
			return
		}
		resp := rec.Msg()

		sort.Sort(coretest.RRSet(resp.Answer))
		sort.Sort(coretest.RRSet(resp.Ns))
		sort.Sort(coretest.RRSet(resp.Extra))

		if !coretest.Header(t, tc, resp) {
			t.Logf("%v\n", resp)
			continue
		}

		if !coretest.Section(t, tc, coretest.Answer, resp.Answer) {
			t.Logf("%v\n", resp)
		}
		if !coretest.Section(t, tc, coretest.Ns, resp.Ns) {
			t.Logf("%v\n", resp)

		}
		if !coretest.Section(t, tc, coretest.Extra, resp.Extra) {
			t.Logf("%v\n", resp)
		}
	}
}

const dbMiekENTNL = `
$TTL    30M
$ORIGIN miek.nl.
@       IN      SOA     linode.atoom.net. miek.miek.nl. (
                             1282630057 ; Serial
                             4H         ; Refresh
                             1H         ; Retry
                             7D         ; Expire
                             4H )       ; Negative Cache TTL
                IN      NS      linode.atoom.net.

a.b.c           IN      A       139.162.196.78
d.e.f.g         IN      A       139.162.196.78`
