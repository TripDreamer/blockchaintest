FROM google/golang

WORKDIR /gopath/src/blockchain
ADD .  /gopath/src/blockchain
RUN go get github.com/Kevin-miu/blockchain/cli

CMD []
EXPOSE 9119
ENTRYPOINT cli
