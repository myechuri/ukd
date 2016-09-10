#!/bin/bash
# Provision Host Env
#
# Sets up host env for ukd.
set -ex

mkdir -p /var/lib/ukd/images
cp nativeexample.img /var/lib/ukd/images
cp server /usr/bin
cp ukd /usr/bin
cp ukdctl /usr/bin
rm ukdctl server ukd ukd-v0.1-linux-amd64.tar.gz nativeexample.img
