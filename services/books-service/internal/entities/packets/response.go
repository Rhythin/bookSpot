package packets

type BookDetails struct {
	ID           string `json:"id" gorm:"column:id"`
	Title        string `json:"title" gorm:"column:title"`
	Author       string `json:"author" gorm:"column:author"`
	Description  string `json:"description" gorm:"column:description"`
	ChapterCount int64  `json:"chapterCount" gorm:"column:chapter_count"`
}

type ListBooksResponse struct {
	Books       []*BookDetails `json:"books"`
	TotalCount  int64          `json:"totalCount"`
	SearchCount int64          `json:"searchCount"`
}

type ChapterDetails struct {
	Title  string `json:"title" gorm:"column:title"`
	Number int    `json:"number" gorm:"column:number"`
}
type ListChaptersResponse struct {
	Chapters    []*ChapterDetails `json:"chapters"`
	TotalCount  int64             `json:"totalCount"`
	SearchCount int64             `json:"searchCount"`
}

type ReadingListEntryDetails struct {
	BookID          string `json:"bookID" gorm:"column:book_id"`
	LastReadChapter string `json:"lastReadChapter" gorm:"column:last_read_chapter"`
	ChapterCount    int64  `json:"chapterCount" gorm:"-"`
}

type ListReadingListResponse struct {
	Entries     []*ReadingListEntryDetails `json:"entries"`
	TotalCount  int64                      `json:"totalCount"`
	SearchCount int64                      `json:"searchCount"`
}
