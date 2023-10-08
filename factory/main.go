package factory

func main() {
	message := "hello, world"

	stdoutPrinter := NewPrinter(Stdout)
	filePrinter := NewPrinter(File)

	stdoutPrinter.Output(message)
	filePrinter.Output(message)
}
