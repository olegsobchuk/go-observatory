[OBSERVATORY LANDING PAGE](../README.md)

# Problem Requirements Document

### What do we have?

Our Organisation. Voice and text messages, data transfers via internet and local network.

Currently we use only network IP system for communication (phones, data transfering etc). So that, we have to support users and maintain this complicated network and telecomunication system. So that it is mandatory point to have as much as it is possible information about network throughput and availability of different end-points (PC, Routers).

Now, due to lack of appropriate programs we use command line interface (terminal on OS Windows 10) and `ping xxx.xxx.xxx.xxx -t` command. It show us availability (in case there is a response) and latency.

### What do we need?

It is too difficult to set up new window (terminal) with needed parameters, especially when we want to observe dozen endpoints (IPs). To simplify this process we want to run all-in-one application, window, process, whatever.

Would be good to have some target endpoints, addresses, IPs in some storage, like file, DataBase etc. This would allow all users to be envolved into this process.

### Requirements:

1. Works on local PC.
2. Can follow/observe multiple devices by IP address or domain name
3. Collects information from devices.
4. Groups information.
5. Shows information in chart for time period.
6. Shows device's online/current status (available or not)
