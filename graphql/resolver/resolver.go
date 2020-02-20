package resolver

type QueryResolver struct{}
type MutationResolver struct{}
type TodoQueryResolver struct{}
type TodoMutationResolver struct{}

func (r *QueryResolver) Todo() *TodoQueryResolver {
	return &TodoQueryResolver{}
}

func (r *MutationResolver) Todo() *TodoMutationResolver {
	return &TodoMutationResolver{}
}
