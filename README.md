Reverse proxy to rate limit outgoing requests

best way to test is either to bang on it with curl like: 

curl http://localhost:1337/throttle OR  curl http://localhost:1337/error

Apache bench (ab) will show the rate limit in action if you have the patience. 


####
Notes

burst is the queue size, if we exceed that we will start getting 429's from the throttle endpoint. This will need to be tuned. Unfortunatly, delay=0 is invalid config, so when the queue goes empty, two requests to go back to back without a delay. Hopefully this will fall within burst limits of the upstream host, if not it should be possible to keep the queue primed through other means or lua script. 
