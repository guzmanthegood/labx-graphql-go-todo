package model

type User struct {
	ID 		string
	Name 	string
}

// CreateUser SQL mutation
func (s store) CreateUser(name string) (*User, error) {
	var res []User
	sql := "INSERT INTO public.user(name) VALUES($1) RETURNING *"
	err := s.db.Select(&res, sql, name)
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, nil
	}
	return &res[0], nil
}

// UpdateUser SQL mutation
func (s store) UpdateUser(id int32, name string) (*User, error) {
	var res []User
	sql := "UPDATE public.user SET name=$1 WHERE id=$2 RETURNING *"
	err := s.db.Select(&res, sql, name, id)
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, nil
	}
	return &res[0], nil
}

// DeleteUser SQL mutation
func (s store) DeleteUser(id int32) error {
	sql := "DELETE FROM public.user WHERE id=$1"
	r, err := s.db.Queryx(sql, id)
	defer r.Close()
	return err
}
