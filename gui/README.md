# message-client

```bash
go get -d https://github.com/akavel/rsrc
go install https://github.com/akavel/rsrc
rsrc -manifest main.manifest -o rsrc.syso

go build -ldflags="-H windowsgui"
```

The usual default message loop includes calls to win32 API functions, which incurs a decent amount of runtime overhead coming from Go. As an alternative to this, you may compile Walk using an optional C implementation of the main message loop, by passing the walk_use_cgo build tag:

```bash
go build -tags walk_use_cgo
```

## debug

```bash
go get -d github.com/go-delve/delve/cmd/dlv
go install github.com/go-delve/delve/cmd/dlv
```
