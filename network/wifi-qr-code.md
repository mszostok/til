# Generate QR code for your WiFi login details

You can generate these codes using an online service, or using the QR generator of your choice and feeding it this text:
<!--more-->

```text
WIFI:S:My_SSID;T:WPA;P:key goes here;H:false;
^    ^         ^     ^               ^
|    |         |     |               +-- hidden SSID (true/false)
|    |         |     +-- WPA key
|    |         +-- encryption type
|    +-- ESSID
+-- code type
```

Escape any special characters (", ', ;, ,, or \) with a backslash.

_source: https://superuser.com/a/1459233_

_source: https://github.com/bndw/wifi-card_
