#!/bin/bash

user=pi
server=192.168.1.233
declare workers=("192.168.1.234" "192.168.1.235")

ssh ${user}@${server} << EOF
    curl -sfL https://get.k3s.io | sh -
EOF

for w in "${workers[@]}" do

done