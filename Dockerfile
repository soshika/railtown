FROM golang:1.18
EXPOSE 9093

WORKDIR /prj

COPY . /prj/
RUN apt-get install gcc
RUN go build main.go
CMD ./main