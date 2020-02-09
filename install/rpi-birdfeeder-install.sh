#!/bin/bash

wget --directory-prefix=/usr/bin --output-document=rpi-birdfeeder https://github.com/littlehawk93/rpi-birdfeeder/tree/install/install/rpi-birdfeeder 

mkdir /etc/rpi-birdfeeder

if [ -f "/etc/rpi-birdfeeder/config.json"]; then
    echo "Config file already exists"
else
    wget --directory-prefix=/etc/rpi-birdfeeder --output-document=config.json https://github.com/littlehawk93/rpi-birdfeeder/tree/install/install/default-config.json
fi

wget --directory-prefix=/etc/systemd/system --output-document=feeder.powermon.service https://github.com/littlehawk93/rpi-birdfeeder/tree/install/install/feeder.powermon.service
wget --directory-prefix=/etc/systemd/system --output-document=feeder.watch.service https://github.com/littlehawk93/rpi-birdfeeder/tree/install/install/feeder.watch.service

systemctl enable feeder.powermon
systemctl enable feeder.watch
