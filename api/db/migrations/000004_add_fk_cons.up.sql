ALTER TABLE rooms
ADD COLUMN host_id BIGINT,
ADD CONSTRAINT fk_host_id FOREIGN KEY (host_id) REFERENCES hosts (id);