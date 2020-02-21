package model

type Todo struct {
	ID 			int32	`db:"id"`
	Text 		string	`db:"description"`
	Status  	string	`db:"status"`
	UserID  	int32	`db:"user_id"`
}

// GetTodo SQL query
func (s store) GetTodo(id int32) (*Todo, error) {
	var res []Todo
	sql := "SELECT * FROM public.todo WHERE id=$1"
	err := s.db.Select(&res, sql, id)
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, nil
	}
	return &res[0], nil
}

func (s store) CreateTodo(text string, userID int32) (*Todo, error) {
	var res []Todo
	sql := "INSERT INTO public.todo(description, user_id) VALUES($1, $2) RETURNING *"
	err := s.db.Select(&res, sql, text, userID)
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, nil
	}
	return &res[0], nil
}

func (s store) UpdateTodo(id int32, text, status string) (*Todo, error){
	var res []Todo
	sql := "UPDATE public.todo SET description=$1, status=$2 WHERE id=$3 RETURNING *"
	err := s.db.Select(&res, sql, text, status, id)
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, nil
	}
	return &res[0], nil
}

// DeleteTodo SQL mutation
func (s store) DeleteTodo(id int32) error {
	sql := "DELETE FROM public.todo WHERE id=$1"
	r, err := s.db.Queryx(sql, id)
	defer r.Close()
	return err
}