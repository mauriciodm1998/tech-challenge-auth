#mauricio1998/tech-challenge-auth

FROM golang as builder

RUN mkdir app
COPY ./ app
WORKDIR app
RUN CGO_ENABLED=0 go build -o dist/tech-challenge-auth cmd/auth/main.go

FROM alpine as runner

RUN mkdir app
COPY --from=builder go/app/dist/tech-challenge-auth app/
RUN chmod +x app

EXPOSE 3002
EXPOSE 3003
WORKDIR /app

ENTRYPOINT [ "./tech-challenge-auth", "--config-dir", "." ]