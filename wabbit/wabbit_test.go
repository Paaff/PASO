package wabbit

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/NeowayLabs/wabbit"
	"github.com/NeowayLabs/wabbit/amqptest"
	"github.com/NeowayLabs/wabbit/amqptest/server"
)

func TestCreateDialPath(t *testing.T) {
	// Test data
	username := "username"
	pass := "pass"
	address := "192.168.0.150"
	port := "8080"

	expected := "amqp://username:pass@192.168.0.150:8080/"
	actual := createDialPath(username, pass, address, port)
	if expected != actual {
		t.Errorf("Expected: %s\nActual: %s", expected, actual)
	}

}

type TestBT struct {
	Bdaddress string
	Class     string
}

func TestPublisherConsumer(t *testing.T) {
	go testConsumeHelp(t)
	go testPublishHelp(t)

}

func testConsumeHelp(t *testing.T) {
	// Mocking server.
	mockServer := server.NewServer("amqp://testUser:testPass@192.168.0.106:5672/")
	mockServer.Start()

	w := &Wabbit{
		Connection: nil,
		Channel:    nil,
	}

	mockConn, err := amqptest.Dial("amqp://testUser:testPass@192.168.0.106:5672/")
	if err != nil {
		t.Error("Connection to the mock server failed.")
	}
	w.Connection = mockConn

	mockCh, err := mockConn.Channel()
	if err != nil {
		t.Error("Channel creation failed.")
	}

	if mockCh.ExchangeDeclare("TestExchangeName", "direct", wabbit.Option{
		"durable":  true,
		"delete":   false,
		"internal": false,
		"noWait":   false,
	}); err != nil {
		t.Error("Mock Exchange declaration failed")
	}
	w.Channel = mockCh

	mockQ, err := mockCh.QueueDeclare("TestQueue", wabbit.Option{
		"durable":   true,
		"delete":    false,
		"exclusive": false,
		"noWait":    false,
	})
	if err != nil {
		t.Error("Mock Queue declaration failed")
	}

	mockCh.QueueBind(mockQ.Name(), "TestRoutingKey", "TestExchangeName", wabbit.Option{
		"durable":  true,
		"delete":   false,
		"internal": false,
		"noWait":   false,
	})
	if err != nil {
		t.Error("Queue bind failed")
	}

	// Consuming
	var btData TestBT
	recieverChan := make(chan []byte)
	//go w.ConsumeMessage(mockQ.Name(), recieverChan)

	for data := range recieverChan {
		err := json.Unmarshal(data, &btData)
		if err != nil {
			t.Errorf("Unmarshalling failed: %v", err)
		}
		fmt.Println(btData.Bdaddress, btData.Class)
	}

}

func testPublishHelp(t *testing.T) {
	// Mocking server.
	mockServer := server.NewServer("amqp://testUser:testPass@192.168.0.106:5672/")
	mockServer.Start()

	w := &Wabbit{
		Connection: nil,
		Channel:    nil,
	}

	mockConn, err := amqptest.Dial("amqp://testUser:testPass@192.168.0.106:5672/")
	if err != nil {
		t.Error("Connection to the mock server failed.")
	}
	w.Connection = mockConn

	mockCh, err := mockConn.Channel()
	if err != nil {
		t.Error("Channel creation failed.")
	}

	if mockCh.ExchangeDeclare("TestExchangeName", "direct", wabbit.Option{
		"durable":  true,
		"delete":   false,
		"internal": false,
		"noWait":   false,
	}); err != nil {
		t.Error("Mock Exchange declaration failed")
	}
	w.Channel = mockCh

	// Publishing
	d := &TestBT{
		Bdaddress: "TestBTAddress",
		Class:     "TestClass",
	}
	json, err := json.Marshal(d)
	if err != nil {
		t.Error("Test BT json marshal failed")
	}
	err = w.PublishMessage(json, "TestExchangeName", "TestRoutingKey")
	if err != nil {
		t.Errorf("Publishing failed, err: %v", err)
	}

}
