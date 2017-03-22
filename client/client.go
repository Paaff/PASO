package client

import (
	"fmt"
	"os/exec"
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
