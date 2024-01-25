[OBSERVATORY LANDING PAGE](../README.md)

# Problem Requirements Document

### What do we have?

Our Organization. Voice and text messages, data transfers via the internet, and the local network.

Currently, we use only the network IP system for communication (phones, data transferring, etc.). Consequently, we have to support users and maintain this complicated network and telecommunication system. Therefore, it is mandatory to have as much information as possible about network throughput and the availability of different endpoints (PCs, routers).

Now, due to a lack of appropriate programs, we use the command line interface (terminal on OS Windows 10) and the `ping xxx.xxx.xxx.xxx -t` command. It shows us availability (in case there is a response) and latency.

### What do we need?

It is too difficult to set up a new window (terminal) with needed parameters, especially when we want to observe dozens of endpoints (IPs). To simplify this process, we want to run an all-in-one application, window, process, or whatever.

It would be good to have some target endpoints, addresses, IPs in some storage, like a file, database, etc. This would allow all users to be involved in this process.

### Requirements:

1. Works on a local PC.
2. Can follow/observe multiple devices by IP address or domain name
3. Collects information from devices.
4. Groups information.
5. Shows information in a chart for a specific time period.
6. Shows the device's online/current status (available or not)
