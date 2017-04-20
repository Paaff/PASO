package client

import (
	"log"
	"reflect"
	"testing"
)

func TestConvertBTClassHexToBinary(t *testing.T) {
	// Hex string matching Smartphone.
	hex := "5a020c"
	expected := []uint64{0, 1, 0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 1, 0, 0}

	actual, err := convertBTClassHexToBinary(hex)
	if err != nil {
		t.Errorf("There was an error even though the correct types were used. Err: %v", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Actual was not as expected.\nActual: %v\nExpected: %v", actual, expected)
	}

}

func TestCheckBitN(t *testing.T) {
	// Hex string matching Smartphone.
	hex := "5a020c"
	classBits, err := convertBTClassHexToBinary(hex)
	if err != nil {
		log.Fatal(err)
	}

	/*
	  Service classes from the hex string is:
	  Networking - Bit 17
	  Capturing - Bit 19
	  Object Transfer - Bit 20
	  Telephony - Bit 22

	  Major Device Class is Phone.
	  Bit 	8 - 9 - 10 - 11 - 12
	  Value 0 - 1 - 0  - 0  - 0
	*/

	// Service classes
	if !isFlipped(classBits, 17) {
		t.Errorf("Networking bit 17 was not flipped, we expect it to be. ClassBits: %v", classBits)
	}
	if !isFlipped(classBits, 19) {
		t.Errorf("Capturing bit 19 was not flipped, we expect it to be. ClassBits: %v\nBit 19, index 19: %v", classBits, classBits[19])
	}
	if !isFlipped(classBits, 20) {
		t.Errorf("Object Transfer bit 20 was not flipped, we expect it to be. ClassBits: %v", classBits)
	}
	if !isFlipped(classBits, 22) {
		t.Errorf("Telephony bit 22 was not flipped, we expect it to be. ClassBits: %v", classBits)
	}

}

func TestIsMajorDeviceClassPhone(t *testing.T) {
	// Hex string matching Smartphone.
	hex := "5a020c"
	classBits, err := convertBTClassHexToBinary(hex)
	if err != nil {
		log.Fatal(err)
	}

	// This should be a Phone in the major device.
	actual := isMajorDeviceClassPhone(classBits)
	if !actual {
		t.Errorf("We expected the hex string to match a major device as a Phone, but it returned: %v\n%v", actual, classBits)
	}
}

func TestIsMinorDeviceClassSmartPhone(t *testing.T) {
	hex := "5a020c"
	classBits, err := convertBTClassHexToBinary(hex)
	if err != nil {
		log.Fatal(err)
	}

	actual := isMinorDeviceClassSmartPhone(classBits)
	if !actual {
		t.Errorf("We expected the hex string to match a minor device as a Smartphone, but it returned: %v\n%v", actual, classBits)
	}
}
