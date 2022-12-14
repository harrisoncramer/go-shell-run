# GoShellRun 🐚

This repository is a simple Golang HTTP server designed to allow for the remote
execution of shell commands.

It was originally designed to automate deployment workflows via webhooks. 

NOTE: This is an insecure approach because the token is not encrypted, and is mainly for demonstration purposes. If you need something like this look into just running your commands with SSH or via Puppet if you need something more robust.

## Security

In order to validate an inbound request, this server checks that a specific header called `token` 
matches what you set in your server configuration. You may also want to firewall the port you
choose only for specfic IP addresses, since this server enables shell access to your
machine. Be careful!

## Installation and Use

Compile the binary with `go build` for your operating system of choice. Then send the binary to your server:

```terminal
$ GOOS=linux GOARCH=amd64 go build .
$ scp ./go-shell-run ubuntu@12.345.67.89:/home/harrison
```

Run the binary on an open port of your choosing and pass it your API key:

```terminal
$ ssh ubuntu@12.345.67.89
$ ./go-shell-run --port=3012 --token=89fnoq8yeho8h1y3o
```

You can now pass arbitrary shell commands to the server by POSTing them to the `/jobs` endpoint.

```
$ curl --location --request POST 'http://12.345.67.89:3012/jobs' \
--header 'token: 89fnoq8yeho8h1y3o' \
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

## Running as a Service

The `/resources` folder contains a service manifest that you could use to turn this
binary into a service controlled by `systemctl`, so that the server is
persistent even across command failures.

Copy this file to your `/etc/systemd/system` directory and restart your system
manager. On Ubuntu, this series of commands might look something like this:

```terminal
$ scp ./resources/go-shell-run.service ubuntu@12.345.67.89:/home/harrison
$ ssh ubuntu@12.345.67.89
$ sudo mv /home/harrison/go-shell-run.service /etc/systemd/system
$ sudo systemctl daemon reload
$ sudo systemctl status go-shell-run.service
● go-shell-run.service - Service that runs an HTTP server meant for arbitrary shell execution remotely.
     Loaded: loaded (/etc/systemd/system/go-shell-run.service; disabled; vendor preset: enabled)
     Active: active (running) since Sun 2022-08-07 17:20:42 UTC; 10min ago
   Main PID: 57135 (go-shell-run)
      Tasks: 3 (limit: 2354)
     Memory: 852.0K
        CPU: 5ms
     CGroup: /system.slice/go-shell-run.service
             └─57135 /home/harrison/c2c-visualization/go-shell-run --port=3012 --token=s72870h!b98f0uA(
```

You can follow the logs with `journalctl` like this:

```terminal
sudo journalctl -u go-shell-run.service --follow
```
