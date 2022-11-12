# Arqui-Sw-2
## Para Poner las cosas en el Solr
docker run -d -p 8983:8983 solr solr-precreate Properties
docker exec -it [container-id] bash

curl -o /tmp/properties.csv https://raw.githubusercontent.com/SantiagoRivSal/Arqui-Sw-2/main/properties.csv?token=GHSAT0AAAAAAB2B76XQWSLTMSIDL4YBACGUY3OX6JQ

bin/post -c Properties /tmp/properties.csv



"http://localhost:8983/solr/Properties/select?defType=lucene&fq=city%3A%22"
+query1+
"%22&fq=country%3A%22"
+query2+
"%22&fq=service%3A%22"
+query3+
"%22&indent=true&q.op=OR&q=*%3A*"

