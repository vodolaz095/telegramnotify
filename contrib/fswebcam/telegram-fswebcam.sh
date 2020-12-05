#!/usr/bin/env bash

DATE=$(date +%d-%m-%Y)

# Make image from webcam and upload it to channel, can be started by cron/systemd-timer daily

/usr/bin/fswebcam --save /home/shared/Webcam_Pictures/img_$DATE.png
/usr/bin/telegramnotify upload /home/shared/Webcam_Pictures/img_$DATE.png
