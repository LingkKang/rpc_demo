# A simple RPC demo

Remote Procedure Call (RPC)

## Usage

### Server

The server is already hosted on `test.lingkang.dev:8333`, but you can also run it locally:

```bash
go run server/main.go
```

Or build and run the binary:

```bash
go build -o server/target/rpc_demo server/main.go
nohup ./server/target/rpc_demo &
```

Use `nohup ./server/target/rpc_demo >/dev/null 2>&1 &` to discard the output.

### Client

```bash
cd client
cargo run
```

#### Output

A naive implementation of logger based on `log` crate is used in this demo:

![The running output of the client side](./img/client_running.png)
