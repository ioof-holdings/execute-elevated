package main

func platformElevate(command string, args []string) error {
	sudoArgs := append([]string{"-n", command}, args...)
	return plainExec("sudo", sudoArgs)
}
