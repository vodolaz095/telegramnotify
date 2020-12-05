#!/usr/bin/env bash

UNIT=$1

UNITSTATUS=$(systemctl status $UNIT)
ALERT=$(echo -e "\u26A0")

/usr/bin/telegramnotify send "$ALERT Unit failed $UNIT $ALERT
Status:
$UNITSTATUS" work
