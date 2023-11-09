package sample

var FormServiceMockSingleton = &FormServiceMock{}

type FormService interface {
	SaveForm(formId string, form map[string]any) error
	LoadForm(formId string) (map[string]any, error)
}

var FromData map[string]map[string]any

func init() {
	FromData = map[string]map[string]any{}
}

type FormServiceMock struct {
}

func (s *FormServiceMock) SaveForm(formId string, form map[string]any) error {
	FromData[formId] = form
	return nil
}

func (s *FormServiceMock) LoadForm(formId string) (map[string]any, error) {
	if form, ok := FromData[formId]; ok {
		return form, nil
	}
	return nil, nil
}