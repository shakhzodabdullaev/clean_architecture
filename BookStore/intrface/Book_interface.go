package intrface

import (
	"clean_arch/BookStore/model"
	"context"
)

type BookService interface {
	PrintBookTitle(ctx context.Context, book *model.Book)
	TestBookService(ctx context.Context)
}
type BookDataLayer interface {
	GetBookById(ctx context.Context, bookID int16)
	TestBookDataLayer(ctx context.Context)
}
