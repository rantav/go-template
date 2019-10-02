<!-- BEGIN __DO_NOT_INCLUDE__ -->
# Welcome to go-template

[![Actions Status](https://github.com/rantav/go-template/workflows/Go/badge.svg)](https://github.com/rantav/go-template/actions)

This project will jumpstart your Golang project and provide a set of tempaltes you may use to keep your code tidy and consistent.

## TLDR
Run this snippet:

    cd your-work-dir
    bash <(curl https://raw.githubusercontent.com/rantav/go-template/master/scripts/init.sh) rantav my-awesome-go-project

    # replace rantav with your github user or github group
    # replace my-awesome-go-project with your project or service name

Here's how it looks like in action:

[![asciicast](https://asciinema.org/a/oHvy0RWx80bexcFE1nJyREGyk.svg)](https://asciinema.org/a/oHvy0RWx80bexcFE1nJyREGyk)

## What does the script go-template.sh do?
This script installs the `go-archetype` tool and `protoc` and uses the `transformations.yml` file in order to initialize a working Golang project on your local disk based on the template project `go-template`.
The script also initializes an empty git repo and adds the origin specific by your username/group and project-name.

This tool then asks for your input in order to determine the contents and structure of your project.

## What's in the box?
The generated project results in:

* `Makefile`
* `README.md` (which you have to edit to add contents)
* `admin` package, which includes healthchecks, basic support for gRPC tracing and metrics
  * Admin http pages at `/_/health.json` and `/_/metrics.json`
  * A first healthcheck (checking gRPC to our local service)
* `cmd` package which includes CLI args support
  * Two CLI args: `admin-bind-address` and `grpc-bind-address`
* Go modules files (`go.mod` and a generaged `go.sum`)
* `grpc` package with an `idl` file, a `client` and a `server` (which you'll have to edit in order to implement your business logic)
  * gRPC instrumentation interface at `/_/zpages/rpcz`, `/_/zpages/tracez` and `/_/channelz`
  * grpc already has a gRPC healthcheck
* git precommit hood that runs lint.
* *Tracing*, using Jaeger.
* and some more...


When creating the project you are asked whether you'd like to include gRPC support. This of course results in different code.

### gRPC flavour
Here's the source tree when choosing to use gRPC:

```
➜  my-awesome-go-project git:(master) ✗ tree
.
├── Makefile
├── README.md
├── admin
│   ├── grpc-healthcheck.go
│   ├── grpc-healthcheckt_test.go
│   ├── healthcheck.go
│   ├── healthcheck_test.go
│   ├── package.go
│   ├── server.go
│   └── tracing.go
├── assets
│   └── README.md
├── cmd
│   ├── flags.go
│   ├── root.go
│   ├── serve.go
│   ├── serve_test.go
│   ├── test-grpc-client.go
│   └── test-grpc-client_test.go
├── docs
│   └── README.md
├── examples
│   └── README.md
├── go.mod
├── go.sum
├── grpc
│   ├── client
│   │   └── grpc-client.go
│   ├── idl
│   │   └── service.proto
│   └── server
│       ├── grpc-server.go
│       └── grpc-server_test.go
├── internal
│   └── generated
│       └── grpc
│           └── service.pb.go
├── main.go
├── pkg
│   └── README.md
├── scripts
│   └── README.md
├── service
│   └── service.go
├── test
│   └── README.md
├── tools
│   └── README.md
├── types
│   └── types.go
└── web
    └── README.md
```

### Non gRPC flavour
Here's the source tree when choosing not to use gRPC:

```
➜  my-awesome-go-project git:(master) ✗ tree
.
├── Makefile
├── README.md
├── admin
│   ├── healthcheck.go
│   ├── healthcheck_test.go
│   ├── package.go
│   └── server.go
├── assets
│   └── README.md
├── cmd
│   ├── flags.go
│   ├── root.go
│   ├── serve.go
│   └── serve_test.go
├── docs
│   └── README.md
├── examples
│   └── README.md
├── go.mod
├── go.sum
├── main.go
├── pkg
│   └── README.md
├── scripts
│   └── README.md
├── service
│   └── service.go
├── test
│   └── README.md
├── tools
│   └── README.md
├── types
│   └── types.go
└── web
    └── README.md
```

## In Depth
We use a tool called [go-archetype](https://github.com/rantav/go-archetype) for templating.
This tool is used in order to transition a blueprint codebase to a project instance. The user provides the required input and thereafter the tool transforms the blueprint project into your own personal project.


### Tools
We use and install several tools, namely `go-archetype`, `protoc` and some related tools.
We install them all in `/tmp` and they are versioned.
In the future, if prooved useful, we might install them in other location outside of `tmp`


---
The following part is going to be part of the actual generated project
---

<!-- This section will be generated for the new project -->
<!-- END __DO_NOT_INCLUDE__ -->

# Welcome to go-template

A template project

## Purpose
TODO: Explain what this service does

## High level design
TODO: Explain the design

Use http://asciiflow.com/ to show your diagrams

```
+---------+           +---------+
| Service |           | Service |
|    A    |           |    B    |
|         +---------->+         |
|         |           |         |
+---------+           +---------+
```

## How to use
TODO

### Debugging the service.
* Use `AF_LOG_LEVEL` to increase log level (levels are: panic, fatal, error, warning, info, debug, trace)
* Use `GRPC_GO_LOG_VERBOSITY_LEVEL=99` and `GRPC_GO_LOG_SEVERITY_LEVEL=info` to increase logging verbosity of gRPC
* Access `http://localhost:8080/_/health.json` for healthchecks
* Access `http://localhost:8080/_/metrics.json` for healthchecks
* Access `http://localhost:8080/_/zpages/rpcz` for gRPC stats
* Access `http://localhost:8080/_/zpages/tracez` for gRPC tracing
* Access `http://localhost:8080/_/channelz` for gRPC channel UI
