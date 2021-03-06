#!/bin/sh
cd `dirname $0` || exit

CUR_DIR=`pwd`
LIB_DIR=${CUR_DIR}/lib

if [ -z "${LD_LIBRARY_PATH}" ]; then
    export LD_LIBRARY_PATH="${LIB_DIR}"
else
    export LD_LIBRARY_PATH="${LIB_DIR}:${LD_LIBRARY_PATH}"
fi

APP_NAME=bin/novelserver
LOAD_TOOL=bin/supervise.novelserver
LOAD_DIR=status/novelserver

mkdir -p log
mkdir -p ${LOAD_DIR}

listApps() {
    ps -ef | grep ${APP_NAME} | grep -v grep
    ps -ef | grep ${LOAD_TOOL} | grep -v grep | grep -v ${APP_NAME}
}

start() {
    ${LOAD_TOOL} -u ${LOAD_DIR} ${APP_NAME}
    listApps
    echo "End"
}

stop() {
    if [ -f ${LOAD_DIR}/lock ]; then
        tool_pid=`/sbin/fuser ${LOAD_DIR}/lock`
        #tool_pid=`fuser ${LOAD_DIR}/lock`
        `ps -ef | grep "$tool_pid" | grep "${LOAD_TOOL}" | grep -v grep > /dev/null 2>&1`
        if [ $? -eq 0 ] && [ "$tool_pid" != "" ] ; then
            echo "kill ${LOAD_TOOL} process:"${tool_pid}
            kill -9 $tool_pid
        fi
    fi                                                                                                      

    if [ -f ${LOAD_DIR}/status ]; then                                                       
        app_pid=`od -An -j16 -N2 -tu2 ${LOAD_DIR}/status`                    
        `ps -ef | grep "$app_pid" | grep "${APP_NAME}" | grep -v grep > /dev/null 2>&1`
        if [ $? -eq 0 ]; then
            echo "kill ${APP_NAME} process:"${app_pid}                              
            kill -9 $app_pid 
        fi
    fi

    listApps
    echo "End"
}

other() {
    echo "Usage: $0 {start|stop|list}"
}

case "$1" in
    start)
        start
        ;;
    stop)
        stop
        ;;
    list)
        listApps
        ;;
    restart)
        stop
        start
        ;;
    *)
        other
        ;;
esac
