# services:   cd .. && ./start
# server:     cd ./server && go run main.go
request:   nsq_tail -topic=nsq_rr.req -lookupd-http-address=localhost:4161
response:  nsq_tail -topic=z...rsp.main.node01 -lookupd-http-address=localhost:4161 
# client:     sleep 2 && cd ./client && go run main.go