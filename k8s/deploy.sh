REGION="asia-southeast1"
ZONE="asia-southeast1-b"
CLUSTER_NAME="sg1"
CERT=$CLUSTER_NAME
DNS="$CLUSTER_NAME.biddlr.com"
PROJECT="hazel-framing-184014"
INSTANCE_TYPE="n2d-custom-4-5120"
INSTANCE_DISK_SIZE="20"
MIN_NODES="1"
MAX_NODES="3"
DEFAULT_NODE="1"
# IP="34.36.108.165"

gcloud auth login
gcloud auth application-default login

gcloud compute addresses create $CLUSTER_NAME --project=$PROJECT --global

IP=$(gcloud compute addresses list --project=$PROJECT --global | grep $CLUSTER_NAME | awk '{print $2}')

echo "=========================="
echo "IP: $IP"

gcloud beta compute ssl-certificates create $CLUSTER_NAME --project=$PROJECT --global --domains=$DNS

gcloud beta container --project $PROJECT clusters create $CLUSTER_NAME --region $REGION --enable-network-policy --no-enable-basic-auth --cluster-version "1.27.3-gke.100" --release-channel "stable" --machine-type "$INSTANCE_TYPE" --image-type "COS_CONTAINERD" --disk-type "pd-ssd" --disk-size "$INSTANCE_DISK_SIZE" --metadata disable-legacy-endpoints=true --scopes "https://www.googleapis.com/auth/devstorage.read_only","https://www.googleapis.com/auth/logging.write","https://www.googleapis.com/auth/monitoring","https://www.googleapis.com/auth/servicecontrol","https://www.googleapis.com/auth/service.management.readonly","https://www.googleapis.com/auth/trace.append" --max-pods-per-node "50" --num-nodes "$DEFAULT_NODE" --logging=SYSTEM,WORKLOAD --monitoring=SYSTEM --enable-ip-alias --network "projects/$PROJECT/global/networks/default" --subnetwork "projects/$PROJECT/regions/$REGION/subnetworks/default" --no-enable-intra-node-visibility --default-max-pods-per-node "110" --node-labels pool=regular --enable-autoscaling --total-min-nodes "$MIN_NODES" --total-max-nodes "$MAX_NODES" --location-policy "BALANCED" --no-enable-master-authorized-networks --addons HorizontalPodAutoscaling,HttpLoadBalancing,GcePersistentDiskCsiDriver --enable-autoupgrade --enable-autorepair --max-surge-upgrade 1 --max-unavailable-upgrade 0 --enable-shielded-nodes --node-locations $ZONE

gcloud container clusters get-credentials $CLUSTER_NAME --region $REGION

kubectl apply -f secrets.yaml

kubectl apply -f ws-backend-config.yaml

kubectl apply -f deployments

sleep 5

kubectl apply -f ingress.yaml

# gcloud container clusters delete sg1 --zone asia-southeast1 --quiet
