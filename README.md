##DOCKER SANDBOX

This is a simplistic (yet!) sandbox to test applications consisting of multiple containers, typically of containers of your application and some infrastructural containers. It should save efforts writing about 5/6 of `docker-compose.yaml` (such as definitions of the services, networking etc). 
One may alter the resulting yaml as you want (`names`, `ports`, `images`, `ENVs` etc).
Infrastructural services such as popular databases (`mongo`, `postgres`, `mysql`, `cassandra` etc) and other popular tools (`redis`, `kafka`, `schema-registry`) will be supported with some generic definitions which one also can tweak. Also there'll be an API to add your own services or you can request support of the service you need.