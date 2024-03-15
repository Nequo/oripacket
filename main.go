package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/pcapgo"
)

func main() {
	handle, err := pcap.OpenOffline("test.pcap")
	if err != nil {
		log.Fatal(err)
	}
	packetSource := gopacket.NewPacketSource(handle, layers.LayerTypeEthernet)
	packet, _ := packetSource.NextPacket()

	ethernetLayer, _ := packet.LinkLayer().(*layers.Ethernet)
	ipLayer, _ := packet.NetworkLayer().(*layers.IPv4)
	udpLayer, _ := packet.TransportLayer().(*layers.UDP)

	f, _ := os.Create("out.pcap")
	defer f.Close()

	w := pcapgo.NewWriter(f)
	w.WriteFileHeader(65535, layers.LinkTypeEthernet)

	inbytes, err := os.ReadFile("bytes.txt")
	if err != nil {
		log.Fatal(err)
	}

	buffer := gopacket.NewSerializeBuffer()
	options := gopacket.SerializeOptions{}
	err = gopacket.SerializeLayers(buffer, options,
		ethernetLayer,
		ipLayer,
		udpLayer,
		gopacket.Payload(inbytes),
	)
	if err != nil {
		log.Fatal(err)
	}
	packetBytes := buffer.Bytes()

	c := gopacket.CaptureInfo{}
	c.Timestamp = time.Unix(1710282234, 100)
	c.CaptureLength = len(packetBytes)
	c.Length = len(packetBytes)

	w.WritePacket(c, packetBytes)
	fmt.Println("Wrote out.pcap")
}
