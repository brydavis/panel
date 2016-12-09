package main

type Panel map[string][]interface{}

func New(data interface{}) Panel {
	p := make(Panel)
	switch t := data.(type) {
	case map[string][]interface{}:
		p = Panel(p)
	case map[string][]string:
		for head := range t {
			p.Add(head, t[head])
		}
	case map[string][]bool:
		for head := range t {
			p.Add(head, t[head])
		}

	case map[string][]int:
		for head := range t {
			p.Add(head, t[head])
		}

	case map[string][]float64:
		for head := range t {
			p.Add(head, t[head])
		}
	default:

	}
	return p.Clean()

}

func (p Panel) Add(header string, data interface{}) Panel {
	switch t := data.(type) {
	case []interface{}:
		p[header] = t
	case []int:
		slice := make([]interface{}, len(t))
		for k, v := range t {
			slice[k] = v
		}
		p[header] = slice
	case []string:
		slice := make([]interface{}, len(t))
		for k, v := range t {
			slice[k] = v
		}
		p[header] = slice
	case []float64:
		slice := make([]interface{}, len(t))
		for k, v := range t {
			slice[k] = v
		}
		p[header] = slice
	case map[int]interface{}:
		slice := make([]interface{}, len(t))
		for k, v := range t {
			slice[k] = v
		}
		p[header] = slice

	}
	// todo: add map[int]interface...
	return p
}

func (p Panel) Select(cols ...string) Panel {
	tempPanel := make(Panel)
	if len(cols) == 0 || cols[0] == "*" {
		p.Clone(tempPanel)
	} else {
		for _, c := range cols {
			tempPanel.Add(c, p[c])
		}
	}
	return tempPanel
}

func (p Panel) Clone(dst Panel) {
	for header, series := range p {
		dst.Add(header, series)
	}
}
