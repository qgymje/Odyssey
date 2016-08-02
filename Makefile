pack: clean
	cd ./servers/api_services && godep save

clean:
	go clean

server:
	fswatch
