Mail Manager Service:
====================

Requirements:
------------

[docker](https://www.docker.com/)

Usage:
-----

First of all get public key that should be used for grpc server.

1. Make a build:
```bash
docker-compose build
```

2. Run containers:
```bash
docker-compose up
```

For queue management open [http://localhost:15672](http://localhost:15672/)  
Use credentials from .env file. By default: *guest*/*guest*
Go to [http://localhost:15672/#/queues](http://localhost:15672/#/queues)  
Into Publish messages put the following json:
```json
{"action":"register", "payload":{"name":"serg", "to":"siarhei.sharykhin@itechart-group.com", "token":"12345"}}
```
It will send a real mail to *siarhei.sharykhin@itechart-group.com*

Currently only two mail-boxes are supported for getting mails:  
- *artsem.vasilevich@itechart-group.com*  
- *siarhei.sharykhin@itechart-group.com*

To test failing mails use the following env variable:
```bash
docker-compose run -e TEST_FAIL=OK gl-mail-manager-golang
```
This will imitate some sort of failing and write it into a log file.

Example:
```bash
curl -u guest:guest -H "content-type:application/json" -X POST -d'{"properties":{"delivery_mode":2},"routing_key":"mail","payload":"{\"action\":\"register\", \"payload\":{\"name\":\"serg\", \"to\":\"siarhei.sharykhin@itechart-group.com\", \"token\":\"12345\"}}","payload_encoding":"string"}' http://localhost:15672/api/exchanges/%2f/amq.default/publish
```