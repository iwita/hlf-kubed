#!/bin/bash

INTERFACE="eth0"
DEVICE="Device0"


interface_bytes_in_old=$(awk "/^ *${INTERFACE}:/"' { if ($1 ~ /.*:[0-9][0-9]*/) { sub(/^.*:/, "") ; print $1 } else { print $2 } }' /proc/net/dev)
interface_bytes_out_old=$(awk "/^ *${INTERFACE}:/"' { if ($1 ~ /.*:[0-9][0-9]*/) { print $9 } else { print $10 } }' /proc/net/dev)
# start=`date +%s`
idle_cpu_old=$(mpstat | tail -n 1 | awk '{print $12}')
#echo 
used_cpu_old=$(echo "100 - $idle_cpu_old" | bc -l )
#echo "Used CPU" $USED_CPU

used_mem_old=$(free -m | grep Mem | awk '{print $3}')
#echo "Used mem" $USED_MEM

#ecoo "$interface_bytes_in_old,$interface_bytes_out_old,$used_mem_old,$used_cpu_old"

while true
do
    start=$(($(date +%s%N)/1000000))
   # echo $start
    idle_cpu_new=$(mpstat | tail -n 1 | awk '{print $12}')
    
    used_cpu_new=$(echo "scale=1; 100 - $idle_cpu_new" | bc -l)
    # echo "Used CPU" $USED_CPU

    used_mem_new=$(free -m | grep Mem | awk '{print $3}')
    # echo "Used mem" $USED_MEM

    interface_bytes_in_new=$(awk "/^ *${INTERFACE}:/"' { if ($1 ~ /.*:[0-9][0-9]*/) { sub(/^.*:/, "") ; print $1 } else { print $2 } }' /proc/net/dev)
    interface_bytes_out_new=$(awk "/^ *${INTERFACE}:/"' { if ($1 ~ /.*:[0-9][0-9]*/) { print $9 } else { print $10 } }' /proc/net/dev)

    

    # Calculate CPU output
    sum_cpu=$( echo "scale=2; $used_cpu_new + $used_cpu_old" | bc -l )
    used_cpu=$( echo "scale=2; $sum_cpu/2" | bc -l )

    # Calculate Memory output
    sum_mem=$( echo "scale=2; $used_mem_new + $used_mem_old" | bc -l )
    used_cpu=$( echo "scale=2; $sum_mem/2" | bc -l )

    # Calculate Network bandwidth output
    temp_out=$( echo "scale=1; $interface_bytes_out_new - $interface_bytes_out_old" | bc -l)
    temp_in=$( echo "scale=1; $interface_bytes_in_new - $interface_bytes_in_old" | bc -l )
    
    end=$(($(date +%s%N)/1000000))
    runtime_ms=$((end-start))

    runtime=$( echo "scale=3; 1+$runtime_ms/1000" | bc -l )
    netw_out=$( echo "scale=1; $temp_out/$runtime" | bc -l )
    netw_in=$( echo "scale=1; $temp_in/$runtime" | bc -l )

#    echo "$DEVICE,$used_cpu,$user_mem,$netw_out"
    # Store in ledger
    sed  -e "s/__device__/$DEVICE/g" -e "s/__ipc__/$used_cpu/g" -e "s/__mem__/$used_mem/g" -e "s/__bw__/$netw_out/g" template.sh | bash
    end=$(($(date +%s%N)/1000000))
    runtime_ms=$((end-start))

    runtime=$( echo "scale=3; 1+$runtime_ms/1000" | bc -l )
    echo "runtime $runtime"

    # TODO: Sleep 10 seconds for testing
    # sleep $((300-runtime))
    echo "sleep for $(echo "scale=2; 5-$runtime" | bc -l)"
    sleep $( echo "scale=2; 5-$runtime" | bc -l )
done

