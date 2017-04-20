package client

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// BT detection
func detectBluetooth(dataChannel chan blueData) {
	// Periodically scan for bluetooth devices.
	t := time.NewTicker(15 * time.Second)
	for {
		scan(dataChannel)
		<-t.C
	}

}

func scan(dataChannel chan blueData) {
	exec.Command("hcitool", "scan")
	out, err := exec.Command("hcitool", "inq").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}

	findAndDiscoverBClass(out, dataChannel)

}

// Trimming the output of Bluetooth inq command
func findAndDiscoverBClass(inq []byte, dataChannel chan blueData) {

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
				dataChannel <- phone
			}

		}

	}

}

// Takes a hexadecimal number and interprets the binary representation as what class is embedded there.
func checkBtClass(hexClass string) bool {
	// Strip the identifier 0x
	rawHex := hexClass[2:]

	// Find out if the binary representation matches that of a phone.
	classBits, err := convertBTClassHexToBinary(rawHex)
	if err != nil {
		log.Fatal(err)
	}
	return isFlipped(classBits, 22)
}

func isFlipped(val []uint64, n int) bool {
	index := len(val) - 1 - n
	return val[index] == 1
}

// Function to convert hex number in string type to its integer representation.
func convertBTClassHexToBinary(classHex string) ([]uint64, error) {
	// Class bit array is 24 bits long.
	bitArray := []uint64{}

	// Convert string to int
	classInt, err := strconv.ParseUint(classHex, 16, 64)
	if err != nil {
		return bitArray, fmt.Errorf("Error in string to int conversion: %v", err)
	}

	// Convert int to binary representation in an array type.
	bitArray = asBits(classInt)
	return bitArray, nil
}

func asBits(val uint64) []uint64 {
	bits := []uint64{}
	for i := 0; i < 24; i++ {
		bits = append([]uint64{val & 0x1}, bits...)
		val = val >> 1
	}
	return bits
}

func isMajorDeviceClassPhone(classBits []uint64) bool {
	/*
		Major Device Class Phone.
		Bit 	8 - 9 - 10 - 11 - 12
		Value 0 - 1 - 0  - 0  - 0
	*/
	return !isFlipped(classBits, 8) && isFlipped(classBits, 9) && !isFlipped(classBits, 10) &&
		!isFlipped(classBits, 11) && !isFlipped(classBits, 12)
}

func isMinorDeviceClassSmartPhone(classBits []uint64) bool {
	/*
		Minor Device Class Smartphone.
		Bit 	2 - 3 - 4 - 5 - 6 - 7
		Value 1 - 1 - 0 - 0 - 0 - 0
	*/
	return isFlipped(classBits, 2) && isFlipped(classBits, 3) && !isFlipped(classBits, 4) &&
		!isFlipped(classBits, 5) && !isFlipped(classBits, 6) && !isFlipped(classBits, 7)
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
