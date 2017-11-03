package rule

const (
	tipA = `-A, --append chain rule-specification
Append  one  or  more  rules to the end of the selected chain.  When the source and/or destination names resolve to more than one address, a
rule will be added for each possible address combination.
-4, --ipv4
This  option  has  no  effect  in  iptables and iptables-restore.  If a rule using the -4 option is inserted with (and only with) ip6tables-
restore, it will be silently ignored. Any other uses will throw an error. This option allows to put both IPv4 and IPv6  rules  in  a  single
rule file for use with both iptables-restore and ip6tables-restore.`

	tip6 = `-6, --ipv6
If  a  rule using the -6 option is inserted with (and only with) iptables-restore, it will be silently ignored. Any other uses will throw an
error. This option allows to put both IPv4 and IPv6 rules in a single rule file for use with both  iptables-restore  and  ip6tables-restore.
This option has no effect in ip6tables and ip6tables-restore.`

	tipp = `[!] -p, --protocol protocol
The  protocol of the rule or of the packet to check.  The specified protocol can be one of tcp, udp, udplite, icmp, icmpv6,esp, ah, sctp, mh
or the special keyword "all", or it can be a numeric value, representing one of these protocols or a different one.  A  protocol  name  from
/etc/protocols  is  also  allowed.   A  "!" argument before the protocol inverts the test.  The number zero is equivalent to all. "all" will
match with all protocols and is taken as default when this option is omitted.  Note that, in ip6tables, IPv6 extension  headers  except  esp
are  not  allowed.   esp and ipv6-nonext can be used with Kernel version 2.6.11 or later.  The number zero is equivalent to all, which means
that you cannot test the protocol field for the value 0 directly. To match on a HBH header, even if it were the last, you cannot use  -p  0,
but always need -m hbh.`

	tips = `[!] -s, --source address[/mask][,...]
Source  specification. Address can be either a network name, a hostname, a network IP address (with /mask), or a plain IP address. Hostnames
will be resolved once only, before the rule is submitted to the kernel.  Please note that specifying any name to be resolved with  a  remote
query such as DNS is a really bad idea.  The mask can be either an ipv4 network mask (for iptables) or a plain number, specifying the number
of 1's at the left side of the network mask.  Thus, an iptables mask of 24 is equivalent  to  255.255.255.0.   A  "!"  argument  before  the
address  specification  inverts  the sense of the address. The flag --src is an alias for this option.  Multiple addresses can be specified,
but this will expand to multiple rules (when adding with -A), or will cause multiple rules to be deleted (with -D).`

	tipd = `[!] -d, --destination address[/mask][,...]
Destination specification.  See the description of the -s (source) flag for a detailed description of the syntax.   The  flag  --dst  is  an
alias for this option.`

	tipm = `-m, --match match
Specifies  a  match  to use, that is, an extension module that tests for a specific property. The set of matches make up the condition under
which a target is invoked. Matches are evaluated first to last as specified on the command line and work in short-circuit fashion,  i.e.  if
one extension yields false, evaluation will stop.`

	tipj = `-j, --jump target
This  specifies  the  target of the rule; i.e., what to do if the packet matches it.  The target can be a user-defined chain (other than the
one this rule is in), one of the special builtin targets which decide the fate of the packet immediately, or an  extension  (see  EXTENSIONS
below).   If this option is omitted in a rule (and -g is not used), then matching the rule will have no effect on the packet's fate, but the
counters on the rule will be incremented.`

	tipg = `-g, --goto chain
This specifies that the processing should continue in a user specified chain. Unlike the --jump option return will not  continue  processing
in this chain but instead in the chain that called us via --jump.`

	tipi = `[!] -i, --in-interface name
Name  of  an  interface  via which a packet was received (only for packets entering the INPUT, FORWARD and PREROUTING chains).  When the "!"
argument is used before the interface name, the sense is inverted.  If the interface name ends in a "+", then  any  interface  which  begins
with this name will match.  If this option is omitted, any interface name will match.`

	tipo = `[!] -o, --out-interface name
Name of an interface via which a packet is going to be sent (for packets entering the FORWARD, OUTPUT and POSTROUTING chains).  When the "!"
argument is used before the interface name, the sense is inverted.  If the interface name ends in a "+", then  any  interface  which  begins
with this name will match.  If this option is omitted, any interface name will match.`

	tipf = `[!] -f, --fragment
This  means  that the rule only refers to second and further IPv4 fragments of fragmented packets.  Since there is no way to tell the source
or destination ports of such a packet (or ICMP type), such a packet will not match any rules which specify them.  When the "!" argument pre‐
cedes  the "-f" flag, the rule will only match head fragments, or unfragmented packets. This option is IPv4 specific, it is not available in
ip6tables.`

	tipc = `-c, --set-counters packets bytes
This enables the administrator to initialize the packet and byte counters of a rule (during INSERT, APPEND, REPLACE operations).`
)

var tooltip = map[string]string{
	"A": tipA, "append": tipA,
	"6": tip6, "ipv6": tip6,
	"p": tipp, "protocol": tipp,
	"s": tips, "source": tips,
	"d": tipd, "destination": tipd,
	"m": tipm, "match": tipm,
	"j": tipj, "jump": tipj,
	"g": tipg, "goto": tipg,
	"i": tipi, "ininterface": tipi,
	"o": tipo, "outinterface": tipo,
	"f": tipf, "fragment": tipf,
	"c": tipc, "setcounters": tipc,
}
