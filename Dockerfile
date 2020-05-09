# https://github.com/HouzuoGuo/k8s-hello-world
FROM ubuntu

RUN apt update && apt install -y busybox curl wget

COPY my-app /my-app

ENV PORT 20000
EXPOSE 20000

ENTRYPOINT ["/my-app"]
