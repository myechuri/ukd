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
```
