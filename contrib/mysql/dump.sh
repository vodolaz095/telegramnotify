#!/usr/bin/env bash

# This script makes dump of database and uploads it to telegram group
# but it cannot be considered reliable long time storage for backups!
# We assume you have set up ~/.my.cnf with database credentials properly

DATE=$(date +%d-%m-%Y)
DATABASE=blog # name of database you want to backup

/usr/bin/mysqldump $DATABASE | gzip > "/tmp/$DATABASE_$DATE.sql.gz"
/usr/bin/telegramnotify share "/tmp/$DATABASE_$DATE.sql.gz" oldcity
rm -f "/tmp/$DATABASE_$DATE.sql.gz" # because leaving clutter is a sin of pride!
