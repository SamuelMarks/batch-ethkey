batch-ethkey
============

Generate keys in batch.

## Usage

    $ batch-ethkey --help
    Usage of batch-ethkey:
      -dir string
            parent directory containing numbered subdirectories containing keys (default ":/required")
      -inc-port
            Increment port numbers instead of IP addresses
      -n uint
            number subdirectories (containing keys) to create (default 5)
      -network string
            network in CIDR, with start address e.g.: 192.168.0.1/16
      -port-start uint
            port to start counting at (default 12000)

## Example

    $ batch-ethkey -n 10 -dir /my/awesome/dir > /my/awesome/dir/peers.json
    $ tree /my/awesome/dir
    ├── 00
    │   ├── priv_key.pem
    │   └── pub_key.pub
    ├── 01
    │   ├── priv_key.pem
    │   └── pub_key.pub
    ├── 02
    │   ├── priv_key.pem
    │   └── pub_key.pub
    ├── 03
    │   ├── priv_key.pem
    │   └── pub_key.pub
    ├── 04
    │   ├── priv_key.pem
    │   └── pub_key.pub
    ├── 05
    │   ├── priv_key.pem
    │   └── pub_key.pub
    ├── 06
    │   ├── priv_key.pem
    │   └── pub_key.pub
    ├── 07
    │   ├── priv_key.pem
    │   └── pub_key.pub
    ├── 08
    │   ├── priv_key.pem
    │   └── pub_key.pub
    ├── 09
    │   ├── priv_key.pem
    │   └── pub_key.pub
    └── peers.json
    
    10 directories, 21 files

## License

Licensed under either of

- Apache License, Version 2.0 ([LICENSE-APACHE](LICENSE-APACHE) or <https://www.apache.org/licenses/LICENSE-2.0>)
- MIT license ([LICENSE-MIT](LICENSE-MIT) or <https://opensource.org/licenses/MIT>)

at your option.

### Contribution

Unless you explicitly state otherwise, any contribution intentionally submitted
for inclusion in the work by you, as defined in the Apache-2.0 license, shall be
dual licensed as above, without any additional terms or conditions.
