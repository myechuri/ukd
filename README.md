# ukd

Unikernel runtime server on a compute node.


### Get Dependencies
```
go get google.golang.org/grpc
go get github.com/spf13/cobra
go get github.com/urfave/cli
```

### Start ``ukd`` server

```
cd ukd
go install
ukd --help
```

### Start ``ukdctl`` client

```
cd ukdctl
go install
ukdctl --help
```

## Disclaimer

Tested on Ubuntu 16.04. Installation/test instructions on other platforms will differ.


