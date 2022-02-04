# Wall of Globes

A video wall system based on VLC.

*You contaminated the wall of globes inside the Company, then went to the Source. The portal wall, the magic mirror that 
led back to where the Company came from. You let it trickle in, like a slow-acting poison that was actually. Life, again.* 

Dead Astonauts, Jeff Vandermeer

## Overview

Wall of Globes allows you to spread a single video stream across multiple displays. Currently, it assumes you have
multiple devices (one per display) capable of running a web server and VLC. You can then broadcast a UDP video stream
and coordinate the devices to player the correct fragment of the total video, making it appear to span over the many
displays.

This is still an early work-in-progress, written in Go as I learn the language.



