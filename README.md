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

Assuming you have your Globe displays servers up and running, you can use Company to 
broadcast a stream for them to play.

1. Prepare the video file you would like to stream. Some large files might not 
perform well - try a lower-resolution / bitrate version if you experience choppiness.

2. Prepare the wall configuration file. There is an example of a 4x4 C64 layout here: 
[wall-layouts/4x4_64s.json](https://github.com/millerpeterson/wall-of-globes/blob/main/wall-layouts/4x4_c64s.json). You will need to assign the 
IP addresses of the Pi's in the `server_map` section.

3. Download the latest company for your system from the [releases page](https://github.com/millerpeterson/wall-of-globes/releases).

4. Invoke `company`:
```shell
company wall_config.json video_file.mp4
```
The globe should begin playing the stream. Quit the app (control-c) to stop the stream. The globe servers will remain
running, so you should be able to just run Company again.

## HOWTO

### Routing the Multicast Stream to a Specific Interface

If the video appears choppy, or stutters, it may be that there is insufficient bandwidth on your network. If you are
trying to stream over wifi, you might consider running over Ethernet instead. 

To route the multicast packets using 225.0.0.1 (the address used by this application) as the multicast address
trough a specific network interface:

#### Temporary Routing Change

The effect of the following commands will be lost when the machine reboots:

##### OSX

```shell
sudo route add -net 225.0.0.0/8 -interface en6
```

##### Raspbian

```shell
sudo route add -net 225.0.0.0 netmask 255.0.0.0 eth0
```

Where `en6` and `eth0` are the Ethernet network interfaces.

#### Persisted Routing Change

The following commands should make changes that persist through reboots:

##### OSX

1. Find the name of your desired device through `networksetup`:
```shell
networksetup -listallnetworkservices
```
Note the device's router IP.

2. Set the new route:
```shell
networksetup -setadditionalroutes 225.0.0.0 255.0.0.0 192.168.0.1
```
Assuming 192.168.0.1 is the router IP you noted in 1.

##### Raspbian
1. Run `ifconfig`, and note your desired network device's router IP.

2. Add or modify `/lib/dhcpcd/dhcpcd-hooks/40-route`, adding the following line:
```
ip route add 225.0.0.0/8 via 192.168.0.1
``` 
Again assuming that 192.168.0.1 is the router IP you noted in 1.