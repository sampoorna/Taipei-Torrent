This is a modified version of Taipei Torrent - meant to test plausibility of streaming strategies using the BitTorrent protocol. 

Our measurements will be done through measuring the number of 'written' blocks per second, and averaging that to determine the effective download rate. The reason we don't measure total download bandwidth is because partial blocks can be thrown away, and extra blocks may be send and discarded. We want an accurate measure of the file's download speed.

Different measurements we will be performing the torrent download:
- without any strategy, using randomized block gets
- with getting the recent piece first
- with getting the recent piece first, and if sufficient amount is buffered letting it get random pieces of the torrent until the buffer has gone down.
- with getting recent pieces first using endgame mode, once buffered, allowing normal download

We'd like to graph these both with unrestricted bandwidth, and with restricted bandwidths - to test effectiveness in countries with generally less effective internet.

Profiling:
- time elapsed
- speed
- useful blocks out of total blocks

Taipei Torrent
==============

This is a simple command-line-interface BitTorrent client coded in the go
programming language.

Features:
---------

+ Supports multiple torrent files
+ Magnet links
+ DHT
+ IPv6
+ UDP trackers
+ UPnP / NAT-PMP automatic firewall configuration
+ Socks5 proxy support

Additional Features:
--------------------

+ It can act as a tracker if you start it with the -createTracker flag
=> Usage: go run main.go -createTracker=127.0.0.1:8080 nameOfTorrent.torrent

FAQ:
----

Q: Why is it named Taipei Torrent?

A: I started writing it while visiting beautiful Taipei, Taiwan

Q: What is the license?

A: See the LICENSE file.

Current Status
--------------

+ Tested on Go 1.3
+ Tested on Windows, Linux and Mac OS X.
+ People tell me they've run it on Android, too.

Development Roadmap
-------------------

+ Full UPnP support (need to be able to search for an unused listener port,
  detect we have already acquired the port, defend the report against router
  reboots, release the listener port when we quit.)
+ Clean up source code
+ Deal with TODOs
+ Perhaps a web-based status UI.

Download, Install, and Build Instructions
-----------------------------------------

1. Download and install the Go tools from http://golang.org

2. Use the "go" command to download, install, and build the Taipei-Torrent
app:

    go get github.com/jackpal/Taipei-Torrent

Usage Instructions
------------------

    Taipei-Torrent mydownload.torrent
    Taipei-Torrent --useDHT "magnet:?xt=urn:btih:bbb6db69965af769f664b6636e7914f8735141b3"

or

    Taipei-Torrent -help

Third-party Packages
--------------------

https://github.com/jackpal/bencode-go - Bencode encoder/decoder

http://code.google.com/p/go-nat-pmp - NAT-PMP firewall client

https://github.com/hailiang/gosocks - SOCKS5 proxy support

https://github.com/nictuku/dht      - Distributed Hash Table

https://github.com/nictuku/nettools - Network utilities

Google+ Community
-----------------

https://plus.google.com/u/0/communities/100997865549971977580

