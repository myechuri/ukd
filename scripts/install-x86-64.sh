#!/bin/bash
# Provision Host Env
#
# Sets up host env for ukd.
set -ex

mkdir -p /var/lib/ukd/images
cp hello-world-1.img /var/lib/ukd/images
cp hello-world-2.img /var/lib/ukd/images
cp server /usr/bin
cp ukd /usr/bin
cp ukdctl /usr/bin
rm ukdctl server ukd ../ukd-v0.1dev-linux-x86-64.tar.gz hello-world-1.img hello-world-2.img
