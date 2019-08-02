#!/bin/bash
# a simple script to launch at login, change my background, and loop forever.
#  

trap cleanuplock 1 2 3 6 14 15

cleanuplock()
{
        rm -rf /tmp/temp.wc
        # kill $(jobs -p)
        exit 1
}

if [ -f /tmp/temp.wc ]
then
        # already a lockfile/process running
        exit 1
else 
        touch /tmp/temp.wc
fi

# give gnome-session a little time to start up
sleep 1m & wait;

while true
do
        DIR="/home/tccopela/Pictures/wallpaper"
        PIC=$(ls $DIR/* | shuf -n1)
        gsettings set org.gnome.desktop.background picture-uri "file://$PIC"

        sleep 10m & wait;
done


