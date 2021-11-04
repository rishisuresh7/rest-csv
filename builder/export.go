package builder

import "fmt"

type ExportBuilder interface {
	ExportView(viewName string) string
}

type exportBuilder struct{}

func NewExportBuilder() ExportBuilder {
	return &exportBuilder{}
}

func (e *exportBuilder) ExportView(viewName string) string {
	return fmt.Sprintf(`SELECT * FROM %s`, viewName)
}
