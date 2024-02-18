# MiniTFTPd
*A tiny read-only TFTP server*

This is a tiny read-only zero-configuration utility TFTP server.
It always serves the contents of the directory it was called in.

## Installation
This requires an up-to-date Go toolchain, currently there are no precompiled binaries provided.

```sh
go install git.dolansoft.org/lorenz/minitftpd@latest
```

Binding to the TFTP server port (69) requires special privileges on Unix-style
systems like Linux and MacOS older than 10.14. 

On Linux you can chose between running minitftpd as root, granting CAP_NET_BIND_SERVICE
to the binary (`sudo setcap CAP_NET_BIND_SERVICE=+eip /path/to/minitftpd`) or disabling
the privileged ports feature on the whole machine/network namespace
(`sysctl -w net.ipv4.ip_unprivileged_port_start=0`).

On older versions of MacOS (before 10.14) one needs to run minitftpd as root, more recent
versions do no longer enforce privileged ports.

## Usage

Just start minitftpd in the directory you're trying to serve. That's it.
