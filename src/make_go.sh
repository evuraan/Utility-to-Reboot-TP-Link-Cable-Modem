#!/bin/bash

# Author: evuraan@gmail.com
src="reboot_tc7610.go"

mkdir -pv ../bin/

go get  "golang.org/x/net/publicsuffix" && {
	GOOS=linux GOARCH=arm GOARM=7 go build  -o ../bin/reboot_tc7610-linux-armv71 $src
	GOOS=linux GOARCH=arm GOARM=6 go build  -o ../bin/reboot_tc7610-linux-armv6l $src
	GOOS=windows GOARCH=amd64 go build -o ../bin/reboot_tc7610-win-amd64.exe $src
	GOOS=windows GOARCH=386 go build -o ../bin/reboot_tc7610-win-386.exe $src
	GOOS=linux GOARCH=amd64 go build -o ../bin/reboot_tc7610-linux-amd64 $src
	GOOS=linux GOARCH=386 go build -o ../bin/reboot_tc7610-linux-386 $src
}


