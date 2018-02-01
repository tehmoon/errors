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

  if orig == nil {
    return err
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

  if err.Orig == nil {
    return fmt.Sprintf("%s", err.Err)
  }

  return fmt.Sprintf("%s: %s", err.Err, err.Orig.Error())
}

func (err Error) Root() (error) {
  if orig, ok := err.Orig.(*Error); ok {
    return orig.Root()
  }

  return err.Orig
}

func (err Error) Has(T interface{}) (bool) {
  if err.Err == T {
    return true
  }

  if orig, ok := err.Orig.(*Error); ok {
    return orig.Has(T)
  }

  if err.Orig == T {
    return true
  }

  return false
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
