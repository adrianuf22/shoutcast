#!/bin/bash

./sc_serv sc_serv.conf &
cd sc_trans && ./sc_trans ./sc_trans_basic.conf
