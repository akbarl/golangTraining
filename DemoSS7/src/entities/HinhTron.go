package entities

import "math"

type HinhTron struct {
	BanKinh int
	Ten     string
}

func (hinhTron HinhTron) ChuVi() float64 {
	return float64(2) * math.Pi * float64(hinhTron.BanKinh)
}

func (hinhTron HinhTron) DienTich() float64 {
	return math.Pi * math.Pow(float64(hinhTron.BanKinh), float64(2))
}

func (hinhTron HinhTron) TenHinh() string {
	return hinhTron.Ten
}
