FROM alpine
COPY . goserver
RUN ./goserver
