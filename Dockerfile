FROM golang:1.17.2

RUN mkdir /app

WORKDIR $GOPATH/src/github.com/Arturo0911/measurements-realtime

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

# WORKDIR /app

# ADD go.mod .

# ADD go.sum .

# RUN go mod download

# RUN go get github.com/Arturo0911/measurements-realtime

# RUN go get github.com/Arturo0911/measurements-realtime

# ADD . .

# RUN go get -t -d ./...

# RUN go install ./...

EXPOSE 8000

# ENTRYPOINT measurements-realtime --build="go build main.go" --command=./main

CMD ["measurements-realtime"]