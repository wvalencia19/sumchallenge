# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

COPY bin/api ./api

ENV API_LOGLEVEL=debug
ENV API_JWT_TOKENTTLHOURS=1h
ENV API_JWT_SECRETKEY="secret"

EXPOSE 8080

ENTRYPOINT [ "./api" ]
