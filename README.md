# Autobucket Operator
Kubernetes Operator that automatically creates and manages Cloud Buckets (Object Storage) for k8s Deployments. Built with Go and Operator SDK.

[![Build Status](https://travis-ci.org/bmutziu/autobucket-operator.svg?branch=master)](https://travis-ci.org/bmutziu/autobucket-operator)

This repository is a companion repo for the article: [Autobucket Operator: Building a Cloud Storage Kubernetes Operator with Go and Operator SDK](https://bmutziu.medium.com/autobucket-operator-7e77eca95fcc)

**THIS SOFTWARE IS WORK IN PROGRESS / ALPHA RELEASE AND IS NOT MEANT FOR USAGE IN PRODUCTION SYSTEMS**

## Tests
To run tests:
````
$ make test
````

## Run locally
````
# install the k8s resources
$ make install
# run the operator locally
$ make run
````

## Deploy (GCP example)
Authenticate to GCP
```
gcloud auth login
```
Create a new gcp project (choose a unique GCP_PROJECT)
```
gcloud projects create $GCP_PROJECT
```
Create a service account for the operator
```
SERVICE_ACCOUNT=autobucket-operator
gcloud iam service-accounts create $SERVICE_ACCOUNT \
--project $GCP_PROJECT
```
Grant Storage Admin role to the service account
```
gcloud projects add-iam-policy-binding $GCP_PROJECT \
--member=serviceAccount:$SERVICE_ACCOUNT@$GCP_PROJECT.iam.gserviceaccount.com \
--role=roles/storage.admin \
--project $GCP_PROJECT
```
Create Service Account keys
```
gcloud iam service-accounts keys create sa-operator.json \
--iam-account $SERVICE_ACCOUNT@$GCP_PROJECT.iam.gserviceaccount.com \
--project $GCP_PROJECT
```

*Make sure you KUBECONFIG is set before continuing, the deployment will use your current context*

Create a Kubernetes secret for the service account credentials
````
kubectl create secret generic autobucket-gcp-credentials --from-file=sa-operator.json=sa-operator.json -n autobucket-operator-system
````

Deploy resources and controller manager
````
# install the k8s resources
$ make install
# deploy the controller manager
$ GCP_PROJECT=$GCP_PROJECT make deploy
````

## Usage
Deployment annotations sample:
````
apiVersion: apps/v1
kind: Deployment
metadata:
  name: sample-deployment
  annotations:
    bm.bmutziu.me/cloud: gcp
    bm.bmutziu.me/name-prefix: bm
    bm.bmutziu.me/on-delete-policy: destroy
````

- ````bm.bmutziu.me/cloud````: cloud where the storage bucket is created. Valid options: "gcp". If this annotation is missing or empty, no bucket is created for the deployment.
- ````bm.bmutziu.me/name-prefix````: storage bucket name prefix. Default: "bm" (short name for bucketmigration).
- ````bm.bmutziu.me/on-delete-policy````: bucket deletion policy when the deployment is deleted. Valid options: "ignore" (do nothing), "destroy" (delete the storage bucket).

The full name format for the created storage buckets is "{prefix}-{namespace}-{deployment-name}"

For example, the previous deployment, when deployed to the default namespace will automatically create a GCP Bucket: "bm-default-sample-deployment"


## TODO

- [ ] Add AWS S3 Support
- [ ] Additional Bucket configuration options
- [ ] Helm chart for simpler deployment
