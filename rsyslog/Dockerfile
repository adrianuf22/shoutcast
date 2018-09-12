FROM ubuntu:latest

RUN apt-get update && \
    apt-get install -y rsyslog

RUN rm /etc/rsyslog.d/*
COPY ./rsyslog-receiver.conf /etc/rsyslog.d

RUN service rsyslog restart

EXPOSE 60000