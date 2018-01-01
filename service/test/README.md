To run console tests

terminal 1 start infrastructure (consul, nsqd...)
```
   ./start
```

terminal 2 start server
```
   cd server
   go run main.go
```

terminal 3 run client
```
   cd client
   go run main.go
```
