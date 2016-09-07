pack: clean
	godep save

clean:
	@echo "removing pkgs.." && go clean

build:
	  env GOOS=linux GOARCH=amd64 go build

server:
	fswatch
