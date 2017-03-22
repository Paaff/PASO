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
	}
}

// Bluetooth detection
func detectBluetooth() {
	out, err := exec.Command("hcitool", "inq").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Printf("%s", out)

}

// Wifi detection
func detectWifi() {
	out, err := exec.Command("arp-scan", "-l").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Printf("%s", out)
}

// JSON struct for detected data.
type data struct {
	idData string
}
