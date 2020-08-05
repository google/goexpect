#!/bin/sh
trap 'echo "Got the INTR Signal"; rm -f ~/fifo; trap - INT; kill -s INT "$$"' INT
trap 'echo "Got the QUIT Signal"; rm -f ~/fifo; trap - QUIT; kill -s INT "$$"' QUIT
trap 'echo "Got the USR1 Signal"; rm -f ~/fifo' USR1
trap 'echo "Got the HUP Signal"; rm -f ~/fifo' HUP
mkfifo -m 0400 ~/fifo || exit
echo "Waiting for signal"
true < ~/fifo 2> /dev/null
sleep 10
