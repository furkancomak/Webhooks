FROM golang:1.12

ADD kubebot /opt/
RUN wget --quiet http://storage.googleapis.com/kubernetes-release/release/v1.15.3/bin/linux/amd64/kubectl -O /usr/bin/kubectl && \
    chmod +x /usr/bin/kubectl && \
    chmod +x /opt/kubebot && \
    mkdir -p brain
ENV PATH="/opt/:$PATH"
ENTRYPOINT ["kubebot"]
CMD ["-c", "/opt/conf"]