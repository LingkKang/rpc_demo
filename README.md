# A simple RPC demo

Remote Procedure Call (RPC)

## Server

### Running

The server is already hosted on `test.lingkang.dev:8333`, but you can also run it locally:

``` bash
cd server
go run cmd/main.go
```

Or build and run the binary:

``` bash
cd server
go build -o target/rpc_demo cmd/main.go
nohup ./target/rpc_demo &
```

Use `nohup ./target/rpc_demo >/dev/null 2>&1 &` to discard the output.

### Output

Golang built-in module `log` is used to provide basic running information.

![The running output of the server side.](./img/server_running.png)

---
