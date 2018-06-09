FROM golang:1.10
RUN mkdir -p /go/src/github.com/therealplato/shortlink
ADD . /go/src/github.com/therealplato/shortlink
RUN go build -o /shortlink github.com/therealplato/shortlink
CMD /shortlink
