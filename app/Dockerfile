FROM node:16.17.0-alpine3.16 as app-frontend

LABEL stage=app-frontend-build

WORKDIR /tmp/app-frontend

# Temporary copy placeholder index.html
COPY ./app-frontend /tmp/webapp

COPY ./app-frontend /tmp/app-frontend/

RUN echo $(ls -lai)

FROM golang:1.18-alpine as app-backend

LABEL stage=app-backend-build

RUN apk add --no-cache git

WORKDIR /tmp/app-go

COPY ./app-backend/go.mod .

COPY ./app-backend/go.sum .

RUN go mod download

COPY ./app-backend .

RUN go build -o ./out/app-go .

FROM alpine:3.16

COPY --from=app-frontend /tmp/webapp /app/webapp

COPY --from=app-backend /tmp/app-go/out/app-go /app/app-go

ENV FRONTEND_APP_SOURCES "webapp/dist/public"
ENV APPLICATION_PORT "8080"

EXPOSE 8080

WORKDIR /app

CMD ["./app-go"]