Systemd integration
==================================

Notify via telegram, if systemd unit fails.

Inspiration - https://www.curry-software.com/en/blog/telegram_unit_fail/

How to install

1. Copy `telegramnotify` binary to `/usr/bin/`
2. Copy script `unit-status-telegram.sh` to `/usr/bin/`
3. Copy `unit-status-telegram@.service` to `/etc/systemd/system`
4. Add string `OnFailure=unit-status-telegram@%n.service` to systemd unit file you want to monitor, like this one:


```

[Unit]
Description=Redis persistent key-value database
After=network.target
After=network-online.target
Wants=network-online.target

[Service]
ExecStart=/usr/bin/redis-server /etc/redis.conf --supervised systemd
ExecStop=/usr/libexec/redis-shutdown
Type=notify
User=redis
Group=redis
RuntimeDirectory=redis
RuntimeDirectoryMode=0755

# Notify, when server fails!
OnFailure=unit-status-telegram@%n.service

[Install]
WantedBy=multi-user.target

```
