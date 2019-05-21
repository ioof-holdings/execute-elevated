package main

func platformElevate(command string, args []string) error {
	// elevation is handled by our manifest file
	return plainExec(command, args)
}
