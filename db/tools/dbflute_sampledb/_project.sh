#!/bin/bash

export ANT_OPTS=-Xmx512m

export DBFLUTE_HOME=../mydbflute/dbflute-1.1.6

export MY_PROPERTIES_PATH=build.properties

if [ `uname` = "Darwin" ]; then
  export JAVA_HOME=$(/usr/libexec/java_home -v 1.8)
fi
