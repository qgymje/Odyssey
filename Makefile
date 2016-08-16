pack: clean
	godep save

clean:
	@echo "removing pkgs.." && go clean

server:
	fswatch
