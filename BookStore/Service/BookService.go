package service

import (
	"clean_arch/BookStore/intrface"
	"clean_arch/BookStore/model"
	"context"
	"fmt"
)

type BookServiceImpl struct {
	bookDataLayer intrface.BookDataLayer
}

// this function will create bookService object which represents the BookService Interface
func NewBookServiceImpl(bookDataLayer intrface.BookDataLayer) intrface.BookService {
	return &BookServiceImpl{
		bookDataLayer: bookDataLayer,
	}
}

func (service *BookServiceImpl) PrintBookTitle(ctx context.Context, book *model.Book) {

}

func (service *BookServiceImpl) TestBookService(ctx context.Context) {

	fmt.Println("****** Inside Book Controller ******")
	service.bookDataLayer.TestBookDataLayer(ctx)

}
