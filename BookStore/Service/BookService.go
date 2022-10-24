package service

import (
	"clean_arch/BookStore/intrface"
	"clean_arch/BookStore/model"
	"context"
)

type BookServiceImpl struct {
	bookDL intrface.BookDataLayer
}

// this function will create bookService object which represents the BookService Interface
func NewBookServiceImpl(bookDataLayer intrface.BookDataLayer) intrface.BookService {
	return &BookServiceImpl{
		bookDL: bookDataLayer,
	}
}

func (service *BookServiceImpl) PrintBookTitle(ctx context.Context, book *model.Book)
