FROM golang:1.18

# create directory app
RUN mkdir /app-golang

# set or make /app our working directory
WORKDIR /app-golang

# copy all files to /app
COPY . .

RUN go build -o alta-rest-api

CMD ["./alta-rest-api"]
