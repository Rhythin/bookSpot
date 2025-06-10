package packets

type CreateNotificationDetails struct {
	UserIDs       []string
	BookID        string
	BookTitle     string
	ChapterTitle  string
	ChapterNumber int
}
