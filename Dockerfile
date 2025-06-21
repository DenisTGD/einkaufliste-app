FROM golang:1.24-alpine


WORKDIR /go/src/github.com/DenisTGD/einkaufsliste-backend


COPY backend/go.mod ./
COPY backend/go.sum ./
RUN go mod download

COPY backend/. ./
RUN go build -o server ./cmd

EXPOSE 8080
CMD ["./server"]
