package main

import "fmt"

// Report struct represents the complex object being built.
type Report struct {
	Title, Author, Content, Footer string
	FontSize                       int
	Color                          string
}

// ReportBuilder **Main Report Builder**
type ReportBuilder struct {
	report *Report
}

// NewReportBuilder initializes a new ReportBuilder
func NewReportBuilder() *ReportBuilder {
	return &ReportBuilder{&Report{}}
}

// Build returns the final constructed Report object
func (rb *ReportBuilder) Build() *Report {
	return rb.report
}

// ReportTitleContentBuilder Facet 1: Title & Content Builder
type ReportTitleContentBuilder struct {
	*ReportBuilder
}

func (rb *ReportBuilder) TitleContent() *ReportTitleContentBuilder {
	return &ReportTitleContentBuilder{rb}
}

func (r *ReportTitleContentBuilder) Title(title string) *ReportTitleContentBuilder {
	r.report.Title = title
	return r
}

func (r *ReportTitleContentBuilder) Content(content string) *ReportTitleContentBuilder {
	r.report.Content = content
	return r
}

// ReportAuthorBuilder Facet 2: Author Builder
type ReportAuthorBuilder struct {
	*ReportBuilder
}

func (rb *ReportBuilder) Author() *ReportAuthorBuilder {
	return &ReportAuthorBuilder{rb}
}

func (r *ReportAuthorBuilder) Name(author string) *ReportAuthorBuilder {
	r.report.Author = author
	return r
}

// ReportFormatBuilder Facet 3: Formatting Builder
type ReportFormatBuilder struct {
	*ReportBuilder
}

func (rb *ReportBuilder) Format() *ReportFormatBuilder {
	return &ReportFormatBuilder{rb}
}

func (r *ReportFormatBuilder) FontSize(size int) *ReportFormatBuilder {
	r.report.FontSize = size
	return r
}

func (r *ReportFormatBuilder) Color(color string) *ReportFormatBuilder {
	r.report.Color = color
	return r
}

// ReportFooterBuilder Facet 4: Footer Builder
type ReportFooterBuilder struct {
	*ReportBuilder
}

func (rb *ReportBuilder) Footer() *ReportFooterBuilder {
	return &ReportFooterBuilder{rb}
}

func (r *ReportFooterBuilder) Text(footer string) *ReportFooterBuilder {
	r.report.Footer = footer
	return r
}

// **Main Function**
func main() {
	report := NewReportBuilder().
		TitleContent().
		Title("Design Patterns").
		Content("Builder Pattern in Go").
		Author().
		Name("John Doe").
		Format().
		FontSize(12).
		Color("Blue").
		Footer().
		Text("Confidential").
		Build()

	fmt.Println(*report)
}
