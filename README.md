rarch-remote
============

HTTP-to-Retroarch bridge.


## Purpose

To present a mobile-friendly HTML interface to RetroArch's UDP-based Network Commands.

The idea is to let you:

* Navigate the RetroArch menu&mdash;load cores, load roms, configure input, etc.
* Save and load states
* Bring up the menu in-game

All from a phone, tablet, or other device on the same network as your RetroArch machine.

No more keyboard next to your TV for when you inevitably mess up your controller configs.  No more dedicated RetroArch buttons on your controllers.

## Status

* Reads config from TOML file
* Serves static content
* HTML+CSS+JS in place (could use cleanup)
* Commands handler is a stub
* No send-command-to-UDP functionality
* No launch-retroarch functionality

NOTE: Pulls in Google's hosted JQuery.  If you don't like that, download it yourself, stick it in the /static directory, and change the relevant `<script>` tag to point to it.
