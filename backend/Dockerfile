FROM alpine

COPY . /go/src/github.com/iost-official/explorer/backend/
ENV GOPATH="/go"

WORKDIR /go/src/github.com/iost-official/explorer/backend/

RUN chmod +x ./start.sh
CMD ["/bin/ash", "-c", "./start.sh"]