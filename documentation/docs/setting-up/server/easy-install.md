# Easy-install script

!!! caution alert alert-warning "Caution"
    You can download and check `get-pmm.sh` before running it from our [github]:

## Linux or macOS
Using `curl`:
```sh
curl -fsSL https://www.percona.com/get/pmm2 | /bin/bash
```

Using `wget`: 
```sh
wget -O - https://www.percona.com/get/pmm2 | /bin/bash
```

This script:

- Installs Docker if it is not already installed on your system.
- Stops and renames any currently running PMM Server Docker container from `pmm-server` to `pmm-server-{timestamp}`. This old pmm-server container is not a recoverable backup.
- Pulls and runs the latest PMM Server Docker image.
- Can run in _Interactive mode_ to change the default settings:
```sh
curl -fsSLO https://www.percona.com/get/pmm2 # (or wget https://www.percona.com/get/pmm2)
chmod +x pmm2
./pmm2 --interactive
```


[github]: https://github.com/percona/pmm/blob/main/get-pmm.sh
