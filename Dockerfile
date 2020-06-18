FROM golang:latest as build
WORKDIR /app
ADD . .
RUN echo 'package main' > main.go
RUN echo 'import "github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/api"' >> main.go
RUN echo 'func main() {api.Run()}' >> main.go
RUN CGO_ENABLED=0 go build -o main .

FROM alpine
COPY --from=build /app/main /app/
ENV ADDRESS ":8080"
CMD ["/app/main"]
