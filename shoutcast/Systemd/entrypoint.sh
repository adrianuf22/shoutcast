#!/bin/bash

systemctl start tcpdump.service

systemctl restart rsyslog.service

./sc_serv sc_serv.conf &
cd sc_trans && ./sc_trans ./sc_trans_basic.conf
