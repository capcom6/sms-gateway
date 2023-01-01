#!/bin/bash

if [ -e "$(dirname $0)/env.sh" ]; then
    source "$(dirname $0)/env.sh"
fi

if [ -n "$BITBUCKET_TAG" ]; then
    GIT_TAG="$BITBUCKET_TAG"
else
    GIT_TAG=$(git describe --exact-match --tags $(git log -n1 --pretty='%h'))
fi

if [ -n "$BITBUCKET_BRANCH" ]; then
    GIT_BRANCH="$BITBUCKET_BRANCH"
else
    GIT_BRANCH=$(git rev-parse --abbrev-ref HEAD)
fi

TAG=${GIT_TAG:-v"${GIT_BRANCH////-}"}
if [ -z "$TAG" ]; then
    exit 0
fi

if [ -z "$APP" ]; then
    echo "Need APP environment var"
    exit 1
fi

DOCKER_IMAGE="${DOCKER_CR}/${APP}"

DOCKER_PARAMS=()
if [ -n "$DOCKER_FILE" ]; then
    DOCKER_PARAMS+=(--file "$DOCKER_FILE")
fi
if [ -n "$DOCKER_TARGET" ]; then
    DOCKER_PARAMS+=(--target "$DOCKER_TARGET")
fi

FULL_VERSION=${TAG:1}
PARTS=$(echo ${FULL_VERSION} | tr "." "\n")
VERSION=

set -e

docker build ${DOCKER_PARAMS[*]} -t ${DOCKER_IMAGE}:${FULL_VERSION} --build-arg SSH_PRV_KEY="$SSH_PRV_KEY" --build-arg APP="$APP" .
docker push ${DOCKER_IMAGE}:${FULL_VERSION}

for V in $PARTS
do
    if [ -n "$VERSION" ];
    then
        VERSION=${VERSION}.
    fi
    VERSION=${VERSION}${V}

    if [ "$FULL_VERSION" = "$VERSION" ];
    then
        continue
    fi
    
    docker tag ${DOCKER_IMAGE}:${FULL_VERSION} ${DOCKER_IMAGE}:${VERSION}
    docker push ${DOCKER_IMAGE}:${VERSION}
done

if [[ "$TAG" != "v${GIT_BRANCH}" ]]; then
    docker tag ${DOCKER_IMAGE}:${FULL_VERSION} ${DOCKER_IMAGE}:latest
    docker push ${DOCKER_IMAGE}:latest
fi
