FROM golang
RUN mkdir ./godomicroservice
ADD . ./godomicroservice
WORKDIR /godomicroservice
RUN go build -o /godomicroservice/choose/main.go ./choosemain
RUN go build -o /godomicroservice/api/main.go ./apimain
CMD [ "/godomicroservice/choosemain ; /godomicroservice/apimain" ]
