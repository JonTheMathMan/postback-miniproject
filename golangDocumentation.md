delivery.go uses a continuous/repetitive function call to take requests from the redis queue and then decide what to do with them based off what endpoint["method"] is set to.


### The Functions of delivery.go

**func dialRedis() redis.Conn**
Establishes the connection to redis through redigo and returns the redigo connection. (see the redigo documentation) http://godoc.org/github.com/garyburd/redigo/redis

**func continuousFuncCall(repeatFunc func(redis.Conn),redisConn redis.Conn)**
Takes a function as the first argument and the argument to pass to that function as the second argument. Calls that function every 100 milliseconds.

**func sendAResponse(redisConn redis.Conn)**
Checks whether to send response with GET or POST, then delivers the response, then logs the result to deliveryLog.log

**func getRequestMaps(redisConn redis.Conn) (map[string]string, map[string]string)**
Uses the redis connection and redigo's "StringMap" method to take the redis information and put it into convenient golang maps. Then it deletes the information from redis to keep the information stored in redis to a minimum.

**func placeDataIntoTypeUrlValues(data map[string]string) url.Values**
Takes the data map that was made from the redis hash called data and puts it into the type url.Values so that it can be sent in golang's http.PostForm()

**func placeDataMapIntoGetResponse(urlWithKeys string, data map[string]string) string**
Swaps the keys in the data map with the keys in the endpoint url and returns the new url.

**func getUrlDataKeys(s string) []string**
Takes the endpoint url and returns a string slice containing the keys from the url.
