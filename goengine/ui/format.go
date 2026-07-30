package ui

// If absolute path is too long , this function should make it like (.../ProjectFolder/ProjectName)
func TruncatePath(path string, pathSize int) string {
	if len(path) <= pathSize {
		return path
	}

	return "..." + path[len(path)-pathSize:]

}
