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

- sha1
- md5
- sha224
- sha256
- sha384
- sha512
- whirlpool

## Development

The application is written in go and compiled to machine code for all systems. The tests are written in python and execute the terminal application and look for output.

Build development version

```
make build
```

Run tests

```
make test
```
