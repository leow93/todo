package remote

import (
	"bytes"
	"io"

	"github.com/jlaffaye/ftp"
)

type Remote struct {
	Location string
	Username string
	Password string
}

func New(location, username, pwd string) *Remote {
	return &Remote{
		Location: location,
		Username: username,
		Password: pwd,
	}
}

func (r *Remote) Sync(filename string, data []byte) error {
	c, err := ftp.Dial(r.Location)
	if err != nil {
		return err
	}

	err = c.Login(r.Username, r.Password)
	if err != nil {
		return err
	}

	err = c.Stor(filename, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	return c.Quit()
}

func (r *Remote) Read(filename string) ([]byte, error) {
	c, err := ftp.Dial(r.Location)
	if err != nil {
		return nil, err
	}

	defer c.Quit()

	err = c.Login(r.Username, r.Password)
	if err != nil {
		return nil, err
	}

	rem, err := c.Retr(filename)
	if err != nil {
		return nil, err
	}
	defer rem.Close()

	return io.ReadAll(rem)
}
