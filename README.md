# OSRailway Train

This is a server to control [OSRailway](https://www.thingiverse.com/thing:4408535) trains using the [OSRailway Train firmware](https://github.com/aligator/OSRailwayTrain).
It basically is a mqtt client which provides a simple web-gui (it starts a small webserver) to manage several trains. 

The software is in an early stage.

# How to build

Prequesites
- golang compiler
- go get github.com/markbates/pkger/cmd/
- MQTT broker - without it the whole firmware is not usable.

To run it you just have to execute these commands
- `go generate` to embed the html-js client in the app.
- `go run . --web-port 3000 --mqtt-host 192.168.178.90 --mqtt-port 1883 --mqtt-user yourUser --mqtt-password YourPassword`
Then the UI is available at `http://localhost:3000`. 

# Binary Releases

[The latest release can be found here for all popular platforms.](https://github.com/aligator/OSRailwayControl/releases)
Just download and unpack the executable you need.
To get all possible command line options just run:  
`./osrailway-control --help` (linux / mac)  
`osrailway-control.exe --help` (windows)  

# Web client

The web client is currently built with plain js and html to avoid big dependencies like React or similar.
This may change in the future. It connects to the backend using websockets.
