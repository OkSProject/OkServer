#!/bin/bash
rm -r /etc/systemd/system/okserver.service
deluser okserver
rm -r /home/okserver
rm -r /usr/bin/okserver
rm -r /bin/okserver
