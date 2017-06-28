package demo

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/paaff/PASO/config"
	"github.com/paaff/PASO/store"
	"github.com/paaff/PASO/wabbit"
	"gopkg.in/dixonwille/wmenu.v4"
)

var xName string
var key string
var w *wabbit.Wabbit

// StartDemo presents the available demo actions
func StartDemo(conf *config.Config) {
	conn, err := wabbit.InitWabbitPublisher(conf.Username, conf.Pass, conf.Address, conf.Port, conf.ExchangeName, conf.ExchangeType, conf.RoutingKey)
	if err != nil {
		fmt.Println("Ini of publisher failed")
	}
	w = conn
	defer w.Connection.Close()
	defer w.Channel.Close()
	xName = conf.ExchangeName
	key = conf.RoutingKey
	createMenu()
}

func createMenu() {
	var optFunc = func(option wmenu.Opt) error {
		fmt.Fprintf(os.Stdout, "Scenario %s was chosen.", option.Value)
		switch option.Value {
		case "1":
			scenario1()
		case "2":
			scenario2()
		case "3":
			scenario3()
		}

		return nil
	}
	menu := wmenu.NewMenu("Choose an option.")
	menu.Option("User Scenario 1", "1", false, optFunc)
	menu.Option("User Scenario 2", "2", false, optFunc)
	menu.Option("User Scenario 3", "3", false, optFunc)
	err := menu.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func scenario1() {
	fmt.Fprint(os.Stdout, "User scenario 1 - Detecting Alice.\nPublishing detection to server.")

	alice := store.BlueData{Address: "Alice1_ID", Class: "Smartphone", Timestamp: time.Now()}
	publishScenarioData(alice)

	fmt.Fprint(os.Stdout, "Halted - Continue?\n")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

	fmt.Fprint(os.Stdout, "Detecting Bob.\nPublishing detection to server.")
	bob := store.BlueData{Address: "Bob1_ID", Class: "Smartphone", Timestamp: time.Now()}
	publishScenarioData(bob)

}

func scenario2() {
	fmt.Fprint(os.Stdout, "User scenario 2 - Detecting Bob.\nPublishing detection to server.")

	bob := store.BlueData{Address: "Bob2_ID", Class: "Smartphone", Timestamp: time.Now()}
	publishScenarioData(bob)

	fmt.Fprint(os.Stdout, "Halted - Continue?\n")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

	fmt.Fprint(os.Stdout, "Detecting Alice.\nPublishing detection to server.")
	alice := store.BlueData{Address: "Alice2_ID", Class: "Smartphone", Timestamp: time.Now()}
	publishScenarioData(alice)

}

func scenario3() {
	fmt.Fprint(os.Stdout, "User scenario 3 - Detecting Alice and Bob.\nPublishing detection to server.")

	bob := store.BlueData{Address: "Bob3_ID", Class: "Smartphone", Timestamp: time.Now()}
	publishScenarioData(bob)
	alice := store.BlueData{Address: "Alice3_ID", Class: "Smartphone", Timestamp: time.Now()}
	publishScenarioData(alice)

	fmt.Fprint(os.Stdout, "Halted - Continue?\n")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

	fmt.Fprint(os.Stdout, "Detecting Intruder.\nPublishing detection to server.")
	trudy := store.BlueData{Address: "Alice2_ID", Class: "Motion Sensor", Timestamp: time.Now()}
	publishScenarioData(trudy)
}

func publishScenarioData(data store.BlueData) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("JSON Marshal went wrong")
	}
	if err = w.PublishMessage(jsonData, xName, key); err != nil {
		fmt.Printf("Publish went wrong. Err: %v\n", err)
	}
}
