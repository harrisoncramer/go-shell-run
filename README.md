# ShellRun

This repository is a simple Golang HTTP server designed to allow for the execution of shell commands remotely.

It's original use-case was automating deployment workflows via webhooks.

# Installation and Use

Compile the binary with `go build` for your operating system of choice. Then send the binary to your server:

```terminal
$ GOOS=linux GOARCH=amd64 go build .
$ scp ./golang-webhook ubuntu@12.345.67.89:/home/harrison
```

Run the binary on an open port of your choosing and pass it your API key:

```terminal
$ ssh ubuntu@12.345.67.89
$ ./golang-webhook --port=3012 --token=abc
```

You can now pass arbitrary shell commands to the server by POSTing them to the `/jobs` endpoint.

```
$ curl --location --request POST 'http://12.345.67.89:3012/status' \
--header 'token: abc' \
--header 'Content-Type: application/json' \
--data-raw '{
    "jobs": [
        "pwd",
        "ls -la",
    ]
}'
```

If the server is already executing the jobs it'll return a 503, if not, you'll
recieve a JSON response.
