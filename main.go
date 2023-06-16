package main

func main() {
	err := rootCmd.Execute()
	checkErrorDie(err)
}
