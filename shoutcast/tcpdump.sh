#!/bin/bash
tcpdump -nnttlS "tcp[13] & 8!=0" | cat >> /var/log/messages/access.log