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
			line: "[12:244] -A FORWARD -o br-4fbab684f8fb -m conntrack --ctstate RELATED,ESTABLISHED -j ACCEPT",
			r: &Rule{
				Chain: "FORWARD",
				Match: []Arg{{Key: "-o", Value: []string{"br-4fbab684f8fb"}}, {Key: "-m", Value: []string{"conntrack"}}, {Key: "--ctstate", Value: []string{"RELATED,ESTABLISHED"}}},
				Jump:  "ACCEPT",
				Count: Count{Packets: 12, Bytes: 244},
			},
		},
		{
			line: "[12:244] -A FORWARD -o br-4fbab684f8fb -m conntrack --ctstate RELATED,ESTABLISHED -g ACCEPT",
			r: &Rule{
				Chain: "FORWARD",
				Match: []Arg{{Key: "-o", Value: []string{"br-4fbab684f8fb"}}, {Key: "-m", Value: []string{"conntrack"}}, {Key: "--ctstate", Value: []string{"RELATED,ESTABLISHED"}}},
				GoTo:  "ACCEPT",
				Count: Count{Packets: 12, Bytes: 244},
			},
		},
		{
			line: "[12:] -A FORWARD -o br-4fbab684f8fb -m conntrack --ctstate RELATED,ESTABLISHED -j ACCEPT",
		},
		{
			line: "[0:0] -A FORWARD -j ACCEPT",
			r: &Rule{
				Chain: "FORWARD",
				Jump:  "ACCEPT",
			},
		},
		{
			line: "[1:2] -A FORWARD -j ACCEPT",
			r: &Rule{
				Chain: "FORWARD",
				Jump:  "ACCEPT",
				Count: Count{Packets: 1, Bytes: 2},
			},
		},
		{
			line: "[0:0] -A FORWARD  -o br-4fbab684f8fb -m conntrack --ctstate RELATED,ESTABLISHED",
			r: &Rule{
				Chain: "FORWARD",
				Match: []Arg{{Key: "-o", Value: []string{"br-4fbab684f8fb"}}, {Key: "-m", Value: []string{"conntrack"}}, {Key: "--ctstate", Value: []string{"RELATED,ESTABLISHED"}}},
			},
		},
		{
			line: "[1:2] -A FORWARD  -o br-4fbab684f8fb -m conntrack --ctstate RELATED,ESTABLISHED",
			r: &Rule{
				Chain: "FORWARD",
				Match: []Arg{{Key: "-o", Value: []string{"br-4fbab684f8fb"}}, {Key: "-m", Value: []string{"conntrack"}}, {Key: "--ctstate", Value: []string{"RELATED,ESTABLISHED"}}},
				Count: Count{Packets: 1, Bytes: 2},
			},
		},
		{
			line: "[0:0] -A FORWARD -j REJECT --reject-with icmp-host-prohibited",
			r: &Rule{
				Chain:      "FORWARD",
				Jump:       "REJECT",
				TargetArgs: []Arg{{Key: "--reject-with", Value: []string{"icmp-host-prohibited"}}},
			},
		},
		{
			line: "[15484:860607] -A FORWARD -i docker0 ! -o docker0 -j ACCEPT",
			r: &Rule{
				Chain: "FORWARD",
				Match: []Arg{{Key: "-i", Value: []string{"docker0"}}, {Key: "-o", Value: []string{"docker0"}, Not: true}},
				Jump:  "ACCEPT",
				Count: Count{Packets: 15484, Bytes: 860607},
			},
		},
		{line: "[0:0] -B FORWARD -o br-4fbab684f8fb -m conntrack --ctstate RELATED,ESTABLISHED -j ACCEPT"},
		{line: "[0:0] -A FORWARD -o br-4fbab684f8fb -m conntrack --ctstate RELATED,ESTABLISHED -g"},
		{line: "[0:0] -A FORWARD -o br-4fbab684f8fb -m conntrack --ctstate RELATED,ESTABLISHED -j"},
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
