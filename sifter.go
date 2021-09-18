package main

type (
	sifter interface {
		Sift(strings []string) []string
		String() string
	}

	sifterImp struct {
		stringBank []string
	}
)

func newSifter() sifter {
	return &sifterImp{
		stringBank: []string{},
	}
}

func (si *sifterImp) Sift(strings []string) []string {
	var flag bool
	result := []string{}
	for _, s := range strings {
		flag = false
		for _, t := range si.stringBank {
			flag = flag || (s == t)
		}
		if !flag {
			result = append(result, s)
		}
	}
	si.stringBank = append(si.stringBank, result...)
	return result
}

func (si *sifterImp) String() string {
	result := ""
	for _, v := range(si.stringBank) {
		result = result + v+ "\n"
	}
	return result
}