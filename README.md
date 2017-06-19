# PASO
Presence-Aware Smart Office - Bachelor Thesis 2017 by Peter Fischer.

# Prerequisite
Installing a RabbitMQ server.
Running it on a Raspberry Pi with the ´hcitool´ available.

# Setup & Running
Assumed that the RabbitMQ server has been installed, the server- and client-configuration file must be filled out with the proper information. Please use the available templates in the config folder.

Use the ´go build´ command to compile the system.

Two arguments are available for the system. ´server´ and ´client´. 

Start the system server as such:

./PASO -launch=server

Start the client (detection device) as such:

./PASO
or 
./PASO -launch=client
