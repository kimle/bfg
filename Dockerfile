FROM alpine:3.7
WORKDIR /go/src/app
RUN mkdir static
COPY server .
COPY static ./static
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
EXPOSE 8080
CMD ["./server"]
