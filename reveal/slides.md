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

## Container orchestration

***

Now that we have built our Docker configurations, and decided what we are going to run. Let's go straight to prod!

All is gonna be great, right? Docker solves everything!

***

Up until your server goes down and your containers fail at 3am. Shit.

***

Notification goes off, and you hop out of bed to:

```
$ ssh server
$ docker restart container

```

'cuz that's what we do, right?
***

- Alright, but what if I had a Loadbalancer with multiple servers looking after my containers?
- Sure, but what if they fail?


***

What if there was a better way to restart, scale, reprovision or take care of your applications when your containers, or even servers hosting the containers, fail?

***

Can you guess where this is going?

***

<img class="plain"  src="https://portworx.com/wp-content/uploads/2017/08/wordpress-kubernetes.png"/>

***

Just kidding, we aren't going to use Kubernetes.

Sorry.

(But why?)

***

But we are going to use something!

***

We are going to use **Consul** and **Nomad** to look after our applications
(Notice I say applications, not _containers_)

***

What! What?  

***

### Consul for 'Service Discovery'

<img class="plain" src="https://www.datocms-assets.com/2885/1526066558-image-4.png?fit=max&fm=png&q=80&w=2500"/>

***

### Nomad for applications and containers

<img class="plain" src="https://www.nomadproject.io/assets/images/guide-ui-jobs-list-356d8d7d.png"/>

***

### Demo 
