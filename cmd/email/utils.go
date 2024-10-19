package email

func parseTags(tags []string) string {
	var stringtags string = ""
	if len(tags) == 0 {
		return stringtags
	}
	for _, tag := range tags {
		stringtags += "Tag: " + tag + "\n"
	}
	return stringtags
}
