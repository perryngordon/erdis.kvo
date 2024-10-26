nats -s "nats://<username>:<passowrd>@<nats_server_address>" request erdis.kvo.list.append.<bucket_name>.<key> <value_to_append>
<br>
nats -s "nats://<username>:<passowrd>@<nats_server_address>" request erdis.kvo.list.remove.<bucket_name>.<key> <value_to_remove>
