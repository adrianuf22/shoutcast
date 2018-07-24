FROM ubuntu:latest

WORKDIR /usr/local/shoutcast

RUN apt-get update && \
    apt-get install -y wget && \
    apt-get install -y tcpdump

RUN wget http://download.nullsoft.com/shoutcast/tools/sc_serv2_linux_x64_09_09_2014.tar.gz && \
    tar xfz sc* && \
    mkdir ../sc_server && \
    cp sc_serv ../sc_server && \
    cd ../sc_server  && \
    mkdir control  && \
    mkdir logs  && \
    chmod +x sc_serv

RUN mkdir sct && \
    wget http://mirror.lchost.net/download.nullsoft.com/shoutcast/tools/sc_trans_linux_x64_10_07_2011.tar.gz && \
    tar -xzf sc_trans_linux_x64_10_07_2011.tar.gz -C sct && \
    cd sct && \
    chmod a+x sc_trans

EXPOSE 8000
EXPOSE 7999

COPY sc_serv.conf ./
COPY sc_trans_basic.conf ./sct
COPY music1.mp3 ./sct/music/
COPY music2.mp3 ./sct/music/