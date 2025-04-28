func out[T any](v T, err error) (T, error) {
  if err != nil {
    var zv T
    return zv, err
  }
  return v, nil
}
