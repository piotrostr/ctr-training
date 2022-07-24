# cloud run

The deployment itself is extremely simple, as the configuration wizard
handles all of the steps for the user. All that is required is 

```bash
gcloud run deploy --source .
```

Specifying an image from docker.io repository also works but since the image
is pushed to the artifacts repository under Google Cloud.

Latency is incredible, only 10-20ms ping is really nice!
