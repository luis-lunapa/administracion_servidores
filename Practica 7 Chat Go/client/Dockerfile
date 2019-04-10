FROM golang:1.8-alpine
WORKDIR /var/www/html
COPY . /var/www/html
# EXPOSE 80
# ENV NAME Server
CMD ["go", "get", "https://github.com/chzyer/readline"]
# CMD ["go", "run", "main.go", "echo", "|", "conn", "127.0.0.1:80"]
CMD ["echo", "conn", "127.0.0.1:80", "|", "go", "run", "main.go"]