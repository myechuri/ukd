## Dev steps

#### Go get cli

```
go get github.com/urfave/cli
```

#### Install ``ukd`` in ``$GOPATH/bin``

```
go install
```

#### Start ``ukd``

```
ukd --help
```

#### Troubleshooting

##### Resource busy

Problem: ``StartUK`` request to ``ukd`` fails:
```
2016/08/23 13:00:10 ioctl(KVM_CREATE_VM) failed: 16 Device or resource busy
failed to initialize KVM: Device or resource busy
```

Solution: Check if any VirtualBox VMs are up, and if so, stop them:

```
# vagrant destroy -f node
```
