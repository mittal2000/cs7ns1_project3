# Network
# Anoushka Mittal
one possible structure

## Gateway

| net name | global listen | local listen |
| - | - | - |
| net1 | rasp-019.scss.tcd.ie:33000 | 127.0.0.1:443 |
| net2 | rasp-020.scss.tcd.ie:33000 | 127.0.0.1:443 |

## Devices

| device name | listen1 | listen2 |
| - | - | - |
| dev1 | 127.0.0.2:443 | 127.0.1.1:443 |
| dev2 | 127.0.0.3:443 | 127.0.2.1:443 |

## Sensors

| sensor name | connect to | listen |
| - | - | - |
| dev1/sensor1 | 127.0.1.1:443 | 127.0.1.2:443 |
| dev1/sensor2 | 127.0.1.1:443 | 127.0.1.3:443 |

| sensor name | connect to | listen |
| - | - | - |
| dev2/sensor1 | 127.0.2.1:443 | 127.0.2.2:443 |
| dev2/sensor2 | 127.0.2.1:443 | 127.0.2.3:443 |
