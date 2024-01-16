package scripts

const CreateTableSQL = `
	CREATE TABLE IF NOT EXISTS your_table_name (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		entity_id TEXT NOT NULL,
		status TEXT NOT NULL
);`

const InsertSQL = "INSERT INTO your_table_name (entity_id, status) VALUES (?, ?)"

const SelectSQL = "SELECT * FROM your_table_name WHERE entity_id = ?"

const UpdateSQL = "UPDATE your_table_name SET status = ? WHERE entity_id = ?"

const DeleteSQL = "DELETE FROM your_table_name WHERE entity_id = ?"
