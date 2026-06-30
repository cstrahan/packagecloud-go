#!/bin/bash

host=`hostname -f`
token=`curl -XPOST --data "name=${host}" https://REDACTED_TOKEN:@packagecloud.io/install/repositories/USER/REPOSITORY/tokens.text`

helm repo add USER_REPOSITORY https://packagecloud.io/USER/REPOSITORY/helm --username $token --password does-not-matter

helm repo update

