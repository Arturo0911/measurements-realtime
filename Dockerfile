FROM golang:1.17.2

RUN mkdir /app

WORKDIR $GOPATH/src/github.com/Arturo0911/measurements-realtime

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 8000

# ENTRYPOINT measurements-realtime --build="go build main.go" --command=./main

CMD ["measurements-realtime"]