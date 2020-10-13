# DID Sidechain - elastos.org
# This is an official but unsupported node image

FROM golang:1.15.2-alpine3.12 AS builder

LABEL maintainer="kpachhai"

ENV SRC_DIR="/elastos"

RUN apk update
RUN apk add --no-cache make
RUN apk add --no-cache git
RUN apk add --no-cache gcc
RUN apk add --no-cache libc-dev

ARG REPO_URL=https://github.com/elastos/Elastos.ELA.SideChain.ID.git
ARG REPO_BRANCH=v0.2.0

# Clone repo
RUN git clone -b ${REPO_BRANCH} ${REPO_URL} ${SRC_DIR}

WORKDIR ${SRC_DIR}

# Build
RUN make

#---------------- Multi-container build ------------------##

FROM alpine:3.12

ENV SRC_DIR="/elastos"

COPY --from=builder ${SRC_DIR}/did ${SRC_DIR}/did

RUN apk update \
    && apk add --no-cache curl ca-certificates \
    && addgroup -g 1000 -S elauser \
    && adduser -h $SRC_DIR -u 1000 -S elauser -G elauser \
    && chown -R elauser:elauser $SRC_DIR

USER elauser

WORKDIR ${SRC_DIR}

EXPOSE 20604-20608

ENTRYPOINT ["/bin/sh", "-c", "./did"]
