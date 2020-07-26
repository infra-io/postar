### Docker Installation Manual

#### 1. Pull an image, such as version v0.1.2-alpha.
```bash
$ docker pull fishgoddess/postar:v0.1.2-alpha
```

#### 2. Run this image by two ways.
* By docker

_Before running this image, you should prepare a config file named "postar.ini", and put it to where you execute this command._

_A typical config file will be like this:_
```ini
[smtp]
host = smtp.your-smtp-host.com
port = 587
username = your-smtp-username@your-smtp-server.com
password = your-smtp-password
```

_You should replace the information of yours._

_Then, it's ok! Try to run by this command, and you will see a welcome interface:_
```bash
$ docker run --name postar -p 5779:5779 -v /root/postar/postar.ini:/opt/postar-v0.1.2-alpha/postar.ini -d fishgoddess/postar:v0.1.2-alpha
```

_Remember "/root/postar/postar.ini" is where your config file is._

_The work directory is "/opt/postar-v0.1.2-alpha/"._

_If you want to enter the container, try this:_
```bash
$ docker exec -it postar sh
```

_Notice that the based image we used is alpine, which doesn't have bash._

_So, you should use sh instead of bash._

* By docker-compose

_First, you should prepare a file named "docker-compose.yaml"._

_A typical docker-compose file will be like this:_
```yaml
version: "3.1"

services:
  postar:
    image: fishgoddess/postar:v0.1.2-alpha
    ports:
      - "5779:5779"
      # - "5780:5780"
    volumes:
      - ./logs:/opt/postar-v0.1.2-alpha/logs
      - ./logs/error:/opt/postar-v0.1.2-alpha/logs/error
      - ./postar.ini:/opt/postar-v0.1.2-alpha/postar.ini
```

_Notice that we map directory "logs" in the container to local directory "logs"._

_So, you can read these logs conveniently._

_Then, you can execute docker-compose command to run it:_
```bash
$ docker-compose up -d
```

#### 3. Enjoy it!