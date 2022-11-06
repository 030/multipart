#!/bin/bash -e

curl https://gist.githubusercontent.com/030/620a95e7a699c4db3e76b2b8b0309909/raw/cf904c84ee40f34040a7d13a8a51ef5347674f52/go-build.sh -o go-build.sh
chmod +x go-build.sh
./go-build.sh multipart