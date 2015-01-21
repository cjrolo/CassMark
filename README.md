# CassMark

Simple code to check if and how a Cassandra cluster is working and performing.
The objective is given a simple Write and Read query, go over all servers and perform 
the querys with all the Consistency levels available and get the results and the latency
of each query.

Usage
=====
Usage of CassMark:
  * -h=[]: Hosts to Connect
  * -k="sandbox": Keyspace to use
  * -out="results.out": Output
  * -p="": Password for the Cluster
  * -u="": User to connect to the Cluster
Example:

    CassMark -k="mykeyspace" -u="User" -p="Password" -out="log.out" -h="192.168.1.0" -h="192.168.1.1"

Disclaimer
==========
This is not intended to replace/improve/etc cassandra-stress tool (http://goo.gl/JVpyU4)
