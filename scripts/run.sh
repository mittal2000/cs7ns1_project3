#!/usr/bin/env bash
# Swati Poojary
main() {
    rm -r $1

    # gateway
    mkdir -p $1
    cp ./gateway/* $1
    nohup $1/gateway --dir $1 \
        --host $2 --port 33000 \
        --subhost 127.0.0.1 --subport 8443 \
        --index $3 > $1/log.txt 2>&1 &

    for (( dev=1; dev<=5; dev++ ))
    do
        # device
        mkdir -p $1/device$dev
        cp ./device/* $1/device$dev
        nohup $1/device$dev/device --dir $1/device$dev  \
            --host 127.0.0.$(($dev+1)) --port 8443 \
            --subhost 127.0.$dev.1 --subport 8443 \
            --index 127.0.0.1 > $1/device$dev/log.txt 2>&1 &

        for (( sensor=1; sensor<=8; sensor++ ))
        do
            # sensor
            mkdir -p $1/device$dev/sensor$sensor
            cp ./sensor/* $1/device$dev/sensor$sensor
            nohup $1/device$dev/sensor$sensor/sensor --dir $1/device$dev/sensor$sensor  \
                --host 127.0.$dev.$(($sensor+1)) --port 8443 \
                --index 127.0.$dev.1 > $1/device$dev/sensor$sensor/log.txt 2>&1 &
        done
    done

}

main $1 $2 $3
