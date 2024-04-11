## Installation 
```bash
sudo apt-get install v2ray
```

## Run V2ray
```bash
cd /etc/v2ray
V2RAY_LOCATION_ASSET=/etc/v2ray v2ray
```

## Test Setup
When you are sure that the V2Ray is working properly, at the gateway, execute `curl -x socks5://127.0.0.1:1080 google.com ` to test whether your setup can bypass GFW. (Here socks5 refers to the inbound protocol and 1080 is the inbound port ) . If the output is something like the following, you are good. Otherwise, there's something wrong with your setup and you need to recheck what you have missed.

```html
<HTML><HEAD><meta http-equiv="content-type" content="text/html;charset=utf-8">
<TITLE>301 Moved</TITLE></HEAD><BODY>
<H1>301 Moved</H1>
The document has moved
<A HREF="http://www.google.com/">here</A>.
</BODY></HTML>
```

## Transparent Proxy
```bash
iptables -t nat -N V2RAY # Create a new chain called V2RAY
iptables -t nat -A V2RAY -d 192.168.0.0/16 -j RETURN # Direct connection 192.168.0.0/16
iptables -t nat -A V2RAY -p tcp -j RETURN -m mark --mark 0xff # Directly connect SO_MARK to 0xff traffic (0xff is a hexadecimal number, numerically equivalent to 255), the purpose of this rule is to avoid proxy loopback with local (gateway) traffic
iptables -t nat -A V2RAY -p tcp -j REDIRECT --to-ports 12345 # The rest of the traffic is forwarded to port 12345 (ie V2Ray)
iptables -t nat -A PREROUTING -p tcp -j V2RAY # Transparent proxy for other LAN devices
iptables -t nat -A OUTPUT -p tcp -j V2RAY # Transparent proxy for this machine
```
Then set the iptables rule of UDP traffic for the transparent proxy device:
```bash
ip rule add fwmark 1 table 100
ip route add local 0.0.0.0/0 dev lo table 100
iptables -t mangle -N V2RAY_MASK
iptables -t mangle -A V2RAY_MASK -d 192.168.0.0/16 -j RETURN
iptables -t mangle -A V2RAY_MASK -p udp -j TPROXY --on-port 12345 --tproxy-mark 1
iptables -t mangle -A PREROUTING -p udp -j V2RAY_MASK
```


## Sources 
https://guide.v2fly.org/en_US/app/transparent_proxy.html#procedures