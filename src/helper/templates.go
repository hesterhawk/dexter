package helper

func TmplRange(start, end int) (stream chan int) {
  stream = make(chan int)
  go func() {
      for i := start; i <= end; i++ {
          stream <- i
      }
      close(stream)
  }()
  return
}