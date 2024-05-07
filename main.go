package main

import (
	"fmt"
	"log"
	utils "monitoring/utils"
	"net"
	"time"

	"github.com/gosnmp/gosnmp"
)

func main() {
	snmpwalk()
}

func raw_connect(host string, ports []string) {
	for _, port := range ports {
		timeout := time.Second
		conn, err := net.DialTimeout("udp", net.JoinHostPort(host, port), timeout)
		if err != nil {
			fmt.Println("Connection error", err)
		}
		if conn != nil {
			defer conn.Close()
			fmt.Println("Opened ", net.JoinHostPort(host, port))
		}
	}
}

func myWalkFn(pdu gosnmp.SnmpPDU) error {
	if pdu.Type == gosnmp.OctetString {
		var value interface{} = utils.TranslatePDU(pdu)
		fmt.Printf("OID: %s, Value: %v\n", pdu.Name, value)
	} else {
		fmt.Printf("OID %s, Value (other): %v\n", pdu.Name, pdu.Name)
	}
	return nil
}

func snmpwalk() {
	g := &gosnmp.GoSNMP{
		Target:    "192.168.0.115",
		Port:      161,
		Community: "testing123",
		Version:   gosnmp.Version2c,
		Timeout:   time.Duration(2) * time.Second,
	}

	err := g.Connect()
	if err != nil {
		log.Fatalf("Connect() err: %v", err)
	}

	defer g.Conn.Close()

	var rootOid string = "1.3.6.1.2.1.2.2"
	err = g.BulkWalk(rootOid, myWalkFn)
	if err != nil {
		log.Fatalf("Walk() err: %v", err)
	}

}
