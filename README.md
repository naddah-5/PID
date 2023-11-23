# PID
PID control system for a water tank


## Set up

### Installing an image
Since the revpi ha eMMC you will have to install some additional software in order to install an image on it, the instructions can be found [here](https://www.youtube.com/watch?v=jp_mF1RknU4&t=81s).

After you have completed the above steps you can check out this [page](https://revolutionpi.com/tutorials/downloads#revpiimages) to find an image for the device. If you are not sure which image is compatible with your device you can check find out [here](https://revolutionpi.com/tutorials/images). You should now be ready to install the image to the revpi device. The first step in this process is to connect the device to your pc via the micro-usb port while the device is *off*. You can then connect the power supply which causes the device to boot in a mode that allows us to load a new image to the device.

### Connect to the device
The first step in connecting to the revpi unit is to connect it to a local network. In most cases this can be done by simply connecting it via ethernet to the local router. However, when connecting via institution networks you may need to set up a permitted connection for the device. This is usually done by pairing the mac address to a specific IP, which also has the benefit of a static IP address which will be convenient for ssh. For the revpi connect there are two mac addresses available (A and B), they are printed on the units chassi below the ethernet port and are slightly different. For more information on the difference follow [this link](https://revolutionpi.com/tutorials/ethernet-ports-compact).
