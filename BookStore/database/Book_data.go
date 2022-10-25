package database

import (
	"clean_arch/BookStore/intrface"
	"context"
	"database/sql"
)

type BookDataLayerImpl struct {
	Dbconn *sql.DB
}

func NewBookDatalayerImpl(Conn *sql.DB) intrface.BookDataLayer {
	return &BookDataLayerImpl{}
}
func (bookDL *BookDataLayerImpl) GetBookById(ctx context.Context, bookID int16)

func (dataLayer *BookDataLayerImpl) TestBookDataLayer(ctx context.Context) {

}
