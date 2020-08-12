#!/bin/bash
FIFOFILE=~/fifo
trap 'rm -f $FIFOFILE; echo "Got the INTR Signal"' INT
trap 'rm -f $FIFOFILE; echo "Got the QUIT Signal"' QUIT
trap 'rm -f $FIFOFILE; echo "Got the USR1 Signal"' USR1
trap 'rm -f $FIFOFILE; echo "Got the HUP Signal"' HUP

[[ -f "$FIFOFILE" ]] && rm -f $FIFOFILE

mkfifo -m 0400 $FIFOFILE || exit
echo "Waiting for signal"
cat < $FIFOFILE || echo ""
sleep 10
