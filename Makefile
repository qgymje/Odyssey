pack: clean
	godep save

clean:
	go clean

server:
	fswatch
