# go-multiaddr-vsock

> [multiaddr](https://github.com/multiformats/multiaddr) support for VSOCK

Adds support to `multiaddr` addresses such as `/vsock/3/tcp/55555` to allow `virtio-vsock` as communication channel.

## Install

```sh
go get github.com/balena/go-multiaddr-vsock
```

## Usage

### Example

#### Simple

```go
import ma "github.com/balena/go-multiaddr-vsock"

// construct from a string (err signals parse failure)
m1, err := ma.NewMultiaddr("/vsock/3/tcp/1234")

// construct from bytes (err signals parse failure)
m2, err := ma.NewMultiaddrBytes(m1.Bytes())

// true
strings.Equal(m1.String(), "/vsock/3/tcp/1234")
strings.Equal(m1.String(), m2.String())
bytes.Equal(m1.Bytes(), m2.Bytes())
m1.Equal(m2)
m2.Equal(m1)
```

#### Protocols

```go
// get the multiaddr protocol description objects
m1.Protocols()
// []Protocol{
//   Protocol{ Code: 40, Name: 'vsock'},
// }
```

## Contribute

Check out our [contributing document](https://github.com/multiformats/multiformats/blob/master/contributing.md) for more information on how we work, and about contributing in general. Please be aware that all interactions related to multiformats are subject to the IPFS [Code of Conduct](https://github.com/ipfs/community/blob/master/code-of-conduct.md).

Small note: If editing the README, please conform to the [standard-readme](https://github.com/RichardLitt/standard-readme) specification.

## License

[MIT](LICENSE) © 2023 Guilherme Versiani
