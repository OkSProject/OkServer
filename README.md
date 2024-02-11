# OkServer!
_OkServer!_, or _OkS!_ for short, is a Go-based HTTP server aiming to be lightweight, clean, and straightforward.
<img src="OkS.svg" width="300px" />

## Development
    git clone https://github.com/OkSProject/OkServer.git
    cd OkServer
    cp deployment/example.env .env
    make run

## Installation
<br>

    make && sudo make install
    oksi-http &
    
As simple as it is to install, you can stop it by running...

    killall oksi-http


### `systemD`
**(THIS FEATURE IS CURRENTLY BEING WORKED ON RIGHT NOW. IT DOESN'T WORK.)**<br>
**It seems like the problem lies in how the symbolic link in /etc/systemd/ is trying to read the log writing feature.**
<br>
Before creating the systemD service, create a user named `okserver` that will run *OkS!*.
    
    useradd -m oksihttp -s /bin/bash
    sudo -i -u oksihttp
    mkdir http && cd http
    wget "https://github.com/OkSProject/OkServer/releases/download/(LATEST RELEASE TAG)/(LATEST RELEASE TAR FILE).tar.gz"
    tar -xf (LATEST RELEASE TAR FILE).tar.gz
    rm -r (LATEST RELEASE TAR FILE).tar.gz

The placeholders need to be replaced with whatever is on the releases section of the repo.

If you wish to use `systemD`, see included service file. After making your desired changes, copy the service file to the main directiory, then switch to your regular user and create a symbolic link to systemD.

    cp deployment/oksi-http.service ./oksi-http.service
    ln -s /home/oksihttp/http/oksi-http.service /etc/systemd/system/oksi-http.service
    sudo systemctl daemon-reload
    sudo systemctl enable oksi-http.service
    sudo systemctl start oksi-http.service

## Configuration
Configuration is done via environmental variables.

| Name | Default Value |
| --- | --- |
| `OKSERVER_LOG_DIR` | `./log/` |
| `OKSERVER_ASSET_DIR` | `./www/html/` |
| `OKSERVER_PORT` | `8080` |

## To be worked on...
- Serve PHP, Python, and Go files.
- Ability to manage use/manage database(s).
- Ability to create .conf files for individual sites.

## Credits
_OkServer!_ and the _OkS!_ Project were started by Cole Rathbun (@cbrbygones), and is licensed under GPLv3 for anyone to use, modify, and redistribute.
