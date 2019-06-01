#!/bin/bash

CURDIR=$(cd $(dirname $0); pwd)
PORT=$1
RUNTIME_ROOT=${CURDIR}
RUNTIME_CONF_ROOT=$RUNTIME_ROOT/conf
RUNTIME_LOG_ROOT=$RUNTIME_ROOT/log

BinaryName="DietRegimen"
export GIN_LOG_DIR=$RUNTIME_LOG_ROOT
CONF_DIR=$CURDIR/conf/
args+=" -port=$PORT"

echo "$CURDIR/bin/${BinaryName} $args"
exec $CURDIR/bin/${BinaryName} $args