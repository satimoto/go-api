#!/bin/sh
COMMIT_HASH=$(git log -1 --pretty=%h)
BRANCH=$(git rev-parse --abbrev-ref HEAD)
VERSION=$(echo ${BRANCH} | sed -e 's/.*\/v//g' | sed -e 's/\/.*//g')
UNCOMMITED_FILES=$(git status -s | wc -l | tr -d ' ')

usage()
{
    echo ""
    echo "Usage: tag [OPTIONS]"
    echo ""
    echo "Tags the docker image with the git branch version and git commit hash (e.g. 1.12.0-4e543431)"
    echo ""
    echo "Options:"
    echo "  -b, --build    Rebuild the docker image"
    echo "  -t, --testnet  Tag the image for testnet"
    echo "  -m, --mainnet  Tag the image for mainnet"
    echo "  -c, --commit   Push all tagged images"
    echo "  -f, --force    Forces continuation if there are uncommitted files, otherwise user is prompted"
    echo ""
}

while [ $# -ge 1 ]; do
    case $1 in
        -b|--build) BUILD="SET"            
        ;;
        -t|--testnet) TESTNET="SET"            
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

if [ -n "$BUILD" ]; then
    echo "Building api"
    docker build -t api .
fi

if [ -n "$TESTNET" ]; then
    docker tag api 526438337184.dkr.ecr.eu-west-1.amazonaws.com/api:testnet
    echo "Tagged 526438337184.dkr.ecr.eu-west-1.amazonaws.com/api:testnet"
fi

docker tag api 526438337184.dkr.ecr.eu-west-1.amazonaws.com/api:${VERSION}
echo "Tagged 526438337184.dkr.ecr.eu-west-1.amazonaws.com/api:${VERSION}"
docker tag api 526438337184.dkr.ecr.eu-west-1.amazonaws.com/api:${VERSION}-${COMMIT_HASH}
echo "Tagged 526438337184.dkr.ecr.eu-west-1.amazonaws.com/api:${VERSION}-${COMMIT_HASH}"

if [ -n "$MAINNET" ]; then
    docker tag api 526438337184.dkr.ecr.eu-west-1.amazonaws.com/api:mainnet
    echo "Tagged 526438337184.dkr.ecr.eu-west-1.amazonaws.com/api:mainnet"
fi

if [ -n "$COMMIT" ]; then
    if [ -n "$TESTNET" ]; then
        docker push 526438337184.dkr.ecr.eu-west-1.amazonaws.com/api:testnet
        echo "Pushed 526438337184.dkr.ecr.eu-west-1.amazonaws.com/api:testnet"
    fi

    docker push 526438337184.dkr.ecr.eu-west-1.amazonaws.com/api:${VERSION}
    echo "Pushed 526438337184.dkr.ecr.eu-west-1.amazonaws.com/api:${VERSION}"
    docker push 526438337184.dkr.ecr.eu-west-1.amazonaws.com/api:${VERSION}-${COMMIT_HASH}
    echo "Pushed 526438337184.dkr.ecr.eu-west-1.amazonaws.com/api:${VERSION}-${COMMIT_HASH}"

    if [ -n "$MAINNET" ]; then
        docker push 526438337184.dkr.ecr.eu-west-1.amazonaws.com/api:mainnet
        echo "Pushed 526438337184.dkr.ecr.eu-west-1.amazonaws.com/api:mainnet"
    fi
fi