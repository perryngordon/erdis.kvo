import asyncio
import nats
from nats.errors import TimeoutError


async def l_append(data):
    print("l_append")
    #nats_servers = ["nats://ioiot:ioiot@10.5.96.10:4222","nats://ioiot:ioiot@10.4.96.10:4222","nats://ioiot:ioiot@10.1.96.10:4222"]
    nats_servers = ["nats://ioiot:ioiot@10.1.96.10:4222"]

    #try:
    #    nc = await nats.connect(nats_servers)
    #    js = nc.jetstream()
    #except Exception as e:
    #    print(f"l_append - failed connecting to nats - {str(e)}")

    # use nats cli for kv

    #bucket = 
    #key = 
    #valut_to_append = 

    # get string list from kv
    # cast to list
    # append
    # cast to string
    # update kv
    #return message


    return "l_append..."
