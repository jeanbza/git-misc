package books

type Book struct {
    Title string
    Author string
    Pages int
}

func (b Book) CategoryByLength() string {
    if (b.Pages < 300) {
        return "SHORT STORY"
    } else {
        return "NOVEL"
    }
}
