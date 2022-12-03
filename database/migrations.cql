CREATE KEYSPACE "shoveler" WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1};
USE shoveler;
CREATE TABLE accounts(id text PRIMARY KEY, email text, password_hash text);
CREATE INDEX email ON accounts (email); // Cannot select by email if there isn't an index.
CREATE TABLE beacons(id text PRIMARY KEY, owner_id text, helper_id text, address text);
CREATE INDEX owner_id ON beacons (owner_id); // Cannot select by owner_id if there isn't an index.
