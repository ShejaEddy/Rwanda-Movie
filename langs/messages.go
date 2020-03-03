package langs

func GenerateValidationMessage(field string, rule string, errors map[string][]string) {
	switch rule {
	case "required":
		errors[field] = append(errors[field], "The "+field+" is required")
		break
	case "email":
		errors[field] = append(errors[field], "The "+field+" should be a valid email")
		break
	case "min":
		errors[field] = append(errors[field], "The "+field+" should be a greater than 6 characters")
		break
	default:
		errors[field] = append(errors[field], "The "+field+" is invalid")
		break
	}
}
