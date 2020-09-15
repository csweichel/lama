FROM golang:latest

LABEL maintainer="Peponi <pep0ni@pm.com>" \
      description="will setup the lama server on port 8080"

ARG REPO_URL="https://github.com/csweichel/lama"
ARG VERSION="v0.3.0"
ARG PROJECT_NAME="lama"
ARG USER="lama"
ARG PORT="8080"

ENV REPO_URL=$REPO_URL \
    VERSION=$VERSION \
    PROJECT_NAME=$PROJECT_NAME \
    USER=$USER

RUN set -e;\
  apk update; \
  apk upgrade; \
  apk add --no-cache --virtual \
    curl \
    git; \
  addgroup -g 12345 $USER;  \
  adduser -u 12345 -G $USER -D $USER; \
  mkdir -p /home/$USER/$PROJECT_NAME; \
  cd /home/$USER/$PROJECT_NAME; \
  git clone -b $VERSION --single-branch $REPO_URL .; \
  apk del git

EXPOSE $PORT

USER $USER

WORKDIR /home/$USER/$PROJECT_NAME

CMD go run main.go

HEALTHCHECK CMD curl --fail http://127.0.0.1:$PORT/ || exit 1