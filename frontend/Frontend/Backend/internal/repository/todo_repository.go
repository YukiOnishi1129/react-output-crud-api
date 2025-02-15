type TodoRepository interface {
	FindAll(ctx context.Context) ([]*domain.Todo, error)
	FindByID(ctx context.Context, input dto.FindByIDInput) (*domain.Todo, error)
	Create(ctx context.Context, todo *domain.Todo) error
	Update(ctx context.Context, todo *domain.Todo) error
	Delete(ctx context.Context, id uuid.UUID) error
} 