# APIRestWithGo
Implementing an Api Rest using Golang to get information from a server and to know if the configurations have changed
Technologies USED:
Language: Go
Data Base: CockroachDB
API Router: fasthttprouter
Interface: Vue.js & bootstrap-vue.js

## Project setup

### 1. Create and Run COCKROACH DATABASE (domains)
Run this commands in the folder where you have installed cockroach. In case you don't have the command "cockroach" in your PATH use ./cockroach.exe
```
cockroach start --insecure --listen-addr=localhost
```
```
cockroach sql --insecure --host=localhost:26257
```
```
CREATE USER IF NOT EXISTS davidobando99;
```
```
CREATE DATABASE serversdb;
```
```
GRANT ALL ON DATABASE serversdb TO davidobando99;
```
You can see your database in localhost:8080

### 2. Run backend in Go
```
go run main.go
```

### 3. Run vue.js view
To run the view you have to install node.js and vue.js. And go to the folder view of this proyect. Then use the command..
```
npm run serve
```
Now go to your browser and put localhost:8081 to see the interface
