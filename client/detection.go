package client

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// Collection of found phones
var detectedPhones []blueData

// BT detection
func detectBluetooth(dataChannel chan blueData) {

	// Periodically scan for bluetooth devices.
	t := time.NewTicker(15 * time.Second)
	for {
		dataChannel <- scan()
		<-t.C
	}

}

func scan() blueData {
	exec.Command("hcitool", "scan")
	out, err := exec.Command("hcitool", "inq").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}

	return findAndDiscoverBClass(out)

}

// Trimming the output of Bluetooth inq command
func findAndDiscoverBClass(inq []byte) blueData {

	//TODO: Refactor phone return
	var phone blueData
	// split string up for each
	bluetoothList := strings.Split(string(inq), "\n")
	for i, line := range bluetoothList {

		// Disregard first line of hcitool inq as it just returns "Inquring ..."
		// And the last line, as it is empty
		if i > 0 && i != len(bluetoothList)-1 {
			bluetoothLine := strings.Fields(line)

			// Check that we have the correct class (Phone)
			if checkBtClass(bluetoothLine[5]) {
				phone = blueData{bdaddress: bluetoothLine[0], class: bluetoothLine[5]}
				fmt.Printf("The bluetooth address %v, and the class is %v\n", bluetoothLine[0], bluetoothLine[5])
				return phone
			}

		}

	}
	return phone

}

// Takes a hexadecimal number and interprets the binary representation as what class is embedded there.
func checkBtClass(hexClass string) bool {
	// Strip the identifier 0x
	rawHex := hexClass[2:]

	// Convert string to int
	classInt, err := strconv.ParseUint(rawHex, 16, 32)
	if err != nil {
		fmt.Printf("%s", err)
	}
	// Find out if the binary representation matches that of a phone.
	/*
		Bit 22 = Telephony
		Bit 12-11-10-9-8 = 00010 = Phone
		Bit 7-6-5-4-3-2 = 000011 = Smart Phone
	*/

	return checkBitN(classInt, 22)
}

// TODO: 64 and 32 bit, will this clash?
func checkBitN(val uint64, n uint32) bool {
	return 1<<n&val > 0
}

// Wifi detection
func detectWifi() {
	out, err := exec.Command("arp-scan", "-l").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Printf("%s", out)
}

// blueData - Bluetooth data.
type blueData struct {
	bdaddress string
	class     string
}

// Wifi data
type wifiData struct {
	mac string
}
