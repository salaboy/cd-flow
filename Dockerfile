FROM centos:7
ENTRYPOINT ["cdf", "--help"]
COPY build/linux/cdf /usr/bin/cdf
