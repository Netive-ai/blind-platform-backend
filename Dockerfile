# build stage
FROM golang:alpine AS build-env
RUN apk --no-cache add build-base git bzr mercurial gcc
ADD . /src
RUN cd /src && go build cmd/platform/platform.go

# final stage
FROM alpine
RUN mkdir /app
RUN mkdir /app/cfg
COPY --from=build-env /src/platform /app/
COPY --from=build-env /src/cfg/ /app/cfg
WORKDIR /app
ENTRYPOINT ./platform