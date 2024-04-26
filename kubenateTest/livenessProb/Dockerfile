#Dockerfile
FROM golang:1.21 As builder 
RUN mkdir /src
ADD . /src
WORKDIR /src 

RUN go env -w GO111MODULE=auto  
RUN go build -o main .

FROM ubuntu:latest
WORKDIR /

COPY --from=builder /src/main /main 
EXPOSE 5000
ENTRYPOINT ["/main"]
