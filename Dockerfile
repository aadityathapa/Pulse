FROM golang:1.27-alpine AS build
WORKDIR /src
COPY go.mod main.go ./
RUN CGO_ENABLED=0 go build -o /pulse .

FROM scratch
COPY --from=build /pulse /pulse
EXPOSE 8080
ENTRYPOINT ["/pulse"]
