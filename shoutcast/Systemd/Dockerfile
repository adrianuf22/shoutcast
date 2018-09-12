FROM local/c7-systemd

WORKDIR /usr/local/shoutcast

RUN yum -y check-update || { rc=$?; [ "$rc" -eq 100 ] && exit 0; exit "$rc"; }

RUN yum install -y wget && \
    yum install -y tcpdump && \
    yum install -y vim && \
    yum install -y rsyslog

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

COPY ./rsyslog-sender.conf /etc/rsyslog.d

COPY ./tcpdump.sh ./
RUN chmod a+x ./tcpdump.sh

COPY ./tcpdump.service /lib/systemd/system
RUN systemctl enable tcpdump.service

COPY ./entrypoint.sh ./
RUN chmod a+x ./entrypoint.sh

CMD [ "/usr/sbin/init" ]