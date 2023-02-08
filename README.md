# go socket connection

## Installation

Before start you need to clone this repository:

```shell
git clone git@github.com:Parviz-Makhkamov/go-socket.git
```

## Usage

For start, in the root directory of the cloned repository:

First, start server and create socket:
```golang
cd ./server
go run .  // stated server with default socket path, default path: "/tmp/echo.sock"
or
go run . -socketAddr= socket path // stated server with special socket path
```

From client side:
```golang
cd ./client
go run .  // if server started with default socket path
or
go run . -socketAddr= socket path // if server started with special socket path
```
