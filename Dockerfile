#Build stage
FROM golang:1.14.2 AS build

RUN mkdir /app
ADD . /app/
WORKDIR /app

RUN CGO_ENABLED=0 go build -ldflags "-X main.build=production" -o main main.go

#Second stage - release
FROM alpine:latest
RUN mkdir /app
WORKDIR /app

COPY --from=build /app/main .
#COPY --from=build /app/config.json .

EXPOSE 8080
CMD ["/app/main"]
