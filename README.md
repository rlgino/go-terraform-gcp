# Go GCP Cloud Function Example

## Commands for Cloud function

````shell
gcloud functions deploy HelloHTTP --runtime go119 --trigger-http --allow-unauthenticated
gcloud functions describe HelloHTTP
curl -X POST https://us-central1-pruebas-378003.cloudfunctions.net/HelloHTTP -H "Content-Type:application/json"  -d '{"name":"Gino"}'
gcloud functions describe HelloHTTP
gcloud functions delete HelloHTTP
````

## Commands for Cloud Run

````shell
gcloud init deploy
gcloud run deploy
````

## Source Google Cloud Run

* [Cloud Functions vs. Cloud Run: when to use one over the other](https://cloud.google.com/blog/products/serverless/cloud-run-vs-cloud-functions-for-serverless)
