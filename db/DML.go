package db

const (
	CreatingTask = `INSERT INTO tasks (title, description, user_id)
						VALUES ($1, $2, $3)`
	GetAllTasks = `SELECT t.id,
							   t.title,
							   t.description,
							   u.full_name AS user_full_name
						FROM tasks t
						JOIN users u ON t.user_id = u.id
						WHERE t.is_deleted = FALSE;`
	GetTaskByID = `SELECT t.id,
					   t.title,
					   t.description,
					   u.full_name AS user_full_name
					FROM tasks t
					JOIN users u ON t.user_id = u.id
					WHERE t.is_deleted = FALSE
					  AND t.id = $1;`
	UpdateTaskByID = `UPDATE tasks
							SET title = $1,
								description = $2,
								is_done = $3
							WHERE id = $4;`
	SoftDeleteTaskByID = `UPDATE tasks SET is_deleted = TRUE WHERE id = $1;`
)
