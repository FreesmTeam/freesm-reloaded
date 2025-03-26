package launcher

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func HomeDirectory() (string, error) {
	// btw thats OS
	const arch string = runtime.GOOS

	switch arch {
	case "linux":
		var dataHome string = os.Getenv("XDG_DATA_HOME")
		if dataHome == "" {
			dataHome = filepath.Join(os.Getenv("HOME"), ".local", "share")
		}
		return filepath.Join(dataHome, "SakuraLauncher"), nil
	case "windows":
		return filepath.Join(os.Getenv("APPDATA"), "SakuraLauncher"), nil
	case "darwin":
		return filepath.Join(os.Getenv("HOME"), "Library", "Application Support", "SakuraLauncher"), nil
	default:
		return "", fmt.Errorf("unsupported OS: %s", arch)
	}
}
