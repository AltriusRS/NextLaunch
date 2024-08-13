package utils

type TextEnum struct {
	text  string
	value uint
}

func DefineTextEnum(text string, value uint) TextEnum {
	return TextEnum{text, value}
}

// check if the value is equal

func (t TextEnum) Equal(value uint) bool {
	return t.value == value
}

func (t TextEnum) String() string {
	return t.text
}
