#!/bin/bash

# delete the filestore instance
gcloud filestore instances delete kube-fs \
  --zone=us-central1-c \
  --quiet

# disable the filestore api
gcloud services disable file.googleapis.com
