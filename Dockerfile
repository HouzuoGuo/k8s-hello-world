# https://github.com/HouzuoGuo/k8s-hello-world
FROM ubuntu

COPY my-app /my-app

ENV PORT 20000
EXPOSE 20000

ENTRYPOINT ["/my-app"]
