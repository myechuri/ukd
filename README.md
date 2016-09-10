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

Use ``ukdctl`` client to start/stop sample application:
```
root@ubuntu# ukdctl --help
root@ubuntu# ukdctl start --image-location="/var/lib/ukd/images/nativeexample.img" --name testuk
2016/09/09 17:21:17 Application unikernel started: true, IP: 192.168.122.89, Info: Successful start
root@ubuntu# ukdctl stop --name testuk
2016/09/09 17:21:22 Application unikernel stopped: true, Info: Successfully stopped Application (testuk)
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


