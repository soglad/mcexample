FROM busybox
COPY  ./multicast /bin/
ENTRYPOINT [ "multicast" ]