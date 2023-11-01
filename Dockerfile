FROM golang:1.21

WORKDIR /usr/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build .

EXPOSE 2345
CMD ["go build . && ./jenkins-golang"]