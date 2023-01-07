# build stage
FROM golang:1.19-alpine AS build-env
ARG APP_NAME=flyer

RUN apk add --no-cache curl bash git openssh
    
COPY . /src/flyer
WORKDIR /src/flyer
RUN go mod download && go build -o $APP_NAME

# final stage
FROM alpine:3.17
RUN apk -U add ca-certificates

WORKDIR /app
COPY --from=build-env /src/flyer/$APP_NAME /app/
COPY --from=build-env /src/flyer/api /app/api

EXPOSE 9080

CMD [ "/app/flyer" ]