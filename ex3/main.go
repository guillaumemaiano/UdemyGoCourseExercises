package main

// for printing data out
import "fmt"
// for reading the cli
import "os"
// for logging errors and actions
import "log"
// for dumping data from the file to somewhere else - eg, Copy
import "io"

func main() {

		arguments := os.Args
		if  (len(arguments) != 2 ) {
			fmt.Println("The Program requires a file path string as an argument.");
			os.Exit(-1);
		}
		fmt.Println("Hello");
		// totally safe, no way this could be a bad idea, right? 
		filename := arguments[1]
		// Using func Open(name string) (*File, error)
		file, err := os.Open(filename)
		if err != nil {
			log.Println(err)
		}
		// Dump the file contents to the screen, woohoo.
		log.Println("####Contents of file %s", filename)
		// Copy(dest, src)
		io.Copy(os.Stdout, file)
}
