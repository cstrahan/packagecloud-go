#!/bin/bash

host=`hostname -f`
token=`curl -XPOST --data "name=${host}" https://REDACTED_TOKEN:@packagecloud.io/install/repositories/USER/REPOSITORY/tokens.text`

gem source --add https://${token}:@packagecloud.io/USER/REPOSITORY/

