Reverse proxy to rate limit outgoing requests

best way to test is either to bang on it with curl like: 

curl http://localhost:1337/throttle OR  curl http://localhost:1337/error

Apache bench (ab) will show the rate limit in action if you have the patience. 
