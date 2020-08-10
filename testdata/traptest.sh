#!/bin/sh
FIFOFILE=~/fifo
trap 'rm -f $FIFOFILE; echo "Got the INTR Signal"' INT
trap 'rm -f $FIFOFILE; echo "Got the QUIT Signal"' QUIT
trap 'rm -f $FIFOFILE; echo "Got the USR1 Signal"' USR1
trap 'rm -f $FIFOFILE; echo "Got the HUP Signal"' HUP

if [ -f $FIFOFILE ]
then
  rm -f $FIFOFILE
fi

mkfifo -m 0400 $FIFOFILE || exit
echo "Waiting for signal"
true < $FIFOFILE 2> /dev/null
sleep 10
