package main

import (
	"fmt"
	"github.com/markbates/grift/grift"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/js"
	"io/ioutil"
	"os"
	"path/filepath"
)

func BundleAndMinifyJS(c *grift.Context) error {
	sourceDir := "src"
	outputDir := "dist"
	outputFile := "bundle.js"

	// Ensure output directory exists
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return err
	}

	// Get all JavaScript files in the source directory
	matches, err := filepath.Glob(filepath.Join(sourceDir, "*.js"))
	if err != nil {
		return err
	}

	// Bundle and minify JavaScript files
	var bundledContent []byte
	m := minify.New()
	m.AddFunc("text/javascript", js.Minify)
	for _, file := range matches {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}
		bundledContent = append(bundledContent, content...)
	}

	// Minify the bundled content
	minifiedContent, err := m.Bytes("text/javascript", bundledContent)
	if err != nil {
		return err
	}

	// Write the minified content to the output file
	outputPath := filepath.Join(outputDir, outputFile)
	if err := ioutil.WriteFile(outputPath, minifiedContent, 0644); err != nil {
		return err
	}

	fmt.Printf("JavaScript files bundled and minified successfully. Output: %s\n", outputPath)
	return nil
}

func main() {
	// Register the task
	grift.Desc("bundle-js", "Bundle and minify JavaScript files")
	grift.Add("bundle-js", BundleAndMinifyJS)

	// Use the Grift CLI to run the task
	taskName := "bundle-js"
	context := &grift.Context{}
	if err := grift.Run(taskName, context); err != nil {
		if err.Error() == "task not found" {
			fmt.Println("Task not found.")
			os.Exit(1)
		}
		panic(err)
	}
}
