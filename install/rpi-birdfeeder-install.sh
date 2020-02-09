#!/bin/bash

wget --output-document="/usr/bin/rpi-birdfeeder" https://raw.githubusercontent.com/littlehawk93/rpi-birdfeeder/install/install/rpi-birdfeeder

if [ -f "/etc/rpi-birdfeeder" ] && [ -d "/etc/rpi-birdfeeder" ]; then
    echo "Config folder already exists"
else
    mkdir /etc/rpi-birdfeeder
fi

if [ -f "/etc/rpi-birdfeeder/config.json" ]; then
    echo "Config file already exists"
else
    wget --output-document="/etc/rpi-birdfeeder/config.json" https://raw.githubusercontent.com/littlehawk93/rpi-birdfeeder/install/install/default-config.json
fi

wget --output-document="/etc/systemd/system/feeder.powermon.service" https://raw.githubusercontent.com/littlehawk93/rpi-birdfeeder/install/install/feeder.powermon.service
wget --output-document="/etc/systemd/system/feeder.watch.service" https://raw.githubusercontent.com/littlehawk93/rpi-birdfeeder/install/install/feeder.watch.service

systemctl enable feeder.powermon
systemctl enable feeder.watch
