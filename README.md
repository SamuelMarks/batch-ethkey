batch-ethkey
============

Generate keys in batch.

## Usage

    $ batch-ethkey --help
    Usage of batch-ethkey:
      -d string
            parent directory containing numbered subdirectories containing keys (default ":/required")
      -n uint
            number subdirectories (containing keys) to create (default 5)

## Example

    $ batch_ethkey -n 10 -dir /my/awesome/dir
    $ tree
    .
    ├── 0
    │   └── priv_key.pem
    ├── 1
    │   └── priv_key.pem
    ├── 2
    │   └── priv_key.pem
    ├── 3
    │   └── priv_key.pem
    ├── 4
    │   └── priv_key.pem
    ├── 5
    │   └── priv_key.pem
    ├── 6
    │   └── priv_key.pem
    ├── 7
    │   └── priv_key.pem
    ├── 8
    │   └── priv_key.pem
    └── 9
        └── priv_key.pem
    
    10 directories, 10 files
