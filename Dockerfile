FROM golang:latest as build
WORKDIR /app
ADD . .
RUN mv main.dogo main.go
RUN CGO_ENABLED=0 go build -o main .

FROM alpine
COPY --from=build /app/main /app/
ENV ADDRESS ":8080"
CMD ["/app/main"]
