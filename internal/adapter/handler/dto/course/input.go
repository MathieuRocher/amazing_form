package course

type CourseInput struct {
	Title       string `form:"title" validate:"required,min=3"`
	Description string `form:"description" validate:"required,min=10"`
}
