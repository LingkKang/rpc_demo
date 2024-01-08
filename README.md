# A simple RPC demo

Remote Procedure Call (RPC)

- [1. Server](#1-server)
  - [1.1. Run](#11-run)
- [2. Client](#2-client)
  - [2.1. Run](#21-run)

## 1. Server

The server side is implemented in Go, under the `./server` directory.

### 1.1. Run

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

---

## 2. Client

The client side is implemented in Rust, under the `./client` directory.

### 2.1. Run

It can be run locally:

``` bash
cd client
cargo run
```
