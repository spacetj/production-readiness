FROM golang:1.14.3-alpine AS build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o todo ./cmd/todo
RUN chmod a+x todo

FROM scratch
COPY --from=build /app/todo /
ENTRYPOINT [ "/todo" ]
EXPOSE 8000
