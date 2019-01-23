#!/bin/sh

echo "export GOPATH=$HOME/go" >> ~/.bashrc;
echo "export PATH=$PATH:$HOME/go/bin:/usr/local/go/bin" >> ~/.bashrc;
source ~/.bashrc

exec "$@";