deploy:
	GOOS=linux GOARCH=amd64 go build -o build/server \
	 && docker build -t rest-api . \
	 && docker run -p 8000:8000 rest-api 
