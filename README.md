## Introduction

An IM project using Golang.

## Build && Run

This repo use Go Module.You can just build with command`go build`.

Don't forget to edit conf/application.yml before running.

Or just using Dockerfile
```
docker build -f Dockerfile -t pusidun/im_golang .
docker run  -p 8085:8085 --name im_golang pusidun/im_golang -it /bin/bash
```

PS:MySQL can also start using docker 
```
docker pull mysql:5.7
docker run --name golang-mysql -e MYSQL_ROOT_PASSWORD=12345abcde -p 3306:3306 -d mysql:5.7
```

## Design


## LICENSE

This repo is under *MIT LICENSE*
