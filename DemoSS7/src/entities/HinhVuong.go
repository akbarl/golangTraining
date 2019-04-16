package entities

type HinhVuong struct {
	Canh int
	Ten  string
}

func (hinhVuong HinhVuong) ChuVi() float64 {
	return float64(hinhVuong.Canh) * float64(4)
}

func (hinhVuong HinhVuong) DienTich() float64 {
	return float64(hinhVuong.Canh) * float64(hinhVuong.Canh)
}

func (hinhVuong HinhVuong) TenHinh() string {
	return hinhVuong.Ten
}
