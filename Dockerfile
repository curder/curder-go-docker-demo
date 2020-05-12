# Compile stage
FROM golang:1.14 AS build-env
ADD . /dockerdev
WORKDIR /dockerdev
RUN go mod download
RUN go build -o /server

# Final stage
FROM debian:buster
WORKDIR /
COPY --from=build-env /server /
EXPOSE 80
CMD ["/server"]