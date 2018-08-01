package data

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

var slugRe = regexp.MustCompile(`^[_a-zA-Z][_a-zA-Z0-9-]*$`)

type Store struct {
	Path string // The path to the base of the data store.
}

func NewStore(path string) *Store {
	return &Store{Path: path}
}

func (s *Store) GetBook(slug string) (Book, error) {
	if !slugRe.MatchString(slug) {
		return Book{}, fmt.Errorf("invalid book slug: %q", slug)
	}
	var b Book
	bb, err := ioutil.ReadFile(filepath.Join(s.Path, slug, "book.xml"))
	if err != nil {
		return b, err
	}

	if err := xml.Unmarshal(bb, &b); err != nil {
		return b, fmt.Errorf("error unmarshaling xml for %q: %v", slug, err)
	}

	return b, nil
}

func (s *Store) WriteBook(book Book) error {
	if !slugRe.MatchString(book.Slug) {
		return fmt.Errorf("invalid book slug: %q", book.Slug)
	}
	if err := os.MkdirAll(filepath.Join(s.Path, book.Slug), 0755); err != nil {
		return err
	}
	bb, err := xml.MarshalIndent(book, "", "  ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(filepath.Join(s.Path, book.Slug, "book.xml"), bb, 0644); err != nil {
		return err
	}
	return nil
}

func (s *Store) GetBooks() ([]Book, error) {
	files, err := ioutil.ReadDir(s.Path)
	if err != nil {
		return nil, err
	}

	var books []Book

	for _, f := range files {
		if f.IsDir() {
			b, err := s.GetBook(f.Name())
			if err != nil {
				return nil, err
			}
			books = append(books, b)
		}
	}

	return books, nil
}
