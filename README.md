batch-ethkey
============

Generate keys in batch.

## Usage

    $ batch-ethkey --help
    Usage of batch-ethkey:
      -dir string
        	parent directory containing numbered subdirectories containing keys (default ":/required")
      -hostname string
        	folder to generate peers.json (default "localhost")
      -n uint
        	number subdirectories (containing keys) to create (default 5)
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
