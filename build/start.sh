#!/bin/bash
#
# Generic building script for the docker packaging of the service.
#
# Maintainer: Diego @ Cyclops Labs GmbH
#
################################################################################

###
# Config variables
###################

# Base configuration.
# In most cases you want/need to touch it
PACKAGE="$(basename $(git rev-parse --show-toplevel))"
TAG="$(git describe --always --long --dirty)"
BRANCH="devel"
MAIN_FOLDER="server"

# Config file from your setup
# The original parameter probably won't be the same in your setup.
# In most cases, you will need to edit this!
DEPLOY_KEY="${HOME}/.ssh/deploy.key"

# Defined name for the files
# You should not edit this.
GIT_FILE="gitconfig"
DEPLOY_FILE="KEY"

###
# Banner
#########

echo 'Welcome to'
echo ''
echo ' $$$$$$\                      $$\ '
echo '$$  __$$\                     $$ |'
echo '$$ /  \__|$$\   $$\  $$$$$$$\ $$ | $$$$$$\   $$$$$$\   $$$$$$$\ '
echo '$$ |      $$ |  $$ |$$  _____|$$ |$$  __$$\ $$  __$$\ $$  _____|'
echo '$$ |      $$ |  $$ |$$ /      $$ |$$ /  $$ |$$ /  $$ |\$$$$$$\  '
echo '$$ |  $$\ $$ |  $$ |$$ |      $$ |$$ |  $$ |$$ |  $$ | \____$$\ '
echo '\$$$$$$  |\$$$$$$$ |\$$$$$$$\ $$ |\$$$$$$  |$$$$$$$  |$$$$$$$  |'
echo ' \______/  \____$$ | \_______|\__| \______/ $$  ____/ \_______/ '
echo '          $$\   $$ |                        $$ |'
echo '          \$$$$$$  |                        $$ |'
echo '           \______/                         \__|'
echo ''
echo '$$\                $$\                 '
echo '$$ |               $$ |                '
echo '$$ |      $$$$$$\  $$$$$$$\   $$$$$$$\ '
echo '$$ |      \____$$\ $$  __$$\ $$  _____|'
echo '$$ |      $$$$$$$ |$$ |  $$ |\$$$$$$\  '
echo '$$ |     $$  __$$ |$$ |  $$ | \____$$\ '
echo '$$$$$$$$\\$$$$$$$ |$$$$$$$  |$$$$$$$  |'
echo '\________|\_______|\_______/ \_______/ '
echo ''
echo 'Docker building system!'
echo ''

##
# Workflow
###########

echo "Copying config deploy key file..."
cp $DEPLOY_KEY $DEPLOY_FILE
chmod 600 $DEPLOY_FILE

echo "Updating the local copy of the repo..."
mkdir -p $PACKAGE && cp -Rf ../{server,models,restapi,client,go.mod} "$_"

echo "Setting vars for Docker build..."

cp Dockerfile${1}.tmp Dockerfile

sed -i "s/PACKAGE/${PACKAGE}/g" Dockerfile
sed -i "s/GIT_FILE/${GIT_FILE}/" Dockerfile
sed -i "s/SSH_FILE/${DEPLOY_FILE}/" Dockerfile
sed -i "s/BRANCH/${BRANCH}/" Dockerfile
sed -i "s/MAIN_FOLDER/${MAIN_FOLDER}/" Dockerfile
sed -i "s/VERSION/${PACKAGE}-${TAG}/" Dockerfile

echo "If this is your first run, please exit this script and edit the file accordingly to your setup, especifically the deploy-key"
echo "Press ENTER to continue if you are ready..."
read -s

echo "Building up docker image..."
echo ""

docker build --no-cache -t $PACKAGE:$TAG .

echo "Building process completed!"
rm -rf ${PACKAGE}
rm Dockerfile
