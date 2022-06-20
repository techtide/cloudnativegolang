FROM golang
RUN mkdir ./godomicroservice
ADD . ./godomicroservice
WORKDIR /godomicroservice
# RUN go build -o /godomicroservice/choose/main.go ./choosemain
# RUN go build -o /godomicroservice/api/main.go ./apimain
CMD [ "go run /godomicroservice/choosemain ; go run /godomicroservice/apimain" ]