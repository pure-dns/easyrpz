package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// StringSliceFlag is a custom flag type for string slices
type StringSliceFlag []string

// String returns the string representation of the flag value
func (s *StringSliceFlag) String() string {
	return fmt.Sprintf("%v", *s)
}

// Set adds a value to the flag
func (s *StringSliceFlag) Set(value string) error {
	*s = append(*s, value)
	return nil
}

func main() {
	// Define command-line flags
	var inputFiles StringSliceFlag
	var excludeFiles StringSliceFlag
	outputFlag := flag.String("o", "", "Output file path")
	flag.Var(&inputFiles, "i", "Input file paths or URLs")
	flag.Var(&excludeFiles, "e", "File paths or URLs containing domains to exclude")
	flag.Parse()

	// Check if input files or URLs are provided
	if len(inputFiles) == 0 {
		fmt.Println("Input file paths or URLs are required")
		return
	}

	// Check if output file path is provided
	if *outputFlag == "" {
		fmt.Println("Output file path is required")
		return
	}

	// Create the output file
	outputFile, err := os.Create(*outputFlag)
	if err != nil {
		fmt.Println("Error creating the output file:", err)
		return
	}
	defer outputFile.Close()

	// Write the RPZ file header
	header := generateRPZHeader()
	_, err = fmt.Fprint(outputFile, header)
	if err != nil {
		fmt.Println("Error writing to the output file:", err)
		return
	}

	// Create a set to store excluded domains
	excludedDomains := make(map[string]bool)
	// Create a set to store unique hosts
	uniqueHosts := make(map[string]bool)

	// Process each exclude file or URL
	for _, excludeFile := range excludeFiles {
		// Trim whitespace from the input
		excludeFile = strings.TrimSpace(excludeFile)

		// Read the exclude file or URL
		var file io.Reader
		if strings.HasPrefix(excludeFile, "http://") || strings.HasPrefix(excludeFile, "https://") {
			resp, err := http.Get(excludeFile)
			if err != nil {
				fmt.Println("Error retrieving the exclude file:", err)
				continue
			}
			file = resp.Body
		} else {
			file, err = os.Open(excludeFile)
			if err != nil {
				fmt.Println("Error opening the exclude file:", err)
				continue
			}
		}

		scanner := bufio.NewScanner(file)

		// Process each line in the exclude file
		for scanner.Scan() {
			domain := strings.TrimSpace(scanner.Text())

			// Skip empty lines and comments
			if domain == "" || strings.HasPrefix(domain, "#") {
				continue
			}

			// Add domain to the excluded domains set
			excludedDomains[domain] = true
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading the exclude file:", err)
			continue
		}

		if closer, ok := file.(io.Closer); ok {
			closer.Close()
		}
	}

	// Process each input file or URL
	for _, input := range inputFiles {
		// Trim whitespace from the input
		input = strings.TrimSpace(input)

		// Read the input file or URL
		var file io.Reader
		if strings.HasPrefix(input, "http://") || strings.HasPrefix(input, "https://") {
			resp, err := http.Get(input)
			if err != nil {
				fmt.Println("Error retrieving the input file:", err)
				continue
			}
			file = resp.Body
		} else {
			file, err = os.Open(input)
			if err != nil {
				fmt.Println("Error opening the input file:", err)
				continue
			}
		}

		scanner := bufio.NewScanner(file)

		// Process each line in the input file
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())

			// Skip empty lines and comments
			if line == "" || strings.HasPrefix(line, "#") {
				continue
			}

			// Skip comments in the middle of the line
			parts := strings.SplitN(line, "#", 2)
			line = strings.TrimSpace(parts[0])

			parts = strings.Fields(line)
			if len(parts) != 2 {
				fmt.Println("Invalid input line:", line)
				continue
			}

			ip := parts[0]
			host := parts[1]

			// Check if the domain should be excluded
			if excludedDomains[host] {
				continue
			}

			// Check if the host is already encountered
			if uniqueHosts[host] {
				continue
			}

			// Add host to the unique hosts set
			uniqueHosts[host] = true

			// Write the RPZ record with A record
			 _, err = fmt.Fprintf(outputFile, "%s IN A %s\n", host, ip)
			if err != nil {
				fmt.Println("Error writing to the output file:", err)
				return
			}

			// Check if IP is "0.0.0.0" and include wildcard record
			if ip == "0.0.0.0" {
				wildcardHost := "*." + host
				_, err = fmt.Fprintf(outputFile, "%s IN A %s\n", wildcardHost, ip)
				if err != nil {
					fmt.Println("Error writing to the output file:", err)
					return
				}
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading the input file:", err)
			continue
		}

		if closer, ok := file.(io.Closer); ok {
			closer.Close()
		}
	}

	fmt.Println("Conversion completed successfully!")
}

// generateRPZHeader generates the RPZ file header with the current date in the format YYYYMMDD.
func generateRPZHeader() string {
	date := time.Now().Format("20060102") // Format: YYYYMMDD

	header := fmt.Sprintf("$TTL 60\n"+
		"@    IN    SOA        localhost.  filters.puredns.org.  (\n"+
		"           %s   ;     serial\n"+
		"           2w         ;     refresh\n"+
		"           2w         ;     retry\n"+
		"           2w         ;     expiry\n"+
		"           2w         ;     minimum\n"+
		"           IN         NS    localhost.\n"+
		")\n\n", date)

	return header
}
