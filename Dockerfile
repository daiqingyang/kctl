FROM scratch
ARG commit
ENV pro=test.${commit}
COPY $pro /test
CMD ["/test","-logtostderr"]
