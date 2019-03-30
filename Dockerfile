FROM golang
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go install recipe-app
CMD ["recipe-app"]
