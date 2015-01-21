# CassMark

Simple code to check if and how a Cassandra cluster is working and performing.
The objective is given a simple Write and Read query, go over all servers and perform 
the querys with all the Consistency levels available and get the results and the latency
of each query.

Usage
=====

Disclaimer
==========
This is not intended to replace/improve/etc cassandra-stress tool (http://goo.gl/JVpyU4)
