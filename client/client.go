package client

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

// Start - Global function to start the client.
func Start() {
	detectBluetooth()

}

// Bluetooth detection
func detectBluetooth() {
	exec.Command("hcitool", "scan")
	out, err := exec.Command("hcitool", "inq").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}

	trimBtOutput(out)

}

// Trimming the output of Bluetooth inq command
func trimBtOutput(inq []byte) {
	var result []blueData
	// split string up for each
	bluetoothList := strings.Split(string(inq), "\n")
	for i, line := range bluetoothList {
		// Disregard first line of hcitool inq as it just returns "Inquring ..."
		// And the last line, as it is empty
		if i > 0 && i != len(bluetoothList) {
			bluetoothLine := strings.Fields(line)
			blueData{
				bdaddress: bluetoothLine[0],
				class:     bluetoothLine[5]}
		}
	}

}

// Takes a hexadecimal number and interprets the binary representation as what class is embedded there.
func discoverBtClass(hexClass string) bool {
	// Strip the identifier 0x
	rawHex := hexClass[2:]

	// Convert string to int
	classInt, err := strconv.ParseUint(rawHex, 16, 32)
	if err != nil {
		fmt.Printf("%s", err)
	}

	// Convert int to binary representation
	// %024b indicates base 2, padding with 0, with 24 characters.
	classBin := fmt.Sprintf("%024b", classInt)

	// Find out if the binary representation matches that of a phone.
	/*
		Bit 22 = Telephony
		Bit 12-11-10-9-8 = 00010 = Phone
		Bit 7-6-5-4-3-2 = 000011 = Smart Phone
	*/

	return false
}

// Wifi detection
func detectWifi() {
	out, err := exec.Command("arp-scan", "-l").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Printf("%s", out)
}

// Bluetooth data.
type blueData struct {
	bdaddress string
	class     string
}

// Wifi data
type wifiData struct {
	mac string
}
