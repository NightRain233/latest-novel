default:
	go build
	nohup ./latest-novel &
clean:
	kill `lsof -t -i:4040`