# Device User Case
# Swati Poojary
## Overview
1. listen local request on 127.0.x1.y:443(mtls) to get data from gateway
2. listen local request on 127.0.x2.y:443(mtls) to get message from sensors
3. a http client to do some cron tasks

## User Case 1

run a p2p server at 127.0.x1.y:443(external)

## User Case 2

run a p2p server at 127.0.x2.y:443(internal)

## User Case 3

- periodically send healthz request to sensors,
  - if 200 OK, do nothing
  - otherwise, remove it from device list
- periodically send register message to other networks, then send collected data to them, then get list from them and update the local list
