package service

import "week04/app/book/internal/data/ent"

type GRPCBookService struct {
	Client *ent.Client
}
