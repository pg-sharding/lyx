FROM golang:1.26

RUN apt-get update \
    && apt-get install -y --no-install-recommends ragel \
    && rm -rf /var/lib/apt/lists/*
RUN go install golang.org/x/tools/cmd/goyacc@latest

WORKDIR /lyx

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENTRYPOINT ["make", "build"]
