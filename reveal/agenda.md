# What is Docker (J)

# How to run containers (J)

## Exercise 1 (J)

# How to create containers (D)

## Exercise 2 (D)

# How does Docker work (J/D)

# Docker compose (J)

So up until now we have been running containers with the CLI. Which works, but if we've got a number of them, with different settings, we don't want to copy and paste commands, rebuild our images, modify our configuration files, etc..

Let's move away from CLI commands, and configure it with some YAML, cuz who don't love YAML?


# Exercise 3 (J)

# Best practices for creating Dockerfiles (D)

# Container clusters and Orchestration (J/D)

Not using Kubernetes, why?:
  - Only works with containers
  - long learning curve
  - complex to get started
  - large amount of unneeded 'extras'

### Consul

##### Background
- What is Consul? (It can be many things, but I use it mostly for Service discovery and health checking)
  - Service discovery:
    - A way of finding your applications if all you have is microservices or applications in a cluster.
    - Useful for running many containers on nodes:
      - no need to remember IP address and port for each
      - applications can dynamically add themselves
      - useful as DNS

#### Health Checks
    - Add a ping / curl / script to determine that application is up
      - this can be used to remove / add it to a Loadbalancer

Nomad:
- What is Nomad?
  - 'Easily deploy applications at any scale' - Tagline
  - A single binary that schedules applications and services
  - Plan changes, run applications, monitor deployments
  - Use for A/B (or blue/green) deployments
  - Rolling deployments
  - Application orchestration:
    - Can failover applications to other Nomad agents (if they exist) in the cluster
    - New agents can share the load if added

- Does this mean Nomad isn't as good?:
  - No:
    - If you don't care about following buzzwords
    - If you need to scale and orchestrate applications of any kind
    - If you don't want to mess with installations:
      - single binary, remember?
    - If you want to have most of the same K8s features, without the learning curve
  - Yes:
    - If your job requires K8s
    - If you want autoscaling of jobs (Nomad cannot (0.8.7) do this yet)

Nomad commands:
  - nomad job plan x
  - nomad job run x
  - nomad status x

Display and Change the job file:
  - nomad job plan x
  - nomad job run x

Note the canary:
  - nomad deployment promote x
  - nomad status x
  - nomad deployment fail x

Kill a node and monitor in web ui
  - nomad job stop x
