FROM golang:1.13.1-stretch

RUN apt-get install git

WORKDIR /go/src/tickerclock

COPY . .

RUN go mod vendor && go build

CMD ./tickerclock -secPMin=60 -secPHour=3600 -allowUpdate=600 -deadline=10800 -port=8080

EXPOSE 8080/tcp
