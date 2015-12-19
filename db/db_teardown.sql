
--Create and extension for connecting to db
CREATE EXTENSION IF NOT EXISTS dblink;

--fucntion to remove the specIFied role
CREATE OR REPLACE FUNCTION remove_role(db_con_name text,role_name text) RETURNS void AS $$
BEGIN
	IF EXISTS( SELECT 1 FROM pg_roles WHERE rolname=role_name) THEN
		PERFORM dblink_exec(db_con_name, format('drop role %s',role_name));
	END IF;
END;
$$
LANGUAGE plpgsql;


--function to remove the specIFied db
CREATE OR REPLACE FUNCTION remove_db(db_con_name text,dbname text) RETURNS void AS $$
BEGIN
	IF EXISTS (SELECT 1 FROM pg_database WHERE datname=dbname) THEN
		PERFORM dblink_exec(db_con_name, format('drop database %s', dbname));
	END IF;
END;
$$
LANGUAGE plpgsql;


--function to terminate other connections that exits 
CREATE OR REPLACE FUNCTION terminate_connections(db_con_name text, dbname text) RETURNS void AS $$
BEGIN
	IF (EXISTS (SELECT 1 FROM pg_database WHERE datname=dbname) and (dbname <> 'postgres') )THEN
		PERFORM pg_terminate_backend(pg_stat_activity.pid) FROM pg_stat_activity WHERE pg_stat_activity.datname=dbname and pid <> pg_backend_pid();
	END IF;
END;
$$
LANGUAGE plpgsql;

--fucntion to create a root connection
CREATE OR REPLACE FUNCTION create_root_connection(db_con_name text) RETURNS void AS $$
BEGIN
	PERFORM dblink_connect(db_con_name, format('dbname=postgres user=postgres password=postgres'));
END;
$$
LANGUAGE plpgsql;


--function to destroy a root connection
CREATE OR REPLACE FUNCTION destroy_root_connection(db_con_name text) RETURNS void AS $$
BEGIN
	PERFORM dblink_disconnect(db_con_name);
END;
$$
LANGUAGE plpgsql;


------------------------------------------------------destroy admin role and db---------------------------------------------------------------
DO
$$
DECLARE 
	con text := 'roo_con';
BEGIN
	--PERFORM destroy_root_connection(con);
	PERFORM create_root_connection(con);
	PERFORM terminate_connections(con, 'koteldb');
	PERFORM removemove_db(con, 'koteldb');
	PERFORM remove_role(con, 'koteladmin');
	PERFORM destroy_root_connection(con);	
END;
$$

