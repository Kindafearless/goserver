FROM ubuntu
ADD . goserver
ENTRYPOINT chmod +x ./goserver
