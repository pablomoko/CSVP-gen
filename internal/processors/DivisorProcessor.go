package processors

import "strconv"

type DivisorProcessor struct {
	Divisor int
}

func (dp DivisorProcessor) Process(value string) string {
	num, err := strconv.Atoi(value)
	if err != nil {
		return "Error: no se pudo convertir el valor a número"
	}
	result := num / dp.Divisor
	return strconv.Itoa(result)
}
