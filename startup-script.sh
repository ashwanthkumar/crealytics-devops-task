#!/bin/bash

set -e

function metadata {
  KEY=$1
  echo `curl -sf -H "Metadata-Flavor: Google" "http://metadata.google.internal/computeMetadata/v1/instance/${KEY}"`
}

CUSTOM_USER_KEY="custom-user"
CUSTOM_USER_PASSWD_KEY="custom-user-passwd"

NAME=$(metadata "name")
ZONE=$(metadata "zone")
USERNAME=$(metadata "attributes/${CUSTOM_USER_KEY}")
PASSWORD=$(metadata "attributes/${CUSTOM_USER_PASSWD_KEY}")

# Create a user identified by ${USERNAME}:${PASSWORD} and having sudo access on the instance

addgroup ${USERNAME}-g
useradd ${USERNAME} --create-home --shell /bin/bash --group ${USERNAME}-g
usermod -aG sudo ${USERNAME}
echo "${USERNAME}:${PASSWORD}" | chpasswd

# Delete the plain text username and password once we've set things up
gcloud compute instances remove-metadata ${NAME} --zone ${ZONE} --keys ${CUSTOM_USER_KEY},${CUSTOM_USER_PASSWD_KEY}
