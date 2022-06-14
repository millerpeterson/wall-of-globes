# Wall of Globes

A video wall system based on VLC.

*You contaminated the wall of globes inside the Company, then went to the Source. The portal wall, the magic mirror that 
led back to where the Company came from. You let it trickle in, like a slow-acting poison that was actually. 
Life, again.* 

Dead Astonauts, Jeff Vandermeer

## Overview

Wall of Globes allows you to tile a video stream across multiple displays. It assumes you have multiple devices, each
attached to their own display, capable of running a web server and VLC. You can then broadcast a video stream to the
display devices, and then tell them (via HTTP request) to play a crop of the total video. With each display acting as a
tile of the total picture, the stream appears to cover the display array.

Currently, the display servers are assumed to be running on Raspberry Pi 4's or OSX. The application is written in Go, 
and so could be compiled on many platforms, but it has only been tested on Raspberry Pi 4 and OSX.

The app consists of three executables:

### globe 

The display server that will play a crop of the video stream.

### company

Broadcasts the video stream, and tells globes what to play.

### multiply

Utilities for setting up a number of globe's at once.

## Setting up a Globe display server on a Raspberry Pi

### Install the Globe binary

From the device(s) that will run the Globe display client:

1. Download the latest Globe binary for Pi's:
```shell
curl https://github.com/millerpeterson/wall-of-globes/releases/download/latest/globe-linux-arm -O ~/globe
```

2. Copy the startup item into place, so that the application runs when the Pi boots:
```shell
curl https://raw.githubusercontent.com/millerpeterson/wall-of-globes/main/globe.desktop -O globe.desktop
cp globe.desktop /home/pi/.config/autostart/
```

3. Reboot the Pi:
```shell
sudo shutdown -r now
```

### Verify that the Globe server is running: 

1. Get the Pi's IP address:
```shell
ifconfig
```

Make a note of this as you will need it when creating your tiling config.

2. Issue an HTTP request to the Globe server's status endpoint:
```shell
curl http://<pi's IP address>:8081/status
```

If you see a 200 OK status, the device is ready to play streams!

## Running a Stream

## HOWTO

### Routing the Multicast Stream to a Specific Interface

If the video appears choppy, or stutters, it may be that there is insufficient bandwidth on your network. If you are
trying to stream over wifi, you might consider running over Ethernet instead. 

To route the multicast packets using 225.0.0.1 as the multicast address through a specific network interface:

OSX:
```shell
sudo route add -net 225.0.0.0/8 -interface en6
```

Raspbian:
```shell
sudo route add -net 225.0.0.0 netmask 255.0.0.0 eth0
```

Where `en6` and `eth0` are the Ethernet network interfaces.