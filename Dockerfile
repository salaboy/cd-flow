FROM scratch
EXPOSE 8080
ENTRYPOINT ["/cd-flow"]
COPY ./build/linux /