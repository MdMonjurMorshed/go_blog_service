FROM golang:1.19-alpine3.17 as build

WORKDIR app/

COPY src/ .

RUN ["go","mod","tidy"]
RUN ["go","build","main/runner.go"]

FROM alpine:latest
WORKDIR /app
COPY --from=build /go/app/runner .
CMD ["./runner"]
EXPOSE 3000