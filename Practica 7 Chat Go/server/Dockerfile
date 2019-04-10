FROM golang:1.8-alpine
WORKDIR /var/www/html
COPY . /var/www/html
EXPOSE 80
ENV NAME Server
CMD ["go", "get", "https://github.com/chzyer/readline"]
CMD ["go", "run", "main.go"]