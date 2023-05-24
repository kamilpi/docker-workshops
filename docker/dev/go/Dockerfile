FROM golang as build 

WORKDIR /app

COPY main.go /app/main.go

RUN go mod init main && \
    go mod tidy && \
    go build -o app

FROM ubuntu

COPY --from=build /app/app /root/app 
COPY docker.img /root/docker.img
CMD ["/root/app"]