# UHash

Universal hashing utility for files and simple inputs written in Go

## Usage

Hash text

```
uhash sha1 apple
```

Hash file contents

```
uhash md5 < cat.jpg
cat cat.jpg | uhash md5
```

### Supported algorithms

- md4
- md5
- sha1
- sha224
- sha256
- sha384
- sha512
- sha3-224
- sha3-256
- sha3-384
- sha3-512
- whirlpool
- ripemd-160

## Development

The application is written in Go and compiled into machine code for all systems. The tests are written in Python and execute the terminal application and look for output.

Build development version

```
make build
```

Run tests

```
make test
```
