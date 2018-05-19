#!/usr/bin/env bash

ip addr add THEIP/255.255.0.0 dev ens7
ip link set ens7 up