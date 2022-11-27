#!/bin/bash

# Add local user
# Either use the LOCAL_USER_ID if passed in at runtime or
# fallback

if [ -z "$LOCAL_USER_ID" ]; then
    USER_ID=-9001
else
    USER_ID=$LOCAL_USER_ID
fi
echo "TEST: $USER_ID"


echo "Starting with UID : $USER_ID"
mkdir -p /home/user
adduser -s /bin/bash -u $USER_ID -H -D user
chown -R $USER_ID: /home/user
export HOME=/home/user
su-exec user "$@"