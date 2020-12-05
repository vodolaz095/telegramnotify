#!/usr/bin/env bash

# This script is started each time Transmission (https://transmissionbt.com/) downloads
# torrent to notify telegram group that we have downloaded it.
# You can run this script by setting parameters in ~/.config/transmission/settings.json like this:
#
# {
#    ....
#    "script-torrent-done-enabled": true,
#    "script-torrent-done-filename": "/usr/bin/transmission-complete-telegram.sh",
#    ....
# }
#
# See for details
# https://github.com/transmission/transmission/wiki/Configuration-Files
# https://github.com/transmission/transmission/wiki/Editing-Configuration-Files

/usr/bin/telegramnotify "We have downloaded torrent ${TR_TORRENT_NAME} into ${TR_TORRENT_DIR}!" torrents
