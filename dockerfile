# Stage 1. Build the go webserver
FROM golang:alpine as stage1
WORKDIR /app
COPY ./go.mod ./
RUN go mod download
COPY ./.env ./
COPY ./main.go ./
COPY ./controller ./controller
COPY ./utils ./utils
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

#Stage2. Build the website
FROM node:12 as stage2
WORKDIR /app
COPY ./website/package.json ./
RUN npm install
COPY ./website .
RUN npm run build

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
# Copy built binary
COPY --from=stage1 /app/main .
# Copy built website
COPY --from=stage2 /app/build ./build
EXPOSE 8000
CMD ["./main"]