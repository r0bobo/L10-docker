```bash
$ docker build --tag app-singlestage -f Dockerfile.singlestage .
$ docker build --tag app-multistage -f Dockerfile.multistage .

$ docker image ls | grep demo-app
app-singlestage            latest              a843e21e18d7        19 hours ago        796MB
app-multistage             latest              1189765559cf        19 hours ago        10.9MB

$ docker run --rm -p 80:80 demo-app
```
