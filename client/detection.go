package client

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/paaff/PASO/store"
)

type scanFunc func(dataChannel chan store.BlueData)

// BT detection
func detectBluetooth(dataChannel chan store.BlueData) {
	t := time.NewTicker(3 * time.Second)
	for {
		scan(dataChannel)
		<-t.C
	}
}

func scan(dataChannel chan store.BlueData) {
	cmd := exec.Command("hcitool", "scan")
	cmdErr := cmd.Start()
	if cmdErr != nil {
		log.Fatal(cmdErr)
	}
	fmt.Printf("Scanning...")
	cmdErr = cmd.Wait()
	if cmdErr != nil {
		log.Fatal(cmdErr)
	}

	out, err := exec.Command("hcitool", "inq").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}

	findAndDiscoverBTClass(out, dataChannel)

}

// Trimming the output of Bluetooth inq command
func findAndDiscoverBTClass(inq []byte, dataChannel chan store.BlueData) {
	var phone store.BlueData
	// split string up for each
	bluetoothList := strings.Split(string(inq), "\n")
	for i, line := range bluetoothList {
		// Disregard first line of hcitool inq as it just returns "Inqusåring ..."
		// And the last line, as it is empty
		if i > 0 && i != len(bluetoothList)-1 {
			bluetoothLine := strings.Fields(line)
			// Check that we have the correct class (Phone)
			className, ok := checkBTClass(bluetoothLine[5])
			if ok {
				phone = store.BlueData{Address: bluetoothLine[0], Class: className, Timestamp: time.Now()}
				fmt.Printf("The bluetooth address %v, and the class is %v\n", bluetoothLine[0], bluetoothLine[5])
				dataChannel <- phone
			} else {
				phone = store.BlueData{Address: bluetoothLine[0], Class: className, Timestamp: time.Now()}
				dataChannel <- phone
			}

		}

	}

}

// Takes a hexadecimal number and interprets the binary representation as what class is embedded there.
func checkBTClass(hexClass string) (string, bool) {
	// Strip the identifier 0x
	rawHex := hexClass[2:]

	// Find out if the binary representation matches that of a phone.
	classBits, err := convertBTClassHexToBinary(rawHex)
	if err != nil {
		log.Fatal(err)
	}
	if isMajorDeviceClassPhone(classBits) && isMinorDeviceClassSmartPhone(classBits) {
		return "Smartphone", true
	} else if isMajorDeviceClassPhone(classBits) && !isMinorDeviceClassSmartPhone(classBits) {
		return "Phone", true
	} else {
		return "Unknown", false
	}
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
