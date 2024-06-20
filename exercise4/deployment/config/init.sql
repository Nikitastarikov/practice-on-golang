CREATE EXTENSION IF NOT EXISTS dblink;

DO
$do$
	BEGIN
		IF EXISTS (SELECT FROM pg_database WHERE datname = 'calhounio_demo') THEN
			RAISE NOTICE 'Database calhounio_demo already exists';
		ELSE
			PERFORM dblink_exec('dbname=' || current_database(), 'CREATE DATABASE calhounio_demo');
		END IF;
	END;
$do$;
