# P2P Server
# Anoushka Mittal
## internal

- handle https://ip:port/register, save other node, get name from request body, reply a token.
- handle https://ip:port/unregister, remove node in the local by the token in body
- handle https://ip:port/healthz, heartbeat, return 200 OK
- handle https://ip:port/list, return all known nodes, identify network by token
- handle https://ip:port/message, send message to an other node, name like a file path, e.g. "./name1/subname1" means find in internal

**All above, if identify failed, return 400 bad request**
