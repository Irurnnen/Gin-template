FROM golang:1.24.2-alpine AS build-stage

WORKDIR /app

# Installing dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy all go files
COPY . .

# Set build tag
ARG BUILD_TAG

RUN CGO_ENABLED=0 GOOS=linux go build --tags ${BUILD_TAG} --ldflags="-s -w" -buildvcs=false -o /app/gin-template ./cmd/


FROM alpine:3.21.3 AS production-stage

WORKDIR /app
COPY --from=build-stage /app/gin-template /app/gin-template

CMD [ "./gin-template" ]