# Arqui-Sw-2
## Para Poner las cosas en el Solr
docker run -d -p 8983:8983 solr solr-precreate Properties
docker exec -it [container-id] bash

curl -o /tmp/properties.csv https://raw.githubusercontent.com/SantiagoRivSal/Arqui-Sw-2/main/properties.csv?token=GHSAT0AAAAAAB2B76XQWSLTMSIDL4YBACGUY3OX6JQ

bin/post -c Properties /tmp/properties.csv
