ingest.php creates a request hash, an endpoint hash, and a data hash in redis for each request it receives. It creates an id number using the redis command INCR. It uses the id number to keep track of which hash belongs to which request.

The actual structure of the id's, lists, and hashes in redis looks like this:
```
requestsCount == 1
queueList{
    "request1"
}
request1{
    "endpoint":"endpoint1",
    "data":"data1"
} 
endpoint1{
    "method":"GET",
    "url":"someplace.com"
}
data1{
    "someKey":"someValue",
    "anotherKey":"anotherValue"
}
```
