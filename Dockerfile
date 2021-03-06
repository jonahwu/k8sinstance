#FROM golang:latest AS build
FROM golang:1.11 AS build

ENV GOARCH_SRC=$GOPATH/src/github.com/k8sinstance
#ENV CGO_ENABLED=1
#ENV GOOS=linux
#ENV NOMS_VERSION_NEXT=1
#ENV DOCKER=1

RUN mkdir -pv $GOARCH_SRC
COPY . ${GOARCH_SRC}
RUN go test github.com/k8sinstance/...
RUN ls $GOPATH/src/github.com/k8sinstance/cmd/k8sinstance -alh
RUN cd $GOPATH/bin && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build github.com/k8sinstance/cmd/k8sinstance
#RUN go install -v 
RUN cp $GOPATH/bin/k8sinstance /bin/k8sinstance
RUN ls $GOPATH/bin/ -alh
RUN ls /bin/ -alh

FROM alpine:latest

COPY --from=build /bin/k8sinstance /k8sinstance
RUN ls / -alh
RUN apk add --no-cache curl
RUN curl -L https://github.com/sequenceiq/docker-alpine-dig/releases/download/v9.10.2/dig.tgz|tar -xzv -C /usr/local/bin/

RUN apk add --update ca-certificates
ADD http://github.com/tsenart/vegeta/releases/download/v5.9.0/vegeta-v5.9.0-linux-amd64.tar.gz /go/bin/vegeta.tar.gz
RUN cd /go/bin && tar xzvf /go/bin/vegeta.tar.gz 
RUN chmod +x /go/bin/vegeta
RUN ln -s /go/bin/vegeta /usr/local/bin/vegeta
COPY httpbench /httpbench
#VOLUME /data
EXPOSE 8000
ENV NOMS_VERSION_NEXT=1
RUN chmod +x ./k8sinstance
ENTRYPOINT [ "./k8sinstance" ]

#CMD ["serve", "/data"] ]
