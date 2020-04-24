# cb-operator
The Operation Tool for Cloud-Barista System Runtime

# Prerequisite

## install docker

## install docker-compose
```
$ sudo curl -L https://github.com/docker/compose/releases/download/1.25.5/docker-compose-$(uname -s)-$(uname -m) -o /usr/local/bin/docker-compose
$ sudo chmod +x /usr/local/bin/docker-compose
$ sudo docker-compose --version
```

# Command to build the operator from souce code
cloud-barista/cb-operator/src$ go build -o operator

# Commands to use the operator

## Help
```
~/go/src/github.com/cloud-barista/cb-operator/src$ ./operator 

The operator is a tool to operate Cloud-Barista system. 
  
  For example, you can setup and run, stop, and ... Cloud-Barista runtimes.
  
  - operator run -f ../docker-compose.yaml
  - operator info
  - operator exec -t cb-tumblebug -c "ls -al"
  - operator stop -f ../docker-compose.yaml
  - operator remove -f ../docker-compose.yaml -v -i

Usage:
  operator [command]

Available Commands:
  exec        Run commands in a target component of Cloud-Barista System
  help        Help about any command
  info        Get informtion of Cloud-Barista System
  remove      Stop and Remove Cloud-Barista System
  run         Setup and Run Cloud-Barista System
  stop        Stop Cloud-Barista System

Flags:
      --config string   config file (default is $HOME/.operator.yaml)
  -h, --help            help for operator
  -t, --toggle          Help message for toggle

Use "operator [command] --help" for more information about a command.
```

## Run
```
~/go/src/github.com/cloud-barista/cb-operator/src$ ./operator run -h

Setup and Run Cloud-Barista System

Usage:
  operator run [flags]

Flags:
  -f, --file string   Path to Cloud-Barista Docker-compose file (default "*.yaml")
  -h, --help          help for run

Global Flags:
      --config string   config file (default is $HOME/.operator.yaml)
```

## Stop
```
~/go/src/github.com/cloud-barista/cb-operator/src$ ./operator stop -h

Stop Cloud-Barista System

Usage:
  operator stop [flags]

Flags:
  -f, --file string   Path to Cloud-Barista Docker-compose file (default "*.yaml")
  -h, --help          help for stop

Global Flags:
      --config string   config file (default is $HOME/.operator.yaml)
```

## Exec
```
~/go/src/github.com/cloud-barista/cb-operator/src$ ./operator exec -h

Run commands in your components of Cloud-Barista System. 
	For instance, you can get an interactive prompt of cb-tumblebug by
	[operator exec cb-tumblebug sh]

Usage:
  operator exec [flags]

Flags:
  -c, --command string   Command to excute
  -h, --help             help for exec
  -t, --target string    Name of CB component to command

Global Flags:
      --config string   config file (default is $HOME/.operator.yaml)

```
