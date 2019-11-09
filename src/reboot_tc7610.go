/*
 * ----------------------------------------------------------------------------
    reboot_tc7610 : utility to reboot TP-Link Cable Modem
    Copyright (C) 2019  Evuraan, <evuraan@gmail.com>

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.
 * ----------------------------------------------------------------------------
*/
package main

import (
	"fmt"
	"golang.org/x/net/publicsuffix"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
	"time"
)

const (
	Ver     = "1.0g"
	Version = "tc7610 rebooter " + Ver
)

var (
	client   = http.Client{Timeout: 15 * time.Second}
	ip       = "192.168.100.1"
	user     = "admin"
	password = "admin"
)

func main() {

	fmt.Printf("%s Copyright (C) 2019 Evuraan <evuraan@gmail.com>\n", Version)
	fmt.Println("This program comes with ABSOLUTELY NO WARRANTY.")

	argc := len(os.Args)
	if argc == 1 {
		fmt.Println("Assuming default address and credentials")
	} else if argc > 1 {
		for i, arg := range os.Args {
			if arg == "help" || arg == "--help" || arg == "h" || arg == "--h" || arg == "-h" || arg == "-help" || arg == "?" {
				showhelp()
				os.Exit(0)
			}

			if arg == "version" || arg == "--version" || arg == "v" || arg == "--v" || arg == "-v" || arg == "-version" {
				fmt.Println("Version:", Version)
				os.Exit(0)
			}

			if arg == "ip" || arg == "--ip" || arg == "ip" || arg == "--ip" || arg == "-ip" || arg == "-ip" {
				next := i + 1
				if argc > next {
					ip = os.Args[next]
				} else {
					fmt.Println("Invalid usage")
					showhelp()
					os.Exit(1)
				}
			}

			//	if arg == "port" || arg == "--port" || arg == "p" || arg == "--p" || arg == "-p" || arg == "-port" {
			if arg == "user" || arg == "--user" || arg == "u" || arg == "--u" || arg == "-u" || arg == "-user" {
				next := i + 1
				if argc > next {
					user = os.Args[next]
				} else {
					fmt.Println("Invalid usage")
					showhelp()
					os.Exit(1)
				}
			}

			if arg == "password" || arg == "--password" || arg == "p" || arg == "--p" || arg == "-p" || arg == "-password" {
				next := i + 1
				if argc > next {
					password = os.Args[next]
				} else {
					fmt.Println("Invalid usage")
					showhelp()
					os.Exit(1)
				}
			}
		}
	}

	//fmt.Printf("ip: %s, user: %s, pass: %s\n", ip, user, password)
	//os.Exit(0)

	login_url := fmt.Sprintf("http://%s/goform/login", ip)
	rebooturl := fmt.Sprintf("http://%s/goform/top", ip)
	Referer := fmt.Sprintf("http://%s/index.htm", ip)

	options := cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	}
	jar, err := cookiejar.New(&options)
	if err != nil {
		fmt.Println("Cookkie err", err)
	}

	client = http.Client{Jar: jar}
	//meh := `{"TopReboot": "1"}`
	meh := "TopReboot=1"
	data := url.Values{
		"loginUsername": {user},
		"loginPassword": {password},
		"data":          {meh},
	}

	r, _ := http.NewRequest("POST", login_url, strings.NewReader(data.Encode()))
	r.Header.Add("user-agent", "")
	r.Header.Add("Referer", Referer)
	doReq, err := client.Do(r)

	if err != nil || doReq.StatusCode != 200 {
		fmt.Println("\nError during post", err)
		os.Exit(1)
	}

	reboot, _ := http.NewRequest("POST", rebooturl, strings.NewReader(data.Encode()))
	reboot.Header.Add("user-agent", "")
	reboot.Header.Add("Referer", Referer)
	doReboot, err := client.Do(reboot)

	if err != nil || doReboot.StatusCode != 200 {
		fmt.Println("Error during reboot", err)
		os.Exit(1)
	} else {
		fmt.Println("Reboot Success!")
	}

}

func showhelp() {
	fmt.Printf("Usage: %s -ip 192.168.100.1 -u administrator -p password\n", os.Args[0])
	fmt.Println("  -h   --help         print this usage and exit")
	fmt.Println("  -ip  --ip           IP address of the modem")
	fmt.Println("  -u   --user         username to use")
	fmt.Println("  -p   --password     password to use")
	fmt.Println("  -v   --version      print version information and exit")
}
