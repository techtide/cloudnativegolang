# godomicroservice
<img src="https://api.cirrus-ci.com/github/techtide/godomicroservice.svg">

Other related projects done here:
* [gotrymicroservice](https://github.com/techtide/gotrymicroservice)
* [godotest](https://github.com/techtide/godotest)

Using Go, testing [kind of], gRPC and Protocol Buffers using ``protoc-gen``, Docker, Mux, Kubernetes, and [kind of] Cirrus CI.

API will take two numbers and do a combinatorics function, just as an example:

* ``/choose/5/3`` is 5 choose 3, the value of which is 10.

* ``/pick/5/3`` is 5 pick 3, the value of which is 60.

The reason this was picked this very random service idea is because I couldn't think of anything else.

### Client

The Releases tab on GitHub contains a link to download the client script, which connects to the API when it is running on the local machine, as per the instructions.

If you're on a system with wget, Python 3, and gunzip installed run, ```wget https://github.com/techtide/godomicroservice/files/3497746/wayra_client.zip && gunzip wayra_client.zip && python3 wayra_client.py```. 

This is just a very simple CLI utility to interact with the API, which should be running on your local machine at 127.0.0.1:8081.

### Justifications

Justification 1: Is it cloud native?
Yes, it is cloud native. There are example Dockerfiles and Kubernetes configs which should, probably, work. This means that the microservice is scalable since it is containarized, but I would hesistate as to why you would ever want a scalable probability calculator.

Justification 2: Does it follow the 12 Factor App Impacts?
1. Codebase is clearly on Git and version control, attempts were made to use CI to communicate with the API.
2. Dependencies are written above, ``go get.`` will traverse through all the dependencies declared in these files and install them as appropriate.
3. Configs are in the environment.
4. N/A
5. That has also been done, in Kubernetes config you can see this.
6. Yes.
7. This is done.
8. N/A
9. It's kind of hard to not have this in something like this.
10. There's no difference.
11. Go logs and testing pretty much do this I think.
12. There are no such admin or management tasks involved!

*That Cirrus CI test really shouldn't be passing, as I haven't found a way for me to allow Cirrus CI to test endpoints, since the endpoints are not running on their end, and aren't running at all.*



