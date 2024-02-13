package mmaths

func LinearRegression(cs [][]float64) (m, b, r2 float64) {
	mx, my := 0., 0.
	for _, c := range cs {
		mx += c[0]
		my += c[1]
	}
	mx /= float64(len(cs))
	my /= float64(len(cs))

	m, r2 = func() (float64, float64) {
		sxy, sxx, syy := 0., 0., 0.
		for _, c := range cs {
			sxy += (c[0] - mx) * (c[1] - my) // covariance
			sxx += (c[0] - mx) * (c[0] - mx) // variance
			syy += (c[1] - my) * (c[1] - my) // variance
		}
		return sxy / sxx, func() float64 {
			if sxx > 0 && syy > 0 {
				if sxy < 0 {
					return -sxy * sxy / (sxx * syy)
				} else {
					return sxy * sxy / (sxx * syy)
				}
			} else {
				return -9999.
			}
		}()
	}()
	b = my - mx*m
	return
}
