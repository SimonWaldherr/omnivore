package ossh

import (
	"github.com/discoriver/massh"
	"github.com/discoriver/omnivore/internal/log"
	"golang.org/x/crypto/ssh"
)

type OmniSSHConfig struct {
	Config     *massh.Config
	StreamChan chan massh.Result
}

func NewConfig() *OmniSSHConfig {
	c := &OmniSSHConfig{
		Config:     massh.NewConfig(),
		StreamChan: make(chan massh.Result),
	}
	return c
}

// Stream executes work contained in the massh.Config, and returns a StreamCycle for monitoring output and status.
func (c *OmniSSHConfig) Stream() (*StreamCycle, error) {
	err := c.Config.Stream(c.StreamChan)
	if err != nil {
		return nil, err
	}

	log.OmniLog.Info("Massh Streaming started successfully.")

	ss := newStreamCycle(c.StreamChan, len(c.Config.Hosts))
	return ss, nil
}

func (c *OmniSSHConfig) AddHosts(h []string) {
	c.Config.SetHosts(h)
}

func (c *OmniSSHConfig) AddSSHConfig(s *ssh.ClientConfig) {
	c.Config.SetSSHConfig(s)
}

func (c *OmniSSHConfig) AddJob(j *massh.Job) {
	c.Config.SetJob(j)
}

func (c *OmniSSHConfig) AddBastionHost(b string) {
	c.Config.SetBastionHost(b)
}

func (c *OmniSSHConfig) AddBastionHostConfig(s *ssh.ClientConfig) {
	c.Config.SetBastionHostConfig(s)
}

func (c *OmniSSHConfig) AddWorkerPool(w int) {
	c.Config.SetWorkerPool(w)
}

func (c *OmniSSHConfig) AddPasswordAuth(user string, password string) {
	c.Config.SSHConfig.User = user
	c.Config.SetPasswordAuth([]byte(password))
}

func (c *OmniSSHConfig) AddPublicKeyAuth(k string, p string) (err error) {
	if err = c.Config.SetPublicKeyAuth(k, p); err != nil {
		return err
	}

	return nil
}

func (c *OmniSSHConfig) AddSSHSockAuth() error {
	c.Config.SetSSHAuthSockAuth()

	return nil
}