#!/bin/bash

dnf intall -y wget
wget https://go.dev/dl/go1.21.3.linux-amd64.tar.gz
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.21.3.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
go version

echo "export go path to /root/.bashrc -> export PATH=\$PATH:/usr/local/go/bin"