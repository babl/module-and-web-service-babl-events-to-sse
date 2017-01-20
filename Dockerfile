FROM alpine:3.4
EXPOSE 7001
RUN wget -O- "http://s3.amazonaws.com/babl/babl-server_linux_amd64.gz" | gunzip > /bin/babl-server && chmod +x /bin/babl-server
RUN mkfifo /tmp/events
ADD module-and-web-service-babl-events-to-sse_linux_amd64 /bin/sse-server
ADD start /bin/start
CMD ["start"]
