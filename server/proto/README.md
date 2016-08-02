
## Dev

Note: Below instructions are for Ubuntu 16.04. Please modify for other platforms accordingly.

If you make changes to ``ukd.proto``, please regenerate ``ukd.pb.go`` by following below steps.

1. Download ``protoc`` version 3.0 from [here](https://github.com/google/protobuf/releases).

2. Place ``protoc`` binary in your ``PATH``.

3. Install ``protoc-gen-go``.

```
apt install golang-goprotobuf-dev
```

4. Run ``protoc`` to generate ``ukd.pb.go``.

```
protoc -I=$PWD --go_out=plugins=grpc:.  $PWD/ukd.proto
```
