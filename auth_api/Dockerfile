FROM golang
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o /godocker

EXPOSE 8080

CMD [ "/godocker" ]