# File Search CLI

This is a command-line application written in Go that searches for a string in a large file.

## Usage

Make sure you have Go installed on your system. To compile and run the program, follow these steps:

1. Clone the repository or download the source code.

2. Open a terminal or command prompt and navigate to the project directory.

3. Run the following command to compile and run the program:

      go run main.go <filename> <search_string>
  
      Ex:go run finder.go data.txt sample
  

## Limitations

- The program assumes that the file is in a readable text format. If the file is in a binary format, the program may not work as expected.

- This program reads the file line by line, so it may take longer to search for the string in very large files.

- The search is case-sensitive. To perform a case-insensitive search, you can modify the code in the `finder.go` file to convert both the line and the search string to lowercase or uppercase before comparison.

## License

This project is licensed under the [MIT License](LICENSE).
