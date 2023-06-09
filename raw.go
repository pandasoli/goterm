package goterm
import (
  "golang.org/x/sys/unix"
  "os"
)


func SetRawMode() (termios *unix.Termios, err error) {
  fd := int(os.Stdin.Fd())
  termios, err = unix.IoctlGetTermios(fd, unix.TCGETS)
  if err != nil {
    return nil, err
  }

  rawTermios := termios
  rawTermios.Lflag &^= unix.ICANON | unix.ECHO | unix.ISIG
  rawTermios.Cc[unix.VMIN] = 1
  rawTermios.Cc[unix.VTIME] = 0
  if err := unix.IoctlSetTermios(fd, unix.TCSETS, rawTermios); err != nil {
    return nil, err
  }

  return termios, nil
}

func RestoreMode(termios *unix.Termios) error {
  fd := int(os.Stdin.Fd())
  if err := unix.IoctlSetTermios(fd, unix.TCSETS, termios); err != nil {
    return err
  }

  return nil
}
