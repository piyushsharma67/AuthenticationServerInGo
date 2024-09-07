# Use an official Golang runtime as a parent image
FROM golang:1.22-alpine
# Set the Current Working Directory inside the container

WORKDIR /app

COPY ./ /app/

RUN go mod download

RUN go install -v github.com/air-verse/air

RUN go build -o chat-application .

CMD ["air", " --build", "./chat-application"]