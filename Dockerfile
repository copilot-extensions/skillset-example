FROM golang:1.21 AS builder

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o bin/skills-example .

FROM ubuntu:20.04 AS run
WORKDIR /app
COPY --from=builder /app/bin/ /app/bin/
EXPOSE 8080
ENTRYPOINT ["bin/skills-example"]