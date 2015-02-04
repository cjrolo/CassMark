/* The Replication should match your cluster replication strategy! */
CREATE KEYSPACE IF NOT EXISTS sandbox WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };
 CREATE TABLE data (
 	id uuid,
 	data text,
 	PRIMARY KEY (id));