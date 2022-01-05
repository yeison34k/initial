FROM golang:alpine

RUN mkdir /pocone
COPY . /pocone
WORKDIR /pocone
RUN go build -o main .
CMD ["/pocone/main"]