package table

import (
	"bufio"
	"log"
	"os/exec"

	"github.com/posener/eztables/rule"
)

// Chain is a list of rules
type Chain struct {
	Name  string
	Rules []rule.Rule
}

// Table is a list of chains
type Table []*Chain

func Load(chain string) (*Table, error) {
	args := []string{"-v", "-S"}
	if chain != "" {
		args = append(args, chain)
	}
	cmd := exec.Command("iptables", args...)
	r, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	err = cmd.Start()
	defer cmd.Wait()
	if err != nil {
		return nil, err
	}

	s := bufio.NewScanner(r)
	t := Table{}
	for s.Scan() {
		line := s.Text()
		if len(line) < 2 || line[0] != '-' {
			continue
		}
		switch line[1] {
		case 'A':
			r, err := rule.Parse(line)
			if err != nil {
				log.Printf("failed parsing A line: %s", err)
				continue
			}
			t.addRule(*r)
		}
	}
	return &t, nil
}

func (t *Table) addRule(r rule.Rule) {
	if len(*t) == 0 || (*t)[len(*t)-1].Name != r.Chain {
		*t = append(*t, &Chain{Name: r.Chain})
	}
	ch := (*t)[len(*t)-1]
	ch.Rules = append(ch.Rules, r)
}
