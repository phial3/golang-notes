package area

import (
	"encoding/json"
	"github.com/phial3/accumulate/file"
	"strings"
)

type Zone struct {
	Citycode  interface{} `json:"citycode"`
	Adcode    string      `json:"adcode"`
	Name      string      `json:"name"`
	Center    string      `json:"center"`
	Level     string      `json:"level"`
	Districts []*Zone     `json:"districts"`
}

func Load(path string) *Zone {
	data, err := file.ReadByteFromFile(path)
	if err != nil {
		return nil
	}

	var result Zone

	_ = json.Unmarshal(data, &result)
	//fmt.Println(result.Districts[0].Name)
	return &result
}

func (z *Zone) GetName() string {
	if z != nil {
		return z.Name
	}
	return ""
}

func (z *Zone) GetLocation() (string, string) {
	if z != nil {
		lonAndLat := strings.Split(z.Center, ",")
		if len(lonAndLat) == 2 {
			return lonAndLat[1], lonAndLat[0]
		}
		return "0", "0"
	}
	return "0", "0"
}

func (z *Zone) GetLevel() bool {
	return true
}

func (z *Zone) isCountry() bool {
	return z.Level == "country"
}

func (z *Zone) isProvince() bool {
	return z.Level == "province"
}

func (z *Zone) isCity() bool {
	return z.Level == "city"
}

func (z *Zone) isDistrictOrStreet() bool {
	return z.Level == "district" || z.Level == "street"
}

func (z *Zone) GetZoneByCityCode(code string) *Zone {
	if z.Citycode == code {
		return z
	}

	for _, e := range z.Districts {
		if e.Citycode == code {
			return e
		}

		for _, ee := range e.Districts {
			if ee.Citycode == code {
				return ee
			}

			for _, eee := range ee.Districts {
				if eee.Citycode == code {
					return eee
				}
			}
		}

	}
	return nil
}

func (z *Zone) GetZoneByAdCode(code string) *Zone {
	if z.Adcode == code {
		return z
	}

	for _, e := range z.Districts {
		if e.Adcode == code {
			return e
		}

		for _, ee := range e.Districts {
			if ee.Adcode == code {
				return ee
			}

			for _, eee := range ee.Districts {
				if eee.Adcode == code {
					return eee
				}
			}
		}

	}
	return nil
}

func (z *Zone) GetFather(obj *Zone) *Zone {
	if z == obj {
		return nil
	}

	for _, e := range z.Districts {
		if e == obj {
			return z
		}

		for _, ee := range e.Districts {
			if ee == obj {
				return e
			}

			for _, eee := range ee.Districts {
				if eee == obj {
					return ee
				}
			}
		}

	}
	return nil
}

func (z *Zone) GetChildren(obj *Zone) []*Zone {
	if z == obj {
		ztmp := make([]*Zone, 0)
		return append(ztmp, z)
	}

	for _, e := range z.Districts {
		if e == obj {
			return z.Districts
		}

		for _, ee := range e.Districts {
			if ee == obj {
				return ee.Districts
			}

			for _, eee := range ee.Districts {
				if eee == obj {
					return eee.Districts
				}
			}
		}

	}
	return nil
}
