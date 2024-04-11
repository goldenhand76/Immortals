# Arm Cloud

## Set DNS
Open the resolv.conf file with an editor, such as nano, to make the
necessary changes. If the file doesn't already exist, this command creates it:

`sudo nano /etc/resolv.conf`

Add lines for the name servers that you want to use.

```config
nameserver 178.22.122.100
nameserver 185.51.200.2
```

## Permit Jenkins to use Docker.sock
```bash
sudo chmod 666 /var/run/docker.sock
```