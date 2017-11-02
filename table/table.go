package table

import (
	"bufio"
	"log"
	"os/exec"

	"fmt"
	"io/ioutil"

	"github.com/posener/eztables/rule"
)

// Chain is a list of rules
type Chain struct {
	Name  string
	Rules []rule.Rule
}

// Table is a list of chains
type Table struct {
	Name   string
	Chains []*Chain
}

// Load loads defined tables from the iptables command
func Load() ([]Table, error) {
	cmd := exec.Command("iptables-save", "-c")
	r, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	defer r.Close()
	errReader, err := cmd.StderrPipe()
	if err != nil {
		return nil, err
	}
	defer errReader.Close()
	err = cmd.Start()
	if err != nil {
		return nil, err
	}

	s := bufio.NewScanner(r)
	t := []Table{}
	for s.Scan() {
		line := s.Text()
		if len(line) == 0 {
			continue
		}
		switch line[0] {
		case '*':
			t = append(t, Table{Name: line[1:]})
		case '[':
			r, err := rule.Parse(line)
			if err != nil {
				log.Printf("Failed parsing line %s: %s", line, err)
				continue
			}
			t[len(t)-1].addRule(*r)
		}
	}
	cmd.Wait()
	if !cmd.ProcessState.Success() {
		errString, _ := ioutil.ReadAll(errReader)
		return nil, fmt.Errorf("command iptables-save failed: %s", string(errString))
	}
	return t, nil
}

// Test tests if iptable command can be run
func Test() error {
	_, err := Load()
	return err
}

func (t *Table) addRule(r rule.Rule) {
	if len(t.Chains) == 0 || t.Chains[len(t.Chains)-1].Name != r.Chain {
		t.Chains = append(t.Chains, &Chain{Name: r.Chain})
	}
	ch := t.Chains[len(t.Chains)-1]
	ch.Rules = append(ch.Rules, r)
}
