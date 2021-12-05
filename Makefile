build:
	go build -o templit 

install: build
	mv templit /usr/local/bin

