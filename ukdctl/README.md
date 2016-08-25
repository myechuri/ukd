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

# ukdctl startUK --name testapp --image-location "/root/osv/build/last/usr.img"
2016/08/25 12:14:28 Application unikernel started: true, IP: 192.168.122.89, Info: Successful start

# ukdctl stopUK --name testapp
2016/08/25 12:14:41 Application unikernel stopped: true, Info: Successfully stopped Application
#
```
