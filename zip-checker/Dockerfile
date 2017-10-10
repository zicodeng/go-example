FROM scratch

COPY zip-checker zip-checker
COPY client client

EXPOSE 443

# Which file should this container execute when it starts running.
ENTRYPOINT ["/zip-checker"]