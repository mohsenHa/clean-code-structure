# Clean Code Structure

Clean code structure for use in other projects.

## Directories:

```
.
├── README.md
├── cmd
│   ├── httpserver
│   │   └── main.go
│   ├── manager
│   │   ├── all.go
│   │   ├── blueprint
│   │   │   ├── handler
│   │   │   │    ├── handler.tmpl
│   │   │   │    ├── route.tmpl
│   │   │   │    └── sample.tmpl
│   │   │   ├── param
│   │   │   │    └── sample.tmpl
│   │   │   ├── service
│   │   │   │    ├── sample.tmpl
│   │   │   │    └── service.tmpl
│   │   │   └── validator
│   │   │       ├── sample.tmpl
│   │   │       └── validator.tmpl
│   │   ├── creator.go
│   │   ├── handler.go
│   │   ├── main.go
│   │   ├── param.go
│   │   ├── service.go
│   │   └── validator.go
│   └── scheduler
│       └── main.go
├── config
│   ├── config.go
│   ├── default.go
│   └── load.go
├── delivery
│   └── httpserver
│       ├── healthhandler
│       │   ├── check.go
│       │   ├── handler.go
│       │   └── route.go
│       ├── middleware
│       └── server.go
├── logger
│   └── logger.go
├── param
│   ├── baseparam.go
│   └── healthparam
│       └── check.go
├── pkg
│   ├── errmsg
│   │   └── messages.go
│   ├── httpmsg
│   │   └── mapper.go
│   └── richerror
│       └── richerror.go
├── scheduler
│   ├── doNothing.go
│   └── scheduler.go
├── service
│   └── healthservice
│       ├── check.go
│       └── service.go
├── validator
│    └── healthvalidator
│        ├── check.go
│        └── validator.go
├── adapter
├── contract
├── entity
├── repository
├── go.mod
├── go.sum
└── config.yml
```
## Directories specification

### cmd
This directory contains all binaries we can build from project

#### httpserver
This directory contains the main binary of project

The main.go is contain all services, validators and ... . All thing must configure in this file.

#### manager
This folder provide a dev tools help to create everything easier by command line interface. Use it as mentioned in bellow:

```shell
go run ./cmd/manager/
```

#### scheduler
Alle scheduler task and worker must register here.

### config
All configuration must add to the config file if default value is required define it in `default.go`

### delivery
This folder contains all delivery layers like httpserver, grpcserver or ...

#### httpserver
All http handler goes here every handler in one directory and contain `route.go` and `handler.go`. Every methods of 
handler in separate files like get,store,edite or ... .

### logger
This folder contain the logger of project.

### param
This folder contain the struct of all parameters. every handler hase a folder here named `handlerparam`.

### pkg
All packages must write to this folder

### scheduler
All scheduler must write to this folder.

### service
This folder is used for write business logics.

### validator
All request validator must write to this folder

### adapter
All third party services go here.

### contract
All contracts used for data transfer protocols like grpc or rabbitmq or ... goes here.

### entity
All entities layer must write in this folder

### repository
All DB calls must manage in this folder

### config.yml
All configuration must write in this file.



