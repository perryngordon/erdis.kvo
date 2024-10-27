clone repo <br>
cd to repo <br>

---

nats -s "nats://<username>:<passowrd>@<nats_server_address>" request erdis.kvo.list.append.<bucket_name>.<key> <value_to_append>

--

export NATS_URL=nats://\<username>:\<passowrd>@<nats_server_address> 
./erdis.kvo 

---

nats -s "nats://\<username>:\<passowrd>@<nats_server_address>" request erdis.kvo.list.append.<bucket_name>.\<key> <value_to_append>

<br>

nats -s "nats://\<username>:\<passowrd>@<nats_server_address>" request erdis.kvo.list.remove.<bucket_name>.\<key> <value_to_remove>

---

[home](https://github.com/perryngordon/erdis.kvo/tree/main)
