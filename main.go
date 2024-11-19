package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/maxmind/mmdbwriter"
	"github.com/maxmind/mmdbwriter/mmdbtype"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: mrt2mmdb <mrtFile> <mmdbOutputFile>")
		os.Exit(1)
	}

	mrtFile := os.Args[1]

	mmdbOutputFile := os.Args[2]

	fmt.Printf("MRT File: %s\n", mrtFile)
	fmt.Printf("MMDB Output File: %s\n", mmdbOutputFile)
	writer, err := mmdbwriter.New(
		mmdbwriter.Options{
			DatabaseType: "My-ASN-DB",
			RecordSize:   24,
		},
	)
	// for {
	record := mmdbtype.Map{}
	asn := 0 // TODO
	as_name := ""
	ip_address := "123.123.123.123/24"

	_, network, err := net.ParseCIDR(ip_address)
	if err != nil {
		log.Fatal(err)
	}

	if asn != 0 {
		record["asn"] = mmdbtype.Uint32(asn)
	}

	if as_name != "" {
		record["as_name"] = mmdbtype.String(as_name)
	}

	err = writer.Insert(network, record)
	// }

	fh, err := os.Create("out.mmdb")
	if err != nil {
		log.Fatal(err)
	}

	_, err = writer.WriteTo(fh)
	if err != nil {
		log.Fatal(err)
	}
}
