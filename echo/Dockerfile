FROM golang:1.19.1-alpine

WORKDIR $GOPATH/src/github.com/Hargeon/kubelrn/echo

COPY . .

RUN go mod tidy
RUN go build -o echo

EXPOSE 3000

CMD ["./echo"]