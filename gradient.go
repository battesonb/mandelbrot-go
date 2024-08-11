package main

type GradientColor struct {
	Stop float64
	R    uint8
	G    uint8
	B    uint8
}

var (
	GRADIENT = []GradientColor{
		{
			Stop: 0,
			R:    45,
			G:    45,
			B:    150,
		},
		{
			Stop: 0.02,
			R:    100,
			G:    240,
			B:    210,
		},
		{
			Stop: 0.025,
			R:    250,
			G:    240,
			B:    180,
		},
		{
			Stop: 0.03,
			R:    250,
			G:    190,
			B:    180,
		},
		{
			Stop: 0.055,
			R:    90,
			G:    70,
			B:    180,
		},
		{
			Stop: 0.06,
			R:    30,
			G:    30,
			B:    90,
		},
		{
			Stop: 0.2,
			R:    220,
			G:    150,
			B:    140,
		},
		{
			Stop: 0.5,
			R:    255,
			G:    255,
			B:    255,
		},
		{
			Stop: 1.0,
			R:    0,
			G:    0,
			B:    0,
		},
	}
)
