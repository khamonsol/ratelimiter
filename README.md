everse proxy to rate limit outgoing requests

best way to test is either to point on curl like: 

curl http://localhost:1337/throttle OR  curl http://localhost:1337/error

Apache bench (ab) will show the rate limit in action if you have the patience. 
