package database

import (
	"clean_arch/BookStore/intrface"
	"context"
	"database/sql"
	"fmt"
)

type BookDataLayerImpl struct {
	Dbconn *sql.DB
}

func NewBookDatalayerImpl() intrface.BookDataLayer {
	return &BookDataLayerImpl{}
}
func (bookDL *BookDataLayerImpl) GetBookById(ctx context.Context, bookID int16) {

}

func (dataLayer *BookDataLayerImpl) TestBookDataLayer(ctx context.Context) {

	fmt.Println("***** Inside Book Datalayer *****")
}
