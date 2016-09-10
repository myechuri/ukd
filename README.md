# ukd

Unikernel runtime server on a compute node.

## Install pre-built binaries

#### Linux

Get the release binaries, and start ukd:
```
curl -L https://github.com/myechuri/ukd/releases/download/v0.1/ukd-v0.1-linux-amd64.tar.gz
tar xvzf ukd-v0.1-linux-amd64.tar.gz
source install.sh
```

Start ``ukd`` server:
```
ukd &
```

Use ``ukdctl`` client:
```
ukdctl --help
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


