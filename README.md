# Oripacket

Tool for crafting packets which can be folded back into
the original capture file.


## Current spec

Input: UDP pcap (test.pcap) + file containing payload bytes (bytes.txt)
Output: pcap with single packet (out.pcap)


```
~/p/oripacket main• ❱ xxd test.pcap 
00000000: 4d3c b2a1 0200 0400 0000 0000 0000 0000  M<..............
00000010: ffff 0000 0100 0000 6c86 5741 80ff f237  ........l.WA...7
00000020: 4b00 0000 4b00 0000 000c 4182 b253 00d0  K...K.....A..S..
00000030: 596c 404e 0800 4500 003d 0a41 0000 8011  Yl@N..E..=.A....
00000040: 7ceb c0a8 3232 c0a8 0001 ff02 ff35 0029  |...22.......5.)
00000050: 07a9 002b 0100 0001 0000 0000 0000 0275  ...+...........u
00000060: 7304 706f 6f6c 036e 7470 036f 7267 0000  s.pool.ntp.org..
00000070: 0100 01                                  ...
~/p/oripacket main• ❱ go run .
Wrote out.pcap
~/p/oripacket main• ❱ xxd out.pcap 
00000000: d4c3 b2a1 0200 0400 0000 0000 0000 0000  ................
00000010: ffff 0000 0100 0000 fad5 f065 0000 0000  ...........e....
00000020: 3c00 0000 3c00 0000 000c 4182 b253 00d0  <...<.....A..S..
00000030: 596c 404e 0800 4500 003d 0a41 0000 8011  Yl@N..E..=.A....
00000040: 7ceb c0a8 3232 c0a8 0001 ff02 ff35 0029  |...22.......5.)
00000050: 07a9 4865 6c6c 6f20 6672 6f6d 2066 696c  ..Hello from fil
00000060: 650a 0000                                e...
```

## References

- [Pcap from wireshark samples](https://wiki.wireshark.org/uploads/__moin_import__/attachments/SampleCaptures/dns_port.pcap)
- [GoPacket examples](https://www.devdungeon.com/content/packet-capture-injection-and-analysis-gopacket)
