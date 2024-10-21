import asyncio
import nats
from nats.errors import TimeoutError

import json
import requests
import time
import psycopg2
import erdis_kvo_list

#hostname = os.environ['HOSTNAME']

nats_servers = ["nats://ioiot:ioiot@10.5.96.10:4222","nats://ioiot:ioiot@10.4.96.10:4222","nats://ioiot:ioiot@10.1.96.10:4222"]

async def main():
############################################################################################
    async def route_command(msg):
     print("route_command!!")
     print(msg.data)

     data = msg.data.decode().split(",")
     print(data)

     respo = "not_set"

     if data[5] == "list":
         if data[6] == "append": 
             respo = await erdis_kvo_list.l_append(data)
             await msg.respond(respo.encode("utf8"))
     #await msg.respond(respo.encode("utf8"))
     #await msg.ack()
###########################################################################################
    try:
        nc = await nats.connect(nats_servers, name="erdis_kvo")
        js = nc.jetstream()
    except Exception as e:
        print(f"failed connecting to nats - {str(e)}")


    #sub = await nc.subscribe("messages.erdis_kvo", cb=route_command)
    sub = await nc.subscribe("messages.erdis_kvo")
    async for msg in sub.messages:
        await route_command(msg)
    #while True:
    #  time.sleep(1)

    """
    pull_sub1 = await js.pull_subscribe("messages.erdis_kvo", stream="test123", durable="messages")
    for i in range(0, 10):
        try:
         msgs1 = await pull_sub1.fetch(1)

         for msg in msgs1:
            print(msg)    
            #print(msg.subject[6:])
            try:
                 #data = json.loads(msg.data.decode('utf-8').strip()) # convert msg string to dictionary
                 data = msg.data.decode('utf-8').strip()

            except Exception as e:
                 print(f" message receive but failed payload formatting - {str(e)}") 
                 # log
                 await msg.ack()
                 continue


            #print(data)
            #resp = route_command(data)
            resp = "is this thing on!?!?!"
            await msg.respond(resp.encode("utf8"))
            #await msg.reply(resp)
            #await msg.ack()


        except TimeoutError:
         print(".")
   """




###
if __name__ == '__main__':
        asyncio.run(main())
