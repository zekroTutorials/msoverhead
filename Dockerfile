FROM golang:1.16-alpine AS build
WORKDIR /build
COPY . .
RUN go build -v -o node ./...

FROM alpine:latest
COPY --from=build /build/node /bin/node
EXPOSE 80
ENTRYPOINT ["node"]