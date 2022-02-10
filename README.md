# Introduction
A simple Go app that acts as a multicast listener and a sender

# How to Use?
After `go build` this, run:
```
$ ./multicaster server
```

In another terminal session, running in the same subnet, run:
```
$ ./multicaster client
```
