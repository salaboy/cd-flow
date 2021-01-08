FROM scratch
COPY build/linux/cdf /
ENTRYPOINT ["/cdf"]