package config

type Config struct {
	MaxWorkers int
	SearchWord string
	InputFiles []string
	OutputFile string
}

// LoadConfig loads the app's configuration
func LoadConfig() Config {
	return Config{
		MaxWorkers: 5,
		SearchWord: "Go", // Word to search in files
		InputFiles: []string{"C:/Users/LENOVO/Desktop/advancedProject/file1.txt", "C:/Users/LENOVO/Desktop/advancedProject/file2.txt", "C:/Users/LENOVO/Desktop/advancedProject/file3.txt"},
		OutputFile: "output/results.txt", // Output file
	}
}
