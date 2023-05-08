# training-go
All Go code and training.

Running the code:
```
cd webservice-start
go build .
./webservice-start
```

APIs:

```
curl localhost:3000/users
curl -X POST localhost:3000/users -d '{"FisrtName":"Frank", "LastName":"Go to NY"}'
curl -X PUT localhost:3000/users/1 -d '{"FisrtName":"Frank", "LastName":"Go to LA"}'
curl localhost:3000/users
curl localhost:3000/users/1
curl -X DELETE localhost:3000/users/1
```
