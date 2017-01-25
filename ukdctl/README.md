# CLI client for ``ukd``

## Install dependecies

```
go get github.com/spf13/cobra
```

## Install ukdctl

```
go install
```

## Test

```
# ukdctl version
2016/08/12 14:28:01 Ukd server version: 0.1

# ukdctl --server-endpoint 'localhost:55555' version
2016/08/12 14:27:13 Ukd server version: 0.1

# ukdctl start --name testapp --image-location "/root/osv/build/last/usr.img"
2016/08/25 14:13:06 Application unikernel started: true, IP: 192.168.122.89, Info: Successful start

# ukdctl start --name testapp --image-location "/root/osv/build/last/usr.img"
2016/08/25 14:13:10 Application unikernel started: false, IP: , Info: testapp is already running. Please choose a different name for the application if you wish to start a second instance using the same image.

# ukdctl stop --name testapp
2016/08/25 12:14:41 Application unikernel stopped: true, Info: Successfully stopped Application

# ukdctl stop --name testapp
2016/08/25 13:58:33 Application unikernel stopped: true, Info: App not found. Nothing to do.

# ukdctl start --name testapp --image-location "/invalid/path/usr.img"
2016/08/25 14:08:35 Application unikernel started: false, IP: , Info: /invalid/path/usr.img does not exist, error: stat /invalid/path/usr.img: no such file or directory

# ukdctl start --name testapp --image-location "/root/osv/build/last/usr.img" --hypervisor invalid
2016/08/25 15:26:11 Application unikernel started: false, IP: , Info: Requested hypervisor (invalid) is not yet supported.

# ukdctl update-image --oldImage "/destination-host-path/images/java/old.img" --newImage "/source-host-path/images/java/new.img"
2016/12/05 17:33:56 Computed new image signature
2016/12/05 17:33:56 Gathered signature of old image on destination
2016/12/05 17:33:56 Calcuated diff of new image over old image: 2286KB
2016/12/05 17:33:56 Transmitting diff over..
2016/12/05 17:33:57 Unikernel image update: true, new image path on destination: /tmp/ukd-update-stage-227435982/newImage.img, Info: Verified signature match for new Image on source and destination

# ukdctl status --name testapp
2016/12/23 10:38:31 Application unikernel status check: true, status: RUNNING, Info: IP: 10.0.2.15

# ukdctl log --name test1
2017/01/25 10:12:08 Unikernel application log retrived: true, Info: 
2017/01/25 10:12:08 Unikernel application log:
OSv v0.24-174-gd82e1b3
eth0: 192.168.122.89

# ukdctl log --name test2
2017/01/25 10:12:27 Unikernel application log retrived: true, Info: Application (test2) is currently stopped. No log to report.
2017/01/25 10:12:27 Unikernel application log:

```
