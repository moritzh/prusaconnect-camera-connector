build: $(wildcard **/*.go)
	go build 

install: build 
	cp prusaconnect-camera-connector /usr/local/bin