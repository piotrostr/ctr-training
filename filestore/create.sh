#!/bin/bash

# enable the filestore api
gcloud services enable file.googleapis.com

# create the filestore instance
gcloud filestore instances create kube-fs \
  --zone=us-central1-c \
  --tier=BASIC_HDD \
  --file-share=name="vol1",capacity=1TB \
  --network=name="default"
