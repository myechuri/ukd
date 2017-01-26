# ukd

Unikernel runtime server on a compute node.

## Installation

### Prerequisites

[KVM](https://help.ubuntu.com/community/KVM/Installation)

### Install pre-built binaries (recommended)

#### Linux

Get the release binaries, and start ukd:

X86-64
```
curl -L https://github.com/myechuri/ukd/releases/download/v0.1dev/ukd-v0.1dev-linux-x86-64.tar.gz 
tar xvzf ukd-v0.1dev-linux-x86-64.tar.gz 
cd ukd-v0.1dev-linux-x86-64
source install-x86-64.sh
```

Aarch64
```
curl -L https://github.com/myechuri/ukd/releases/download/v0.1dev/ukd-v0.1dev-linux-aarch64.tar.gz
tar xvzf ukd-v0.1dev-linux-aarch64.tar.gz
cd ukd-v0.1dev-linux-aarch64.tar.gz
source install-aarch64.sh
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

### Build from source

#### Get Dependencies
```
go get google.golang.org/grpc
go get github.com/spf13/cobra
go get github.com/urfave/cli
go get github.com/satori/go.uuid
```

#### Build and install ``ukd`` server

```
cd ukd
go install
```

#### Build and install ``ukdctl`` client

```
cd ukdctl
go install
```

## Supported platforms

### Hypervisor

[Ubuntu 16.04 + KVM](http://releases.ubuntu.com/16.04/)

### Monitor

[Qemu](http://wiki.qemu.org/Main_Page)

### Unikernel framework

[OSv](http://osv.io/)

## Disclaimer

Tested on Ubuntu 16.04. Installation/test instructions on other platforms will differ.


