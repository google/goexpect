#!/bin/bash
FIFOFILE=$(mktemp -u)
trap 'rm -f $FIFOFILE; echo "Got the INTR Signal"; exit' INT
trap 'rm -f $FIFOFILE; echo "Got the QUIT Signal"; exit' QUIT
trap 'rm -f $FIFOFILE; echo "Got the USR1 Signal"; exit' USR1
trap 'rm -f $FIFOFILE; echo "Got the HUP Signal"; exit' HUP

if [[ -p $FIFOFILE ]]
then
  rm -f $FIFOFILE
fi

mkfifo -m 0400 $FIFOFILE
echo "Waiting for signal"
true < $FIFOFILE
sleep 10
