package intrface

import (
	"clean_arch/BookStore/model"
	"context"
)

type BookService interface {
	PrintBookTitle(ctx context.Context, book *model.Book)
}
type BookDataLayer interface {
	GetBookById(ctx context.Context, bookID int16)
}
