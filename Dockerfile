FROM scratch
COPY test.bin /test.bin
CMD ["/test.bin","-logtostderr"]
