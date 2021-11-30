all: build_gateway build_device build_sensor
	cp ./scripts/run.sh build/package
	tar -cf build/project3.tar build/package

build_gateway:
	env GOOS=linux GOARCH=arm GOARM=5 go build -o build/package/gateway/gateway github.com/fishjump/cs7ns1_project3/cmd/gateway
	cp ./demoCA/keys/bundled.* build/package/gateway
	cp ./demoCA/keys/ca.crt build/package/gateway

build_device:
	env GOOS=linux GOARCH=arm GOARM=5 go build -o build/package/device/device github.com/fishjump/cs7ns1_project3/cmd/device
	cp ./demoCA/keys/bundled.* build/package/device
	cp ./demoCA/keys/ca.crt build/package/device

build_sensor:
	env GOOS=linux GOARCH=arm GOARM=5 go build -o build/package/sensor/sensor github.com/fishjump/cs7ns1_project3/cmd/sensor
	cp ./demoCA/keys/bundled.* build/package/sensor
	cp ./demoCA/keys/ca.crt build/package/sensor