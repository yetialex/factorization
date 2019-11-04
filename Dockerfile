FROM alpine

ENV service_name=prime_factorization
ENV service_group=web
ENV listen_port=8090

COPY dist/* /opt/${service_group}/${service_name}/

WORKDIR /opt/${service_group}/${service_name}

RUN chmod +x ${service_name}

CMD ./${service_name}