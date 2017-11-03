package rule

import (
	"fmt"
	"html/template"
	"regexp"
	"strconv"
	"strings"

	"github.com/dustin/go-humanize"
)

var reCounts = regexp.MustCompile(`^\[(?P<packets>\d+):(?P<bytes>\d+)]$`)

// Rule is an iptables rule
type Rule struct {
	// Chain is the chain that the rule belongs to
	Chain string
	// Match is the matching conditions that are defined on a rule
	Match []Arg
	// Jump is the a jump target defined on a rule
	Jump string
	// GoTo is the a goto target defined on a rule
	GoTo string
	// TargetArgs are arguments that will be passed with the rule to the target
	TargetArgs []Arg
	// Count are counters of the rule
	Count Count
}

// Arg is rule argument
type Arg struct {
	Key   string
	Value []string
	Not   bool
}

// KeyName is the argument key name
func (a Arg) KeyName() string {
	return strings.TrimLeft(a.Key, "-")
}

func (a Arg) String() string {
	not := ""
	if a.Not {
		not = "!"
	}
	return fmt.Sprintf("%s%s=%s", a.KeyName(), not, strings.Join(a.Value, " "))
}

// ToolTipAttributes are HTML attributes for tool tip
func (a Arg) ToolTipAttributes() template.HTMLAttr {
	t := tooltip[a.KeyName()]
	if t == "" {
		return template.HTMLAttr("")
	}
	return template.HTMLAttr(fmt.Sprintf(`data-toggle="tooltip" data-placement="bottom" title="%s"`, t))
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
	var (
		r   = new(Rule)
		err error
	)

	fields := strings.Fields(line)
	if len(fields) < 2 {
		return nil, fmt.Errorf("not a rule")
	}

	counts, fields := fields[0], fields[1:]
	r.Count, err = parseCounts(counts)
	if err != nil {
		return nil, fmt.Errorf("parsing counts")
	}
	args := parseArguments(fields)
	for _, arg := range args {
		switch arg.Key {
		case "-A":
			if len(arg.Value) == 0 {
				return nil, fmt.Errorf("no value for -A")
			}
			r.Chain = arg.Value[0]
		case "-j":
			if len(arg.Value) == 0 {
				return nil, fmt.Errorf("no value for -j")
			}
			r.Jump = arg.Value[0]
		case "-g":
			if len(arg.Value) == 0 {
				return nil, fmt.Errorf("no value for -g")
			}
			r.GoTo = arg.Value[0]
		default:
			if r.Target() == "" {
				r.Match = append(r.Match, arg)
			} else {
				r.TargetArgs = append(r.TargetArgs, arg)
			}
		}
	}
	if r.Chain == "" {
		return nil, fmt.Errorf("missing chain")
	}
	return r, nil
}

// parseArguments converts a list of fields to a list of arguments
func parseArguments(fields []string) []Arg {
	var (
		args []Arg
		cur  Arg
	)
	for i := 0; i < len(fields); i++ {
		// we meet a new key-value if it is a key (starts with '-') or a negative sign (!)
		if (fields[i][0] == '-' || fields[i][0] == '!') && cur.Key != "" {
			args = append(args, cur)
			cur = Arg{}
		}
		// parse current key-value
		switch {
		case fields[i][0] == '-':
			cur.Key = fields[i]
		case fields[i] == "!":
			cur.Not = true
		default:
			cur.Value = append(cur.Value, fields[i])
		}
	}
	if cur.Key != "" {
		args = append(args, cur)
	}
	return args
}

func parseCounts(s string) (Count, error) {
	match := reCounts.FindStringSubmatch(s)
	if len(match) < 3 {
		return Count{}, fmt.Errorf("regex mismatch")
	}
	p, _ := strconv.ParseUint(match[1], 10, 0)
	b, _ := strconv.ParseUint(match[2], 10, 0)
	return Count{
		Packets: Packets(p),
		Bytes:   Bytes(b),
	}, nil
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
	return t == "DROP" || t == "REJECT" || t == "RETURN"
}

// TargetIsChain return true if it is a link to a chain
func (r Rule) TargetIsChain() bool {
	switch r.Target() {
	case "", "ACCEPT", "DROP", "REJECT", "RETURN", "MASQUERADE":
		return false
	default:
		return true
	}
}
