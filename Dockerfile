FROM scratch
ARG commit
ENV pro=test.${commit}
COPY $pro /
CMD ["/$pro","-logtostderr"]
