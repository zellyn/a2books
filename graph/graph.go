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

func (a *App) Book() BookResolver {
	return &bookResolver{a}
}

func (a *App) Chapter() ChapterResolver {
	return &chapterResolver{a}
}

func (a *App) Mutation() MutationResolver {
	return &mutationResolver{a}
}

func (a *App) Query() QueryResolver {
	return &queryResolver{a}
}

type bookResolver struct{ *App }

type chapterResolver struct{ *App }

type mutationResolver struct{ *App }

type queryResolver struct{ *App }

func (a *bookResolver) Chapters(ctx context.Context, obj *data.Book) ([]data.Chapter, error) {
	fmt.Println("bookResolver.Chapters")
	return nil, nil
}

func (a *chapterResolver) Pages(ctx context.Context, obj *data.Chapter) ([]data.Page, error) {
	fmt.Println("chapterResolver.Pages")
	return nil, nil
}

func (a *mutationResolver) CreateBook(ctx context.Context, slug string, input BookInput) (data.Book, error) {
	fmt.Println("mutationResolver.CreateBook")
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

	if err := a.Store.WriteBook(book); err != nil {
		return data.Book{}, err
	}

	return book, nil
}

func (a *mutationResolver) UpdateBook(ctx context.Context, slug string, input BookInput) (data.Book, error) {
	fmt.Println("mutationResolver.UpdateBook")
	book, err := a.Store.GetBook(slug)
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
		if err := a.Store.WriteBook(book); err != nil {
			return data.Book{}, err
		}
	}
	return book, nil
}

func (a *queryResolver) Book(ctx context.Context, slug string) (*data.Book, error) {
	book, err := a.Store.GetBook(slug)
	return &book, err
}

func (a *queryResolver) Books(ctx context.Context) ([]data.Book, error) {
	fmt.Println("queryResolver.Books")
	return a.Store.GetBooks()
}

func orEmpty(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}
