package client

import (
	"fmt"
	"os/exec"
	"runtime"
)

// Start - Global function to start the client.
func Start() {
	if runtime.GOOS == "linux" {
		fmt.Println("Unix type OS detected")
		detectBluetooth()
	}
}

// Bluetooth detection
func detectBluetooth() {
	out, err := exec.Command("hcitool", "inq").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}

	trimOutput(out)

}

// Trimming the output of Bluetooth inq command
func trimOutput(inq []byte) {

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
