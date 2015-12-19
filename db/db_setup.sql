	--create dblink extension for creating database and role in the same transaction
CREATE extension IF not EXISTS dblink;

--function to create a specified admin role
CREATE OR REPLACE FUNCTION create_admin_role(db_con_name text, role_name text, pass text) RETURNS void AS $$
BEGIN 
	IF EXISTS(SELECT 1 FROM pg_roles WHERE rolname=role_name) THEN
		PERFORM dblink_exec(db_con_name,format('drop role %s', role_name));
	END IF;
	
	PERFORM dblink_exec(db_con_name,format('CREATE role %I with superuser login password ''%s''',role_name, pass));
END;
$$
LANGUAGE plpgsql;


--funtion to create a specified db
CREATE OR REPLACE FUNCTION create_db(db_con_name text, dbname text, db_owner text) RETURNS void AS $$
BEGIN
	IF EXISTS(SELECT 1 FROM pg_database WHERE datname=dbname) THEN
		PERFORM dblink_exec(db_con_name,format('drop dtatbase %s',dbname));
	END IF;	
	PERFORM dblink_exec(db_con_name,format('CREATE database %s with owner=%s',dbname,db_owner));
	
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

----------------------------------------------- create admin role and db -----------------------------------

DO
$$
DECLARE
	con text := 'root_con';
	adminname text := 'koteladmin';
	pass text := 'kotelpass';
	dbname text := 'koteldb';
BEGIN
	PERFORM create_root_connection(con);
	PERFORM create_admin_role(con,adminname,pass);
	PERFORM create_db(con,dbname,adminname);
	PERFORM destroy_root_connection(con);
END;
$$
LANGUAGE plpgsql;

