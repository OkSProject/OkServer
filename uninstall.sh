#!/bin/bash
rm -r /etc/systemd/system/oksi-http.service
deluser oksihttp
rm -r /home/oksihttp
rm -r /usr/bin/oksi-http
rm -r /bin/oksi-http
