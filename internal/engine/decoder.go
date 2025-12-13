package engine

import "fmt"

func decodeFileContent(content []byte, extension string) string {
	switch extension {
	case ".md":
		return string(content)
	}

	fmt.Printf("decodeFileContent: unknown file type, extension=%s\n", extension)
	return ""
}
