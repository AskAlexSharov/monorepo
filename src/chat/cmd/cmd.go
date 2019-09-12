// Package cmd has all top-level commands dispatched by main's flag.Parse
// The entry point of each command is Execute function
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// CommonOptionsCommander extends flags.Commander with SetCommon
// All commands should implement this interfaces
type CommonOptionsCommander interface {
	SetCommon(commonOpts CommonOpts)
	Execute(args []string) error
}

// CommonOpts sets externally from main, shared across all commands
type CommonOpts struct {
	SharedSecret string
	Revision     string
}

// SetCommon satisfies CommonOptionsCommander interface and sets common option fields
// The method called by main for each command
func (c *CommonOpts) SetCommon(commonOpts CommonOpts) {
	c.SharedSecret = commonOpts.SharedSecret
	c.Revision = commonOpts.Revision
}

// resetEnv clears sensitive env vars
func resetEnv(envs ...string) {
	for _, env := range envs {
		if err := os.Unsetenv(env); err != nil {
			log.Printf("[WARN] can't unset env %s, %s", env, err)
		}
	}
}

// responseError returns error with status and response body
func responseError(resp *http.Response) error {
	body, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		body = []byte("")
	}

	return fmt.Errorf("error response %q, %s", resp.Status, body)
	//return errors.Errorf("error response %q, %s", resp.Status, body)
}

// mkdir -p for all dirs
func makeDirs(dirs ...string) error {
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0700); err != nil { // If path is already a directory, MkdirAll does nothing
			return fmt.Errorf("can't make directory %s: %w", dir, err)
		}
	}
	return nil
}
