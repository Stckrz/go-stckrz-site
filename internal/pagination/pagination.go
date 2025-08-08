package pagination

import (
	"fmt"
	"strconv"
)

type Result[T any] struct {
	Items       []T   //The slice of items to paginate
	CurrentPage int   //The raw string of the current page
	TotalPages  int   //The total number of pages
	PageNumbers []int //An array of the page numbers
}

//Paginate takes an array of items and slices it into pages of size 'pageSize', 
//using the raw string 'pageString' from the url query to decide which posts to show you.
//
//Bad, missing, or negative page strings will be defaulted to 1. Page numbers greater than
//the total number of pages will be defaulted to the total number of pages.
//TotalPages is insured to be at least 1.
func Paginate[T any](itemsArray []T, pageString string, pageSize int) Result[T] {
	total := len(itemsArray)
	totalPages := (total + pageSize - 1) / pageSize
	//if total pages is 0, default to 1
	if totalPages == 0 {
		totalPages = 1
	}

	//default current page to 1
	page := 1

	//check that the page is parse-able and that it is greater than 0
	if parsedPage, err := strconv.Atoi(pageString); err == nil && parsedPage > 0 {
		page = parsedPage
	}

	//make sure page is not greater than total number of pages
	if page > totalPages {
		page = totalPages
	}

	//calculate start and end indices, ensuring that neither is higher than the total number of items
	start := (page - 1) * pageSize
	end := start + pageSize
	if start > total {
		start = total
	}
	if end > total {
		end = total
	}

	pages := make([]int, totalPages)
	for i := range pages {
		pages[i] = i + 1
	}
	fmt.Println(itemsArray[start:end])

	return Result[T]{
		Items:       itemsArray[start:end],
		CurrentPage: page,
		TotalPages:  totalPages,
		PageNumbers: pages,
	}
}
