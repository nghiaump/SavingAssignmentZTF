image_id="docker-core-user:1.0"
container_id=$(docker container ls -a >/dev/null | awk -v image_id="$image_id" '$2 == image_id {print $1}')
echo "CONTAINER_ID: $container_id"
echo "IMAGE: $image_id"
docker container remove $container_id
docker rmi $image_id

image_id="docker-core-saving:1.0"
container_id=$(docker container ls -a >/dev/null | awk -v image_id="$image_id" '$2 == image_id {print $1}')
echo "CONTAINER_ID: $container_id"
echo "IMAGE: $image_id"
docker container remove $container_id
docker rmi $image_id

image_id="docker-mid-saving:1.0"
container_id=$(docker container ls -a >/dev/null | awk -v image_id="$image_id" '$2 == image_id {print $1}')
echo "CONTAINER_ID: $container_id"
echo "IMAGE: $image_id"
docker container remove $container_id
docker rmi $image_id

image_id="api-gateway:1.0"
container_id=$(docker container ls -a >/dev/null | awk -v image_id="$image_id" '$2 == image_id {print $1}')
echo "CONTAINER_ID: $container_id"
echo "IMAGE: $image_id"
docker container remove $container_id
docker rmi $image_id

