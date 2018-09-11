FROM ubuntu:latest

WORKDIR /usr/local/shoutcast

RUN apt-get update && \
    apt-get install -y wget && \
    apt-get install -y tcpdump && \
    apt-get install -y vim && \
    apt-get install -y supervisor

RUN wget http://download.nullsoft.com/shoutcast/tools/sc_serv2_linux_x64_09_09_2014.tar.gz && \
    tar xfz sc* && \
    mkdir ../sc_server && \
    cp sc_serv ../sc_server && \
    cd ../sc_server  && \
    mkdir control  && \
    mkdir logs  && \
    chmod +x sc_serv

RUN mkdir sc_trans && \
    wget http://mirror.lchost.net/download.nullsoft.com/shoutcast/tools/sc_trans_linux_x64_10_07_2011.tar.gz && \
    tar -xzf sc_trans_linux_x64_10_07_2011.tar.gz -C sc_trans && \
    cd sc_trans && \
    chmod a+x sc_trans

EXPOSE 8000

COPY ./supervisord.conf ./

COPY ./entrypoint.sh ./
RUN chmod a+x ./entrypoint.sh

ENTRYPOINT "./entrypoint.sh"