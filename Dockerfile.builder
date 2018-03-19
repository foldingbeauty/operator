FROM golang:1.10

#switch to our app directory
RUN mkdir -p /go/src/github.com/foldingbeauty/operator
WORKDIR /go/src/github.com/foldingbeauty/operator

#copy the source files
COPY main.go /go/src/github.com/foldingbeauty/operator
COPY vendor /go/src/github.com/foldingbeauty/operator/vendor
COPY pkg /go/src/github.com/foldingbeauty/operator/pkg

#disable crosscompiling 
ENV CGO_ENABLED=0

#compile linux only
ENV GOOS=linux

#build the binary with debug information removed
RUN go build  -ldflags '-w -s' -a -installsuffix cgo -o operator