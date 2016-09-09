# ukd

Unikernel runtime server on a compute node.

## Install pre-built binaries

#### Linux

Get the release binaries, and start ukd:
```
curl -L https://github.com/myechuri/ukd/releases/download/v0.01/ukd-v0.01-linux-amd64.tar.gz
tar xvzf ukd-v0.01-linux-amd64.tar.gz
mkdir -p /var/lib/ukd/images
./ukd &
```

## Build 

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


