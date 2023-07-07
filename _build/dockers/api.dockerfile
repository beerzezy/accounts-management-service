FROM golang:1.19

#SET WORKING DIRECTORY
WORKDIR /src

#COPY CODE INTO WORKSPACE
COPY . .

#BUILD APP
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/servapi ./cmd/server
#REDUCE SIZE
RUN chmod +x -R /go

# API
COPY /_build/env/api-env.yml /src/env/config.yml

CMD ["/go/servapi"]


