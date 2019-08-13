FROM golang
COPY ../pb /go/src/
COPY . /go/src/choose
WORKDIR /go/src/choose
RUN go get .
ENTRYPOINT go run main.go
EXPOSE 3000