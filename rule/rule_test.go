package rule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestALine(t *testing.T) {
	t.Parallel()
	tests := []struct {
		line string
		r    *Rule
	}{
		{
			line: "-A FORWARD -o br-4fbab684f8fb -m conntrack --ctstate RELATED,ESTABLISHED -c 12 244 -j ACCEPT",
			r: &Rule{
				Chain: "FORWARD",
				Match: "-o br-4fbab684f8fb -m conntrack --ctstate RELATED,ESTABLISHED",
				Jump:  "ACCEPT",
				Count: Count{Packets: 12, Bytes: 244},
			},
		},
		{
			line: "-A FORWARD -o br-4fbab684f8fb -m conntrack --ctstate RELATED,ESTABLISHED -c 12 244 -g ACCEPT",
			r: &Rule{
				Chain: "FORWARD",
				Match: "-o br-4fbab684f8fb -m conntrack --ctstate RELATED,ESTABLISHED",
				GoTo:  "ACCEPT",
				Count: Count{Packets: 12, Bytes: 244},
			},
		},
		{
			line: "-A FORWARD -o br-4fbab684f8fb -m conntrack --ctstate RELATED,ESTABLISHED -c 12 -j ACCEPT",
			r: &Rule{
				Chain: "FORWARD",
				Match: "-o br-4fbab684f8fb -m conntrack --ctstate RELATED,ESTABLISHED -c 12",
				Jump:  "ACCEPT",
			},
		},
		{
			line: "-A FORWARD -j ACCEPT",
			r: &Rule{
				Chain: "FORWARD",
				Jump:  "ACCEPT",
			},
		},
		{
			line: "-A FORWARD -c 1 2 -j ACCEPT",
			r: &Rule{
				Chain: "FORWARD",
				Jump:  "ACCEPT",
				Count: Count{Packets: 1, Bytes: 2},
			},
		},
		{
			line: "-A FORWARD  -o br-4fbab684f8fb -m conntrack --ctstate RELATED,ESTABLISHED",
			r: &Rule{
				Chain: "FORWARD",
				Match: "-o br-4fbab684f8fb -m conntrack --ctstate RELATED,ESTABLISHED",
			},
		},
		{
			line: "-A FORWARD  -o br-4fbab684f8fb -m conntrack --ctstate RELATED,ESTABLISHED -c 1 2",
			r: &Rule{
				Chain: "FORWARD",
				Match: "-o br-4fbab684f8fb -m conntrack --ctstate RELATED,ESTABLISHED",
				Count: Count{Packets: 1, Bytes: 2},
			},
		},
		{
			line: "-A FORWARD -c 0 0 -j REJECT --reject-with icmp-host-prohibited",
			r: &Rule{
				Chain:      "FORWARD",
				Jump:       "REJECT",
				TargetArgs: "--reject-with icmp-host-prohibited",
			},
		},
		{
			line: "-A FORWARD -i docker0 ! -o docker0 -c 15484 860607 -j ACCEPT",
			r: &Rule{
				Chain: "FORWARD",
				Match: "-i docker0 ! -o docker0",
				Jump:  "ACCEPT",
				Count: Count{Packets: 15484, Bytes: 860607},
			},
		},
		{
			line: "-B FORWARD -o br-4fbab684f8fb -m conntrack --ctstate RELATED,ESTABLISHED -j ACCEPT",
		},
	}

	for _, tt := range tests {
		r, err := Parse(tt.line)
		if tt.r == nil {
			assert.NotNil(t, err)
		} else {
			if assert.Nil(t, err) {
				assert.Equal(t, tt.r, r)
			}
		}
	}
}
