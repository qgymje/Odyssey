pack: clean
	godep save

clean:
	go clean

serve:
	fswatch
