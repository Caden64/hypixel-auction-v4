FROM golang:1.18.3

WORKDIR ./app

COPY MongoDatabase ./MongoDatabase
COPY HypixelRequests ./HypixelRequests

COPY go.mod ./
COPY go.sum ./
COPY RedisDatabase ./RedisDatabase
COPY MojangRequests ./MojangRequests
COPY main.go ./

RUN go mod download

RUN go build -o /app

CMD [ "/app" ]
