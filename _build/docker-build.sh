# Build a docker image
docker build -t fishgoddess/postar:v0.1.2-alpha .

# Login to docker
docker login

# Push to docker hub
docker push fishgoddess/postar:v0.1.2-alpha

# Logout from docker
docker logout