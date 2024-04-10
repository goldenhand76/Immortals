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