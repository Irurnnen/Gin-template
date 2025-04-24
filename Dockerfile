FROM golang:1.24.2-alpine AS build-stage

WORKDIR /app

# Installing dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy all go files
COPY . .

# Set build tag
ARG BUILD_TAG

ENV GOCACHE=/root/.cache/go-build

# --tags ${BUILD_TAG}  -buildvcs=false CGO_ENABLED=0 
RUN --mount=type=cache,target="/root/.cache/go-build" GOOS=linux go build --ldflags="-s -w" -o /app/gin-template ./cmd/


FROM alpine:3.21.3 AS production-stage

WORKDIR /app
COPY --from=build-stage /app/gin-template /app/gin-template

CMD [ "./gin-template" ]