# watch

```
$> rpi-birdfeeder watch
```

This is the main process designed for the rpi-birdfeeder. This processes listens for incoming signals from an IR motion detector sensor. Whenever motion is detected, it triggers the rpi camera to capture a series of images.

----

### Parameters

The parameters for the Watch sub-process are stored in the */etc/rpi-birdfeeder/config.json* main application configuration object.

----

### Examples

```
rpi-birdfeeder --config /etc/rpi-birdfeeder/config.json watch
```