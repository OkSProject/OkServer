# OkServer!
_OkServer!_, or _OkS!_ for short, is a Go-based HTTP web server aimed to be lightweight, clean, and straightforward.
<img src="OkS.svg" width="300px" />

## Development
    git clone https://github.com/OkSProject/OkServer.git
    cd OkServer
    cp deployment/example.env .env
    make run

## Installation
**(THIS FEATURE IS CURRENTLY BEING WORKED ON RIGHT NOW. IT'S NOT RECOMMENDED YOU RUN THIS UNTIL IT'S BEEN TESTED AND FIXED. INSTRUCTIONS WILL BE UPDATED IF NEEDED.)**
<br>

    make && sudo make install

### `systemD`
**(THIS FEATURE IS CURRENTLY BEING WORKED ON RIGHT NOW. IT DOESN'T WORK.)**
<br>
Before creating the systemD service, create a user named `okserver` that will run *OkS!*.
    
    useradd -m okserver -s /bin/bash
    sudo -i -u okserver
    mkdir http
    wget "https://github.com/OkSProject/OkServer/releases/download/(LATEST RELEASE TAG)/(LATEST RELEASE TAR FILE).tar.gz"
    tar -xf (LATEST RELEASE TAR FILE).tar.gz
    mv ./(LATEST RELEASE TAR FILE) ./http

The placeholders needs to be replaced with whatever is on the releases section of the repo.

If you wish to use `systemD`, see included service file. After making your desired changes, copy the service file to the main directiory and create a symbolic link to systemD.

    cp deployment/okserver.service ./okserver.service
    ln -s /home/okserver/http/okserver.service /etc/systemd/system/okserver.service
    sudo systemctl daemon-reload
    sudo systemctl enable okserver.service
    sudo systemctl start okserver.service

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
