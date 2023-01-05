# ledger-go

[![CircleCI](https://circleci.com/gh/Zondax/ledger-go.svg?style=shield)](https://circleci.com/gh/ZondaX/ledger-go)

This project provides a library to connect to ledger devices. 

It handles APDU encapsulation, Zemu and USB (HID) communication.

Linux, OSX and Windows are supported.

## Building
```bash
# Build module:
make module
# Build & run tests:
make test
# Build & run example code:
make example && ./example-ledger-go
```
