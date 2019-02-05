package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
)

// prettyPrint formats a JSON response as a pretty-printed string
func prettyPrint(buffer []byte) string {
	var indented bytes.Buffer
	json.Indent(&indented, buffer, "", "    ")
	return string(indented.Bytes())
}

// ipLocation returns the location from an IP address
func ipLocation(ip net.IPAddr) ([]byte, error) {
	response, err := http.Get("http://ip-api.com/json/" + ip.String())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	flag.Usage = func() {
		fmt.Println(
			"iplocation - fetch the location of a website or IP address\n\n" +
				"Usage: iplocation <ip or dns address>\n")
		flag.PrintDefaults()
	}

	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	input := flag.Args()[0]
	ip, err := net.ResolveIPAddr("ip", input)
	handleError(err)

	location, err := ipLocation(*ip)
	handleError(err)

	fmt.Println(prettyPrint(location))
}
