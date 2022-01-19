# The dockerfile is currently still WIP and might be broken
FROM golang:1.18-alpine AS build-env
RUN apk --no-cache add build-base git musl-dev linux-headers npm
WORKDIR /src/
ADD go.mod go.sum ./
RUN go mod download -x
ADD . ./
RUN make -B all

# final stage
FROM alpine
WORKDIR /usr/src/app
RUN apk --no-cache add libstdc++ libgcc
COPY --from=build-env /src/bin /usr/src/app/
COPY --from=build-env /src/phase0.yml /usr/src/app/phase0.yml
COPY ./config.yml /usr/src/app/config.yml
EXPOSE 3336
CMD ["./explorer", "--config", "config.yml"]
