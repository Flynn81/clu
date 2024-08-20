package api

type Walker interface {
	Walk(path string) (int64, []int64, []string)
}
