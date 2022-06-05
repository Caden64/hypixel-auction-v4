FROM golang:1.18.3

WORKDIR ./app

ADD database ./database
ADD HypixelRequests ./HypixelRequests

ADD go.mod ./
ADD go.sum ./

ADD main.go ./

RUN go mod download

RUN go build -o /app

CMD [ "/app" ]
