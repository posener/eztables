package rule

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/dustin/go-humanize"
)

const (
	// flags are tuples of -<name> <value>, we don't include -g and -j flags
	flags = `(\s*(?P<not>!\s+)?--?(?P<key>[^(jg|\s)]\S*)\s+(?P<value>\S+))*?`
	// chain name is a mandatory field, of the form -A <chain>
	chain = `-A\s+(?P<chain>\S+)`
	// match are all flags that are not the target
	match = `(\s+(?P<match>` + flags + `))?`
	// count is of the form -c <packages> <bytes>
	count = `(\s+-c\s+(?P<packets>\d+)\s+(?P<bytes>\d+))?`
	// targetArgs are flags after the target name
	targetArgs = `(\s+(?P<targetArgs>` + flags + `))?`
	// target is either -j <target> <targetArgs> or -g <target> <targetArgs>
	target = `(\s+-((?P<jump>j)|(?P<goto>g))\s+(?P<target>\S+)` + targetArgs + `)?`
)

var (
	reARule = regexp.MustCompile(`^` + chain + match + count + target + `$`)
)

// Rule is an iptables rule
type Rule struct {
	// Chain is the chain that the rule belongs to
	Chain string
	// Match is the matching conditions that are defined on a rule
	Match string
	// Jump is the a jump target defined on a rule
	Jump string
	// GoTo is the a goto target defined on a rule
	GoTo string
	// TargetArgs are arguments that will be passed with the rule to the target
	TargetArgs string
	// Count are counters of the rule
	Count Count
}

// Count defines counters of a rule
type Count struct {
	// Packets are the number of packets that matched the rule
	Packets Packets
	// Bytes accumulates the number of bytes that match the rule
	Bytes Bytes
}

// Bytes represents a number in byte units
type Bytes uint64

func (b Bytes) String() string {
	return humanize.Bytes(uint64(b))
}

// Packets represents a number in number of packets
type Packets uint64

func (p Packets) String() string {
	return humanize.Comma(int64(p))
}

// Parse parses an iptables -S line to match a "-A" rule
func Parse(line string) (*Rule, error) {
	matches := reARule.FindStringSubmatch(line)
	if len(matches) < 1 {
		return nil, fmt.Errorf("not a rule: %s", line)
	}
	r := new(Rule)
	result := make(map[string]string)
	for i, name := range reARule.SubexpNames() {
		if i != 0 {
			result[name] = matches[i]
		}
	}
	r.Chain = result["chain"]
	r.Match = result["match"]
	if result["packets"] != "" && result["bytes"] != "" {
		p, _ := strconv.ParseUint(result["packets"], 10, 0)
		b, _ := strconv.ParseUint(result["bytes"], 10, 0)
		r.Count.Packets = Packets(p)
		r.Count.Bytes = Bytes(b)
	}
	if result["jump"] == "j" {
		r.Jump = result["target"]
	}
	if result["goto"] == "g" {
		r.GoTo = result["target"]
	}
	r.TargetArgs = result["targetArgs"]
	return r, nil
}

// Target is the the target that the rule points to
func (r Rule) Target() string {
	if r.Jump != "" {
		return r.Jump
	}
	return r.GoTo
}

// Positive is a rule that passes the packet
func (r Rule) Positive() bool {
	t := r.Target()
	return t == "ACCEPT"
}

// Negative is a rule that does not pass the packet
func (r Rule) Negative() bool {
	t := r.Target()
	return t == "DROP" || t == "REJECT"
}
