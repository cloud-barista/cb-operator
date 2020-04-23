# cb-operator
The Operator for Cloud-Barista system

# install docker-compose
$ sudo curl -L https://github.com/docker/compose/releases/download/1.25.5/docker-compose-$(uname -s)-$(uname -m) -o /usr/local/bin/docker-compose

$ sudo chmod +x /usr/local/bin/docker-compose

$ sudo docker-compose --version

$ sudo docker-compose up

# Command to build the operator from souce
cloud-barista/cb-operator/src$ go build -o operator

# Command to use the operator

## Help
~/go/src/github.com/cloud-barista/cb-operator/src$ ./operator -h

A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

Usage:
  operator [command]

Available Commands:
  exec        A brief description of your command
  help        Help about any command
  info        A brief description of your command
  remove      A brief description of your command
  run         Setup and Run
  stop        Stop

Flags:
      --config string   config file (default is $HOME/.operator.yaml)
  -h, --help            help for operator
  -t, --toggle          Help message for toggle

Use "operator [command] --help" for more information about a command.

## Run
~/go/src/github.com/cloud-barista/cb-operator/src$ ./operator run -h
Setup and Run the Cloud-Barista System

Usage:
  operator run [flags]

Flags:
  -f, --file string   Path to Cloud-Barista Docker-compose file (default "*.yaml")
  -h, --help          help for run

Global Flags:
      --config string   config file (default is $HOME/.operator.yaml)

## Stop
~/go/src/github.com/cloud-barista/cb-operator/src$ ./operator stop -h
Stop the Cloud-Barista System

Usage:
  operator stop [flags]

Flags:
  -h, --help   help for stop

Global Flags:
      --config string   config file (default is $HOME/.operator.yaml)

