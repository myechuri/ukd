#!/bin/bash
# Provision Host Env
#
# Sets up host env for ukd.
set -ex

mkdir -p /var/lib/ukd/images
cp hello-world-loop.img /var/lib/ukd/images
cp ukd /usr/bin
cp ukdctl /usr/bin
rm ukdctl ukd ../ukd-v0.1dev-linux-aarch64.tar.gz hello-world-loop.img
