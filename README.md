# ukd

Unikernel runtime server on a compute node.

## Installation

### Prerequisites

[KVM](https://help.ubuntu.com/community/KVM/Installation)

### Install pre-built binaries (recommended)

#### Linux

Get the release binaries, and start ukd:

##### X86-64
As ``root``:
```
curl -O -L https://raw.githubusercontent.com/myechuri/ukd/master/releases/download/v0.1dev/ukd-v0.1dev-linux-x86-64.tar.gz
tar xvzf ukd-v0.1dev-linux-x86-64.tar.gz 
cd ukd-v0.1dev-linux-x86-64
source install-x86-64.sh
```

The last step in the above workflow (``source install-x86-64.sh``) gives you two sample helloworld images to play with:
```
# ls /var/lib/ukd/images/
hello-world-1.img  hello-world-2.img
#
```

##### Aarch64
As ``root``:
```
curl -O -L https://raw.githubusercontent.com/myechuri/ukd/master/releases/download/v0.1dev/ukd-v0.1dev-linux-aarch64.tar.gz
tar xvzf ukd-v0.1dev-linux-aarch64.tar.gz
cd ukd-v0.1dev-linux-aarch64
source install-aarch64.sh
```

The last step in the above workflow (``source install-aarch64.sh``) gives you a sample helloworld image to play with - the image loops forever printing ``Hello World loop`` to console:
```
root@raspberrypi:~# ls /var/lib/ukd/images/
hello-world-loop.img
root@raspberrypi:~#
```

#### Application lifecyle management using ``ukd`` and ``ukonvrt``

Start ``ukd`` server:
```
# ukd
2017/01/26 14:46:11 Detected arch: x86_64 on the system
```

##### ``ukdctl start``

Use ``ukdctl`` client to start sample application:
```
# ukdctl start --name testApp --image-location /var/lib/ukd/images/hello-world-1.img
2017/01/26 14:47:26 Application unikernel started: true, IP: 192.168.122.89, Info: Successful start
```

##### ``ukdctl status``

Monitor status of a provisioned application:
```
# ukdctl status --name testApp
2017/01/26 14:48:18 Application unikernel status check: true, status: RUNNING, Info: IP: 192.168.122.89
```

##### ``ukdctl log``

Gather log output from an application:
```
# ukdctl log --name testApp
2017/01/26 14:49:05 Unikernel application log retrived: true, Info: 
2017/01/26 14:49:05 Unikernel application log:
OSv v0.24-199-g105c5de
eth0: 192.168.122.89
Hello version 1 from C code
#
```

Gather log output from an application that might not be running:
```
# ukdctl log --name unknownApp
2017/01/26 14:49:40 Unikernel application log retrived: true, Info: Application (unknownApp) is currently stopped. No log to report.
2017/01/26 14:49:40 Unikernel application log:
#
```

##### ``ukdctl stop``

Stop a running application:
```
# ukdctl stop --name testApp
2017/01/26 14:50:38 Application unikernel stopped: true, Info: Successfully stopped Application (testApp)
#
```

##### ``ukdctl update-image``

Update an application's on-disk image:
```
# ukdctl update-image --oldImage /var/lib/ukd/images/hello-world-1.img --newImage /var/lib/ukd/images/hello-world-2.img 
2017/01/26 14:52:16 Computed new image signature
2017/01/26 14:52:16 Gathered signature of old image on destination
2017/01/26 14:52:16 Calcuated diff of new image over old image: 1044KB
2017/01/26 14:52:16 Transmitting diff over..
2017/01/26 14:52:16 Unikernel image update: true, new image path on destination: /tmp/ukd-update-stage-115735916/newImage.img, Info: Verified signature match for new Image on source and destination
#
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


