//go:generate gqlgen

package graph

import (
	context "context"
	"fmt"

	data "github.com/zellyn/transcriber/data"
)

type App struct {
	Path  string
	Store *data.Store
}

func NewApp(path string) *App {
	return &App{
		Path:  path,
		Store: data.NewStore(path),
	}
}

func (app *App) Book_chapters(ctx context.Context, obj *data.Book) ([]data.Chapter, error) {
	fmt.Println("Book_chapters")
	return nil, nil
}

func (app *App) Chapter_pages(ctx context.Context, obj *data.Chapter) ([]data.Page, error) {
	fmt.Println("Chapter_pages")
	return nil, nil
}

func (app *App) Mutation_createBook(ctx context.Context, slug string, input BookInput) (data.Book, error) {
	fmt.Println("Mutation_createBook")
	book := data.Book{
		Slug:             slug,
		Title:            orEmpty(input.Title),
		Authors:          input.Authors,
		URL:              orEmpty(input.Url),
		ISBN:             orEmpty(input.Isbn),
		ImageURLTemplate: orEmpty(input.ImageUrlTemplate),
	}

	if book.Title == "" {
		return book, fmt.Errorf("cannot create book without a title")
	}

	if err := app.Store.WriteBook(book); err != nil {
		return data.Book{}, err
	}

	return book, nil
}

func (app *App) Mutation_updateBook(ctx context.Context, slug string, input BookInput) (data.Book, error) {
	fmt.Println("Mutation_updateBook")
	book, err := app.Store.GetBook(slug)
	if err != nil {
		return book, err
	}
	updated := false

	if input.Title != nil {
		if *input.Title == "" {
			return book, fmt.Errorf("cannot modify a book to have no title")
		}
		book.Title = *input.Title
		updated = true
	}
	if input.Authors != nil {
		book.Authors = input.Authors
		updated = true
	}
	if input.Url != nil {
		book.URL = *input.Url
		updated = true
	}
	if input.Isbn != nil {
		book.ISBN = *input.Isbn
		updated = true
	}
	if input.ImageUrlTemplate != nil {
		book.ImageURLTemplate = *input.ImageUrlTemplate
		updated = true
	}

	if updated {
		if err := app.Store.WriteBook(book); err != nil {
			return data.Book{}, err
		}
	}
	return book, nil
}

func (app *App) Query_book(ctx context.Context, slug string) (*data.Book, error) {
	book, err := app.Store.GetBook(slug)
	return &book, err
}

func (app *App) Query_books(ctx context.Context) ([]data.Book, error) {
	fmt.Println("Query_books")
	return app.Store.GetBooks()
}

func orEmpty(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}
