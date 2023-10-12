Reverse proxy to rate limit outgoing requests

best way to test is either to bang on it with curl like: 

curl http://localhost:1337/throttle OR  curl http://localhost:1337/error

Apache bench (ab) will show the rate limit in action if you have the patience. 


####
Notes

burst is the queue size, if we exceed that we will start getting 429's from the throttle endpoint. 
This will need to be tuned. Unfortunatly, delay=0 is invalid config, so when the queue goes empty, two requests to go back to back without a delay. 
Hopefully this will fall within burst limits of the upstream host, if not it should be possible to keep the queue primed through other means or lua script. 

---

##October 2023 Update##

Upon further testing, it has become evident that HAProxy doesn't excel at handling forward/egress proxy tasks. NGINX, on the other hand, has proven to be significantly more efficient. The files have been updated in accordance with our 2023 requirements.

By setting limit_conn = 1, we've addressed the previous delay=0 concern. This ensures that proxy_pass never allows more than one concurrent connection. With the current configuration, an overload of concurrent connections beyond what proxy_pass can accommodate will result in 503 errors.

Unfortunately, there isn't a reliable way to have existing proxy-based rate limiters queue connections and throttle them without client retries, unless we invest in an enterprise license. Furthermore, even with a paid license, it's uncertain whether the available mechanisms would cater to our exact needs. While one can manipulate Nginx into this behavior using proxy_cache, it's neither officially supported nor guaranteed to work consistently.

To assist with rate limit tuning, I've developed a test REST API that shows the connections reaching the proxy_pass. Included are three executables and a Go source file, which, although currently basic in terms of measurement capabilities, can be expanded upon. For demonstration purposes, the config also connects to an external public site, showcasing that it's possible to connect to sites beyond our control and still successfully negotiate an SSL connection.

For load testing the proxy, I utilized Vegeta. Relevant scripts are available in the repository. It's noteworthy that the proxy's current settings are highly conservative. Hence, even seemingly permissible connections can fail. While I haven't extensively fine-tuned the NGINX configuration, it appears to align closely with our objectives.

I hope you find this resource useful.

###Setup###

**Configuration:** Modify the nginx2023 config file to suit your needs. Note: When running NGINX inside a container and redirecting traffic to localhost, a unique address (already present in the nginx config file for port 8112) is necessary.

**Docker:** Build using the provided Dockerfile-2023 to integrate the updated config and execute it. From my initial experience, simply mounting with a bind mount wasn't superseding the default configuration. However, if mounting is an option for you, it facilitates easier iteration and tuning.

**Local HTTP Server:** To monitor backend connections in real-time, start the local HTTP server. A script (httpMonitorStart.sh) is available for this purpose. Ensure you replace the executable with the one compatible with your OS.

###Testing###

**Direct Load Test:** Initially, run a load test directly against the local server. This will help you gauge its capacity to handle traffic without connection drops.

**Proxy Test:** Following that, conduct the test via the proxy. Monitor the proxy logs for 503 errors, indicating areas that require tuning. For a more holistic understanding of its operation in our context, it might be essential to develop a genuine testing tool capable of executing retries against the 503 errors.
