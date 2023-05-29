FROM golang:1.19
WORKDIR /home
COPY ./ ./
RUN go  build -o damage-f
RUN chmod +x damage-f
CMD ["./damage-f"]
