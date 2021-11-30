FROM golang:1.17.2 as go_builder

WORKDIR $GOPATH/src/github.com/Arturo0911/measurements-realtime

COPY . .

RUN go get -t -d ./...

RUN go install ./...

EXPOSE 8000

CMD ["measurements-realtime"]