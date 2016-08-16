package main

func (p Panel) Columns(cols map[string]string) Panel {
	return Columns(p, cols)
}

func Columns(p Panel, cols map[string]string) Panel {
	if cols != nil {
		tempPanel := make(Panel)
		for header, series := range p {
			if val, ok := cols[header]; ok {
				tempPanel.Add(val, series)
			} else {
				tempPanel.Add(header, series)
			}
		}
		return tempPanel
	}
	return p
}

// func (p Panel) Columns(cols ...string) Panel {
// 	return Columns(p, cols...)
// }

// func Columns(p Panel, cols ...string) Panel {
// 	i := 0
// 	lim := len(cols)
// 	tempPanel := make(Panel)
// 	for header, series := range p {
// 		if i < lim {
// 			tempPanel.Add(cols[i], series)
// 			i += 1
// 		} else {
// 			tempPanel.Add(header, series)
// 		}
// 	}
// 	return tempPanel
// }
