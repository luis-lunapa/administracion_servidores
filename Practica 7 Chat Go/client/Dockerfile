FROM golang:latest
WORKDIR /var/www/html
COPY . /var/www/html
# EXPOSE 80
ENV NAME Client
RUN go get -v github.com/chzyer/readline
RUN go build -o cliente main.go
#CMD ["./cliente", "|", "echo", "conn localhost:8081"]
RUN echo conn servidor:80 | ./cliente
# CMD ["echo", "conn", "localhost:8081", "|", "go", "run", "main.go"]

# CMD ["go", "run", "main.go", "|", "echo", "conn", "localhost:8081"]