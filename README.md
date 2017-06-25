# PASO
Presence-Aware Smart Office - Bachelor Thesis 2017 by Peter Fischer.

## Getting Started
As the system has been developed with the use of multiple Raspberry Pi 3's, it is assumed that such a device is used here.
The following instructions will lead you to running the system on said Raspberry Pi's. Furthermore, there is a client and server take on this system. One Raspberry Pi will run the *server* side of the system, whereas the second Raspberry Pi will act as the *client*.

It is not required, but can be recommended to have the [Raspbian](https://www.raspberrypi.org/downloads/raspbian/) operating system installed on the Raspberry Pi.

### Prerequisite
Before the system can run, there are a few things needed. A **RabbitMQ** server is used for communication between the *client* and *server* sides of the system, whereas the **hcitool** is used for the Bluetooth detection on the *client* only. 

#### RabbitMQ
The RabbitMQ server can in theory be installed on any device available as long as the IP used in the system is correct. However, here it is assumed that it is installed on the Raspberry Pi where the *server* side is ran.

To install the RabbitMQ server, follow the instructions given [here](https://www.rabbitmq.com/install-debian.html).


#### hcitool (Needed on the *client* side ran on a different Raspberry Pi)
Before using, remember to update and install the necessary Bluetooth software for your specific setup. The end goal is getting the **hcitool** to function, using the commmands:

* **hcitool inq**
* **hcitool scan**

If there are problems, searching the Raspberry Pi forums is a great way to get help. Though, the initial way of making sure the tool is working is to do the following:

```
sudo apt-get update
sudo apt-get upgrade
sudo apt-get dist-upgrade
sudo rpi-update
```
This should make sure your Raspbian OS is fully updated. Afterwards, the **hcitool** should be functional, though you might need to start the Bluetooth dongle on the Raspberry Pi with **bluetoothctl** and **hciconfig**. Please use the Raspberry Pi forums if help is needed.

#### Go (Optional)
It is not needed to install Go on any of the devices if you build the system on another device with Go already installed. However, if you do chose to build it on the Raspberry Pi, you have to install the Go binaries. 

Furthermore, it is wise to setup the Go environment following these [instructions](https://golang.org/doc/install)



## Setup & Running
Here it is assumed that a RabbitMQ server is available and the Raspberry Pi in charge of detecting has working **hcitool** functionality. 

### Compiling
In the event that you want to compile the system yourself and have followed the environment setup from the Go instructions, the following command will download the source code in your environment:

```
go get github.com/paaff/PASO
```
Navigate to the *paaff/PASO/* folder and run the following:
```
go build
```
Now a *PASO* executable should appear in the folder.

### Configuration
Before launching the system, the appropiate configuration files has to be filled out. Here, the *config* folder contains templates for the *server* and *client* side of the system, where the RabbitMQ server information has to be inserted. 

Make a copy of the template files, fill out the information needed and remove the *_example* addition in the file.

### Launching
The system is parted in two as explained, a *server* and *client* side. To launch the *server*, run the following command in the folder where the *PASO* executable resides.

```
./PASO -launch=server
```
This should start the *server*, connecting to the RabbitMQ server and starting the webserver on **localhost:3000**. As this is up and running, the the *client* can be launched. Run one of the two commands:

```
./PASO -launch=client
./PASO
```

