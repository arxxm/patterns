package file

import "os"

type Opcs struct {
	GID         int
	UID         int
	Content     string
	Flags       int
	Permissions os.FileMode
}

type fopc func(*Opcs)

func NewContent(content string) fopc {
	return func(o *Opcs) {
		o.Content = content
	}
}

func NewGID(gid int) fopc {
	return func(o *Opcs) {
		o.GID = gid
	}
}

func NewUID(uid int) fopc {
	return func(o *Opcs) {
		o.UID = uid
	}
}

func NewFile(filepath string, fuopc ...fopc) (*Opcs, error) {

	argOpt := &Opcs{
		GID:         os.Getgid(),
		UID:         os.Getgid(),
		Content:     "",
		Permissions: 0666,
		Flags:       os.O_CREATE | os.O_EXCL | os.O_WRONLY,
	}

	for _, onef := range fuopc {
		onef(argOpt)
	}

	return argOpt, nil

}
