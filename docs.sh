#!/bin/bash

# Check availability of docker
hash docker 2>/dev/null || { echo >&2 "* docker is required!"; exit 1; }

# Check availability of docker-compose
hash docker-compose 2>/dev/null || { echo >&2 "* docker-compose is required!"; exit 1; }

function net_port_in_use() {
    port=$1
    nc -z 127.0.0.1 $port 2>/dev/null
    if [ $? -eq 0 ]; then
        return 0
    else
        return 1
    fi

}

function question_ask_default() {

    if [ "$2" = "" ]; then
        echo "Error Default value is missing"
        exit;
    fi;

    read -p "$1" response

    if [ "$response" = "" ]; then
        echo "$2"
    else
        echo $response
    fi;
}

function printLine(){
  echo "$1"
}

function usage() {

  TITLE=$(cat docs/mkdocs.yml | grep "site_name" | sed -e 's/site_name: //g')
  DESCRIPTION=$(cat docs/mkdocs.yml | grep "site_description" | sed -e 's/site_description: //g')

  printLine "Documentation: $TITLE"
  printLine "$DESCRIPTION"
  printLine ""
  printLine "Usage: sh docs.sh [start|stop]"
  printLine "Options:"
  printLine "   start: load documentation app at port 8000 by default"
  printLine "   stop: stop documentation app"

}

case "$1" in
    start)
    shift 1
    
    while : ; do
        dsMkdocsPort=$( question_ask_default "Documentation default port? [8000] " "8000" )

        if ! net_port_in_use $dsMkdocsPort ; then
            echo "The selected port is: $dsMkdocsPort"
            echo $dsMkdocsPort > docs.port
            break
        else
            echo "---> The port $dsMkdocsPort is busy, please select a new one\n"
        fi;
    done

    export DS_MKDOCS_SERVICENAME=$(basename $(pwd))
    export DS_MKDOCS_PORT=$dsMkdocsPort
    docker-compose -p $DS_MKDOCS_SERVICENAME -f docs/docker-compose-docs.yml up -d --build
    sleep 5
    open "http://localhost:$dsMkdocsPort"
    ;;

    stop)
    shift 1
    #Default port only to prevent syntax errors
    export DS_MKDOCS_SERVICENAME=$(basename $(pwd))
    export DS_MKDOCS_PORT=$(head -n 1 docs.port)
    docker-compose -p $DS_MKDOCS_SERVICENAME -f docs/docker-compose-docs.yml down --remove-orphans
    ;;

    *)
    usage
    ;;
esac

exit 0



