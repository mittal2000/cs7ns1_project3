# Gateway User Case
# Anoushka Mittal
## Overview
1. listen remote request on rasp-x.scss.tcd.ie:33000(mtls) to get external messages
2. listen local request on 127.0.x.y:443(mtls) to get device messages
3. a http client to do some cron tasks
4. cli to control manually (optional, do if enough time)

## User Case 1

run a p2p server at rasp-x.scss.tcd.ie:33000(external)

## User Case 2

run a p2p server at 127.0.x.y:443(internal)

## User Case 3

- periodically send healthz request to other external nodes(aka gateways),
  - if 200 OK, do nothing
  - otherwise, remove it from network list
- periodically send healthz request to internal nodes,
  - if 200 OK, do nothing
  - otherwise, remove it from device list
- periodically send register message to other networks, then send collected data to them, then get list from them and update the local list
