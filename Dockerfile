# Dockerfile
FROM golang:1.14.2

WORKDIR /app
COPY . .

RUN go get -u github.com/go-chi/chi
RUN go get -u github.com/swaggo/http-swagger
#RUN swag init reload 제대로 하려면 정리 필요
RUN go get github.com/cespare/reflex
#RUN go get github.com/go-delve/delve/cmd/dlv

EXPOSE 9999
#EXPOSE 2345

#CMD ["reflex", "-r", "(\.go$|go\.mod)", "-s", "--", "sh", "-c", "dlv debug --listen=:2345 --headless=true --api-version=2"]
#CMD ["reflex", "-r",  "'(\.go\$|go\.mod)'",  "-s",  "'go run main.go'"]
CMD ["reflex", "-c", "reflex.conf"]
