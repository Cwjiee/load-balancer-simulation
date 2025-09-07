#/bin/sh

NUMBER_OF_REQUESTS=15

for i in $(seq 1 $NUMBER_OF_REQUESTS)
do
  curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET http://weijie.com:80
done
