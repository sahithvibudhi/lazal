FROM golang:latest 

RUN mkdir /app \
    git clone https://github.com/sahithvibudhi/lazal.git /app 

WORKDIR /app 

RUN go build -o main . 

CMD ["/app/main"]

EXPOSE 5555