#!/bin/sh

/app/goose up \
    && /app/app $@
