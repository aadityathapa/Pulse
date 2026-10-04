FROM golang:alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY main.go ./
RUN CGO_ENABLED=0 go build -o /pulse .

FROM scratch
COPY --from=build /pulse /pulse
EXPOSE 8080
ENTRYPOINT ["/pulse"]
