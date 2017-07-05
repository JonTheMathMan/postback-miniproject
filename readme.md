# The Postback Machine

The postback machine takes an http request with json data containing an endpoint for delivery and data objects to deliver. The postback machine uses php to accept requests, redis to queue requests, and golang to deliver responses. If the method for delivery is set to GET the postback machine will swap the data keys with the keys in the endpoint url. If the method for delivery is set to POST the postback machine will deliver the data key values as POST form key values.

### Examples

```
POST
http://{server_ip}/ingest.php
{
  "endpoint":{ 
    "method":"GET",
    "url":"http://sample_domain_endpoint.com/data?title={mascot}&image={location}&foo={bar}" },
  "data":[ {
    "mascot":"Gopher",
    "location":"https://blog.golang.org/gopher/gopher.png" }
] 
}

Result:
GET
http://sample_domain_endpoint.com/data?title=Gopher&image=https%3A%2F%2Fblog.golang. org%2Fgopher%2Fgopher.png&foo=
```
```
POST
http://{server_ip}/ingest.php
{
  "endpoint":{ 
    "method":"POST",
    "url":"http://sample_domain_endpoint.com/" },
  "data":[ {
    "mascot":"Gopher",
    "location":"https://blog.golang.org/gopher/gopher.png" }
] 
}

Result:
$_POST == array(2){ "mascot"=>"Gopher", "location"=>"https://blog.golang.org/gopher/gopher.png" }
```
