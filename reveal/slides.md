---
title: Docker
theme: solarized
separator: '---'
verticalSeparator: '\*\*\*'
revealOptions:
  transition: slide
---

<img class="plain"  src="https://www.docker.com/sites/default/files/social/docker_facebook_share.png"/>

---

## What is Docker?

***

- Package software

***

- Immutable distribution


---


# Creating images

***

## Dockerfile

```dockerfile
FROM node:alpine  # Container base image

WORKDIR /app/
COPY app.js package.json ./  # Copy files into container
RUN npm install  # Execute command in container build

# What to run when finished container will be started
ENTRYPOINT ["node", "app.js"]
```

***

```sh
.
├── app.js
├── Dockerfile
└── package.json
```

***

```sh
$ docker build myapp .

$ docker tag myapp 0.0.0.0:5000/myapp:v1.0

$ docker push myapp
The push refers to repository [0.0.0.0:5000/newapp]
d0ee75f51314: Layer already exists
7bff100f35cb: Layer already exists
latest: digest: sha256:7deec43013ea79005c342b407e37f42061dc638351512ad6cfec3b754a48a023 size: 739

```

---

## docker-compose

so cool

```yml
version: '3'

services:

  registry:
    image: registry
    ports:
      - 5000:5000
    volumes:
      - data:/var/lib/registry
      - ./config.yml:/etc/docker/registry/config.yml

volumes:
  data:
```

---
