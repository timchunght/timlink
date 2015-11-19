# Timlink
	
Development:

	go get
	gin -p 5000

Build and Deploy:

	go build

To run the binary file, run:

	export PORT=8080 && ./timlink

For the app to work correctly, please ensure you fill out the ``.env`` with the correct environment variable. By default, it is

	MONGO_URL=mongodb://localhost:27017

You can change the port to the one your prefer.

###Usage:

Welcome! This is the simpliest url shortener API

POST   /shorten?url=your_url

Sample Response:

	{
	  "id": "564d422037f69f505565dc64",
	  "url": "https://www.google.com",
	  "hash": "dayYGs",
	  "short_url": "localhost:9000/dayYGs",
	  "created_at": "2015-11-18T22:29:36.900194106-05:00"
	}
		