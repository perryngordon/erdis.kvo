clone repo <br>
cd to repo <br>

---

the nats url will default to localhost if NATS_URL is not set

```
export NATS_URL=nats://<username>:<passowrd>@<nats_server_address>
```

```
./erdis.kvo 
```

multiple running instances of erdis.kvo will behave as a queue group 

---

```
nats -s "nats://<username>:<passowrd>@<nats_server_address>" request erdis.kvo.list.pop.<bucket_name>.<key> pop
```

```
nats -s "nats://<username>:<passowrd>@<nats_server_address>" request erdis.kvo.list.append.<bucket_name>.<key> <value_to_append>
```

```
nats -s "nats://<username>:<passowrd>@<nats_server_address>" request erdis.kvo.list.remove.<bucket_name>.<key> <value_to_remove>
```

---

[home](https://github.com/perryngordon/erdis.kvo/tree/main)
