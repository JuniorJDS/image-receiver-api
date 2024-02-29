FROM rabbitmq:3-management

COPY ./script/rabbitmq.conf /etc/rabbitmq
COPY ./script/definitions.json /etc/rabbitmq

RUN cat /etc/rabbitmq/rabbitmq.conf