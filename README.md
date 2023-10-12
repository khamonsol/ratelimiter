Reverse proxy to rate limit outgoing requests

best way to test is either to bang on it with curl like: 

curl http://localhost:1337/throttle OR  curl http://localhost:1337/error

Apache bench (ab) will show the rate limit in action if you have the patience. 


####
Notes

burst is the queue size, if we exceed that we will start getting 429's from the throttle endpoint. 
This will need to be tuned. Unfortunatly, delay=0 is invalid config, so when the queue goes empty, two requests to go back to back without a delay. 
Hopefully this will fall within burst limits of the upstream host, if not it should be possible to keep the queue primed through other means or lua script. 

This is a POC. Nginx does not provide a means to sync ratelimits across multiple instances without paying for a commerical lisence.This there is no way to make this HA. 
Hopefully our big fancy edge load balancer can provide similar functionality with HA capability. Alternatively we can look at HAproxy since HA is right in the name.


2023 UPDATE
Turns out HA proxy is very bad at this forward/egress proxy thing. Nginx is much better. The file have been updated with 2023 requirements

Check 2023 files for examples.. 

build docker: 
`docker build -t nginix2023-test -f ./Dockerfile-2023 .`

 
