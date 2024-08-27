# Docker and Kubernetes 

- Presentation of containers
https://docs.google.com/presentation/d/1TClOq3s__IAEwI7xQbVTJ1nMv9UNQ1OcoQ6-OKSy3Ko/edit?usp=sharing


- Delete all containers in docker 

```
docker rm -f $(docker ps -aq)
```
- To delete all images

```
docker rmi -f $(docker images -q)
```

- To delete unused networks

```
docker network prune
```

- create a network

```
docker network create --subnet=192.168.10.0/24 my-custom-network
```

- create a busybox container with a simple shell script

```
docker run -d --network my-custom-network --name bb1 busybox sh -c "while true;do sleep 3600; done"
```
```
docker run -d --name app2 -e APP_PORT=8081 -p 48081:8081 \
jpalaparthi/rakutenapp1:v0.9
```
- save an image to tar file

```
docker save -o <tarfile.tar> <imagename>
```

- To build docker images 

```
docker build . -f Dockerfile-scratch-nonroot -t jpalaparthi/rakutenapp1:v0.9
```

- To push Docker images

```
docker push jpalaparthi/rakutenapp1:v0.9
```