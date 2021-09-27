package errors

// Error implementation of error interface
func (cli *ccError) Error() string {
    return cli.callback()
}

// GetCode returns error code
func (cli *ccError) GetCode() int {
    return cli.code
}
