#!/bin/bash
FIFOFILE=$(mktemp -u)
trap 'rm -f $FIFOFILE; echo "Got the INTR Signal"' INT
trap 'rm -f $FIFOFILE; echo "Got the QUIT Signal"' QUIT
trap 'rm -f $FIFOFILE; echo "Got the USR1 Signal"' USR1
trap 'rm -f $FIFOFILE; echo "Got the HUP Signal"' HUP

if [[ -p $FIFOFILE ]]
then
  rm -f $FIFOFILE
fi

mkfifo -m 0400 $FIFOFILE
echo "Waiting for signal"
true < $FIFOFILE
sleep 10
