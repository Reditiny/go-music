package vo

type BannerVo struct {
	ID  uint   `json:"id" copier:"must"`
	Pic string `json:"pic" copier:"must"`
}
