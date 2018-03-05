# go-echo-sqlboiler
Golang sample project using Echo and SQLBoiler.  

## WebFW
### Echo
https://github.com/labstack/echo

## ORM
### SQLBoiler
https://github.com/volatiletech/sqlboiler
```
# generate new orm codes
$ make gen

# create new db schema rs ... replace schema
$ make rs
```

## DB Migration
### DBFlute
https://github.com/dbflute/dbflute-core

## Swagger
https://github.com/swaggo/swag
https://github.com/swaggo/echo-swagger

## Hot reload
https://github.com/codegangsta/gin
```
# run echo web server via gin
$ make run
$ curl http://localhost:1314
```

## Test watch
Use looper for test watching.  
looper  
https://github.com/nathany/looper  

Starting test watching when you entered blow command.
```
$ make watch
```
