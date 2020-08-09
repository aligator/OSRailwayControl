# /bin/sh
mkdir build

go generate

echo "Building darwin"
GOOS=darwin GOARCH=amd64 go build
mv OSRailwayControl build/OSRailwayControl_macOS_amd64

echo "Building Linux"
GOOS=linux GOARCH=amd64 go build
mv OSRailwayControl build/OSRailwayControl_linux_amd64
GOOS=linux GOARCH=arm	GOARM=7 go build
mv OSRailwayControl build/OSRailwayControl_linux_arm
GOOS=linux GOARCH=arm64 go build
mv OSRailwayControl build/OSRailwayControl_linux_arm64

echo "Building windows"
GOOS=windows GOARCH=amd64 go build
mv OSRailwayControl.exe build/OSRailwayControl_windows_amd64.exe

echo "Completed"