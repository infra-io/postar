FROM alpine:3.12.0
MAINTAINER  fishgoddess

# Deploy postar to work directory
WORKDIR /postar/
COPY ./postar-v0.1.0-alpha ./
COPY ./src/logit.conf ./
COPY ./_examples/config/postar.ini ./
EXPOSE 5779
EXPOSE 5780

# Run postar
CMD ["./postar-v0.0.1-alpha", "-c", "./postar.ini"]