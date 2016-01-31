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

### Dependency Vendoring

We are currently using ``godep`` to manage dependencies. All dependencies are tracked in ``Godeps.json`` and copied into the ``Godeps`` directory. If the project directory is changed, say from ``tim`` to ``timothy``, run the following. However, if any of the source files are referencing same-project package(s), you need to change those import paths first.

For example, I have a package in the same project named ``connection`` and my old project name is ``tim``, I will have to reference ``connection`` pacakge in my ``main.go`` as ``tim/connection``. Now the directory name is ``timothy`` and I will have to change it to ``timothy/connection`` and then run the following,

	godep save -r ./...

``godep`` is not yet smart enough to distinguish whether the package belongs to the same project or is an external dependency.

Go Debugging
---

If you have used Rails, you will miss using binding.pry. However, Go does have similar. Go has a package called "Godebug"

Install Godebug by running:

	go get github.com/mailgun/godebug

Insert a breakpoint anywhere in a source file you want to debug:

	_ = "breakpoint"

Replace pkgs with the packages we are will be debugging if it is not the ``main`` package

	godebug build -instrument <pkgs>

For example:

	godebug build -instrument planit/connection

godebug will generate a binary named ``yourprojectname.debug``, run that binary with the necessary arguments or environment variables and use it as you would binding.pry

For example,

	PORT=8080 ./planit.debug
