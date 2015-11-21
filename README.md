# do2sshconfig
Command line tool to create an entry in your ~/.ssh/config for each of your DigitalOcean droplets in Go

## How to use
1. Create a Personal Access Token (PAT) in your DigitalOcean account. It's recommended to give it only read-only access.

2. Download a binary for your system from the [Releases](https://github.com/pengux/do2sshconfig/releases) page, then extract and run:
```sh
./do2sshconfig [PAT] >> ~/.ssh/config
```

or better yet, install [Go](https://golang.org/doc/install) and then you can clone and run it:

```sh
git clone https://github.com/pengux/do2sshconfig.git
cd do2sshconfig
go build
./do2sshconfig [PAT] >> ~/.ssh.config
```

### Update
If you want to update the entries with new ones, make a backup first and then you can remove the old entries with this command:
```sh
sed '/# --- DigitalOcean hosts - Start ---/,/# --- DigitalOcean hosts - End ---/d' ~/.ssh/config > /tmp/ssh_config && mv /tmp/ssh_config ~/.ssh/config
```

## Options
```sh
./do2sshconfig -h
Usage of ./do2sshconfig:
  -idFile string
        Path of private SSH key to use
  -ipv6
        Use IPv6 instead of IPv4 as hostnames where possible
  -user string
        Use a different user instead of 'root' (default "root")
```
