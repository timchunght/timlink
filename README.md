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