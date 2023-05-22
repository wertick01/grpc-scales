FROM golang:1.20-alpine AS build

WORKDIR /app

# prevent the re-installation of vendors at every change in the source code
COPY . .
RUN go mod tidy

# Build app
RUN CGO_ENABLED=0 go build -o /server ./cmd/server


FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /server /server

EXPOSE 8000

USER nonroot:nonroot

CMD [ "/server" ]