google-chrome \
--proxy-server='socks5://localhost:$portTunnel' \
--proxy-bypass-list='<-loopback>' \
--host-resolver-rules='MAP * ~NOTFOUND , EXCLUDE localhost'


