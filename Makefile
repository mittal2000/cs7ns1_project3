all: build_gateway

build_gateway:
	env GOOS=linux GOARCH=arm GOARM=5 go build -o build/gw/gateway github.com/fishjump/cs7ns1_project3/cmd/gateway
	cp ./demoCA/keys/* build/gw