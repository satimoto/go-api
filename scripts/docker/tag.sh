#!/bin/sh
COMMIT_HASH=$(git log -1 --pretty=%h)
BRANCH=$(git rev-parse --abbrev-ref HEAD)
VERSION=$(echo ${BRANCH} | sed -e 's/.*\/v//g' | sed -e 's/\/.*//g')
UNCOMMITED_FILES=$(git status -s | wc -l | tr -d ' ')
ACCOUNT_ID="909899099608"
NETWORK="testnet"

usage()
{
    echo ""
    echo "Usage: tag [OPTIONS]"
    echo ""
    echo "Tags the docker image with the git branch version and git commit hash (e.g. 1.12.0-4e543431)"
    echo ""
    echo "Options:"
    echo "  -b, --build    Rebuild the docker image"
    echo "  -t, --tag      Tag the image with network"
    echo "  -m, --mainnet  Switch to mainnet"
    echo "  -c, --commit   Push all tagged images"
    echo "  -f, --force    Forces continuation if there are uncommitted files, otherwise user is prompted"
    echo ""
}

while [ $# -ge 1 ]; do
    case $1 in
        -b|--build) BUILD="SET"            
        ;;
        -t|--tag) TAG="SET"            
        ;;
        -m|--mainnet) MAINNET="SET"            
        ;;
        -c|--commit) COMMIT="SET"            
        ;;
        -f|--force) FORCE="SET"            
        ;;
        -h|--help) usage
                   exit
                   ;;
        *) usage
           exit 1
    esac
    shift
done

if [ "$UNCOMMITED_FILES" -ne "0" ] && [ -z "$FORCE" ]; then
    read -p "There are ${UNCOMMITED_FILES} uncommitted file(s). Continue? [Y/n]: " CONFIRM

    if [ "$CONFIRM" != "Y" ]; then
        exit 1
    fi
fi

if [ -z "$VERSION" ]; then
    VERSION=${BRANCH}
fi

if [ -n "$MAINNET" ]; then
    ACCOUNT_ID="490833747373"
    NETWORK="mainnet"
fi

if [ -n "$BUILD" ]; then
    echo "Building api"
    docker build -t api .
fi

docker tag api ${ACCOUNT_ID}.dkr.ecr.eu-central-1.amazonaws.com/api:${VERSION}
echo "Tagged ${ACCOUNT_ID}.dkr.ecr.eu-central-1.amazonaws.com/api:${VERSION}"
docker tag api ${ACCOUNT_ID}.dkr.ecr.eu-central-1.amazonaws.com/api:${VERSION}-${COMMIT_HASH}
echo "Tagged ${ACCOUNT_ID}.dkr.ecr.eu-central-1.amazonaws.com/api:${VERSION}-${COMMIT_HASH}"

if [ -n "$TAG" ]; then
    docker tag api ${ACCOUNT_ID}.dkr.ecr.eu-central-1.amazonaws.com/api:${NETWORK}
    echo "Tagged ${ACCOUNT_ID}.dkr.ecr.eu-central-1.amazonaws.com/api:${NETWORK}"
fi

if [ -n "$COMMIT" ]; then
    docker push ${ACCOUNT_ID}.dkr.ecr.eu-central-1.amazonaws.com/api:${VERSION}
    echo "Pushed ${ACCOUNT_ID}.dkr.ecr.eu-central-1.amazonaws.com/api:${VERSION}"
    docker push ${ACCOUNT_ID}.dkr.ecr.eu-central-1.amazonaws.com/api:${VERSION}-${COMMIT_HASH}
    echo "Pushed ${ACCOUNT_ID}.dkr.ecr.eu-central-1.amazonaws.com/api:${VERSION}-${COMMIT_HASH}"

    if [ -n "TAG" ]; then
        docker push ${ACCOUNT_ID}.dkr.ecr.eu-central-1.amazonaws.com/api:${NETWORK}
        echo "Pushed ${ACCOUNT_ID}.dkr.ecr.eu-central-1.amazonaws.com/api:${NETWORK}"
    fi
fi