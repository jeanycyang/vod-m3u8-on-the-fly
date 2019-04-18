#!/bin/sh
openssl rand 16 > /var/vod/file.key
echo $BASE_URL$FILE_NAME/file.key > /var/vod/file.keyinfo
echo file.key >> /var/vod/file.keyinfo
echo $(openssl rand -hex 16) >> /var/vod/file.keyinfo