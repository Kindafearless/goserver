FROM alpine
COPY ./goserver .
ENTRYPOINT ["./goserver"]
