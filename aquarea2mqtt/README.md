# Home Assistant Add-on: Aquarea2MQTT

Panasonic Aquarea Service Cloud to MQTT gateway

![Current version][version] ![Project Stage][project-stage-shield]

![Supports aarch64 Architecture][aarch64-shield]
![Supports amd64 Architecture][amd64-shield]
![Supports armhf Architecture][armhf-shield]
![Supports armv7 Architecture][armv7-shield]
![Supports i386 Architecture][i386-shield]

**This addon is deprecated**

## Important information

- When you set the interval < 60s you have the chance that Panasonic will block you temporarily.
- All sensors and switches are automatically discovered by Home Assistant MQTT

## Installation

Note: This addon requires an mqtt broker. Make sure you have this already running.

1. [Add my addon repository](https://github.com/fbloemhof/hassio-addons) to your Home Assistant instance
2. Install the addon
3. Configure your credentials to the Aquarea Service Cloud and save the config
4. Start the addon
5. Check the logfile to see if the addon is running fine.

If all went well you can find the Aquarea device automatically discovered under the MQTT integrations

## Credits

This addon is based on <https://github.com/rondoval/aquarea2mqtt> which is based on <https://github.com/lsochanowski/Aquarea2mqtt>. Many thanks to both.

## Disclaimer

I am not related to Panasonic and can not take any responsibility in issues you may cause with the Service Cloud or your Panasonic heatpump when using this addon.

[aarch64-shield]: https://img.shields.io/badge/aarch64-yes-green.svg
[amd64-shield]: https://img.shields.io/badge/amd64-yes-green.svg
[armhf-shield]: https://img.shields.io/badge/armhf-yes-green.svg
[armv7-shield]: https://img.shields.io/badge/armv7-yes-green.svg
[i386-shield]: https://img.shields.io/badge/i386-yes-green.svg
[project-stage-shield]: https://img.shields.io/badge/project%20stage-experimental-yellow.svg
[version]: https://img.shields.io/badge/version-v2023.1.15.9-blue.svg
