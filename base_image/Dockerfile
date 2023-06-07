ARG BUILD_OS=ubuntu
ARG BUILD_TAG=18.04

FROM $BUILD_OS:$BUILD_TAG

RUN apt-get update && apt-get upgrade -y \
    && apt-get install --no-install-recommends -y ca-certificates curl \
    && apt-get autoremove -y && apt-get clean \
    && rm -rf /tmp/* /var/tmp/* \
    && rm -rf /var/lib/apt/lists/*

RUN mkdir -p /etc/envoy

ADD envoy /usr/local/bin/envoy
ADD envoy.yaml /etc/envoy/envoy.yaml

RUN apt-get update && apt-get install python3-pip -y \
    && apt-get autoremove -y && apt-get clean \
    && rm -rf /tmp/* /var/tmp/* \
    && rm -rf /var/lib/apt/lists/*

RUN pip3 install --no-cache-dir gunicorn httpbin

RUN mkdir -p /home/envoy/logs

WORKDIR /work
ADD tianchi/envoy-ecds.yaml /work/envoy-ecds.yaml
COPY tianchi/build /build
COPY tianchi/conf /conf
ADD entrypoint.sh /work/
RUN chmod +x /work/entrypoint.sh

ENTRYPOINT ["/work/entrypoint.sh"]
