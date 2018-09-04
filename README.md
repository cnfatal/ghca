# GHCA

## What's this

 `China Telecom 189 fei young` encrypted account caculator

 Source Code From:[007xiaoxingxing/GhcaDialer](https://github.com/007xiaoxingxing/GhcaDialer )(deleted)

## Make

```bash
go build .
chmod +x ./ghca
```
or
```bash
./build.sh
```
it will build a docker image locally named:`fatalc/ghca`

## Usage

### Params

|name|type|default|
|---|---|---|
|username|string|""|
|password|string|""|
|server|bool|false|
|port|int|8080|

### Example

Binary Run Locally:
```bash
$ ./ghca --username <username> --password <password>
~ghcaxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
```
Binary Run As A Service:
```bash
$ ./ghca --server
```
on a special port:
```bash
$ ./ghca --port 8000 --server
```
Then:
```bash
$ curl 127.0.0.1:8000
XXXXXYoung 3.09 XXXX
XXXX: GET /?username=<username>&password=<password>
```

```bash
$ curl 127.0.0.1:8000/?username=<username>\&password=<password>
ghcaxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
```

## Docker

Run with docker are the same
```bash
$ docker run -it --rm fatalc/ghca --usernmae <username> --password <password>
~ghca...
```
or as a service
```bash
docker run -it --name ghca -d -p 8080:8080 fatalc/ghca
```