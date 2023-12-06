#!/bin/sh

/app/app db:migrate up \
    && /app/app $@
