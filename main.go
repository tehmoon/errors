package errors

import (
  "fmt"
  "errors"
)

type Error struct {
  Orig error
  Err error
}

func WrapErr(orig error, err error) (error) {
  if err == nil {
    return orig
  }

  return &Error{
    Orig: orig,
    Err: err,
  }
}

func (err Error) Error() (string) {
  if orig, ok := err.Orig.(*Error); ok {
    return fmt.Sprintf("%s: %s", err.Err.Error(), orig.Error())
  }

  return fmt.Sprintf("%s: %s", err.Err, err.Orig.Error())
}

func New(message string) (error) {
  return errors.New(message)
}

func Errorf(message string, args ...interface{}) (error) {
  return fmt.Errorf(message, args...)
}

func Wrap(orig error, message string) (error) {
  return &Error{
    Orig: orig,
    Err: New(message),
  }
}

func Wrapf(orig error, message string, args ...interface{}) (error) {
  return &Error{
    Orig: orig,
    Err: Errorf(message, args...),
  }
}
