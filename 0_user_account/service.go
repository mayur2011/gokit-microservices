package account

import "context"

//Service which will be exposed to outside users
type Service interfacee{
	CreateUser(ctx context.Context, email string, password string) (string, error)
	GetUser(ctx context.Context, id string) (string, error)
}
