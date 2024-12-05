package utils

import (
	"fmt"
	"strings"
	"time"
)

type Tqdm struct {
	total   int
	current int
	barLen  int
	prefix  string
	suffix  string
}

func NewTqdm(total int, barLen int, prefix, suffix string) *Tqdm {
	return &Tqdm{
		total:   total,
		barLen:  barLen,
		prefix:  prefix,
		suffix:  suffix,
	}
}

func (t *Tqdm) Update(n int) {
	t.current += n
	t.render()
}

// HACK
func (t *Tqdm) render() {
	percent := float64(t.current) / float64(t.total)
	filledLen := int(percent * float64(t.barLen))
	bar := strings.Repeat("=", filledLen) + strings.Repeat(" ", t.barLen-filledLen)
	fmt.Printf("\r%s [%s] %3.2f%% %s", t.prefix, bar, percent*100, t.suffix)
	if t.current >= t.total {
		fmt.Println()
	}
}

func (t *Tqdm) Close() {
	t.current = t.total
	t.render()
}
