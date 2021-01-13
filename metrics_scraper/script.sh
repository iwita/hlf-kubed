#!/bin/bash

INTERFACE="eth0"
DEVICE="Device0"

while true
#for i in 1 2
do
    start=`date +%s`
    IDLE_CPU=$(mpstat | tail -n 1 | awk '{print $12}')
    echo 
    USED_CPU=$(bc <<< "100 - $IDLE_CPU")
    echo "Used CPU" $USED_CPU

    USED_MEM=$(free -m | grep Mem | awk '{print $3}')
    echo "Used mem" $USED_MEM
    
    LINE=`grep "$INTERFACE" /proc/net/dev | sed s/.*://`
    RECEIVED=`echo $LINE | awk '{print $1}'`
    TRANSMITTED=`echo $LINE | awk '{print $9}'`
    USED_BW=$(($RECEIVED+$TRANSMITTED))
    echo "USED BW" $USED_BW

    # Store in ledger
    sed  -e "s/__device__/$DEVICE/g" -e "s/__ipc__/$USED_CPU/g" -e "s/__mem__/$USED_MEM/g" -e "s/__bw__/$USED_BW/g" template.sh | bash

    end=`date +%s`
    runtime=$((end-start))
    echo "runtime $runtime"

    # TODO: Sleep 10 seconds for testing
    # sleep $((300-runtime))
    # sleep $((10-runtime))
done
