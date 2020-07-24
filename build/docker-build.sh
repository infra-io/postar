# Build a docker image
docker build -t fishgoddess/postar:v0.1.0-alpha .

# Login to docker
docker login

# Push to docker hub
docker push