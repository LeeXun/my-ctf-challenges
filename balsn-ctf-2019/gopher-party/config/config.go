package config

import "unicode"

var (
	Redis   RedisConfig
	App     AppConfig
	Scripts = map[string]*unicode.RangeTable{
		"Adlam":                  unicode.Adlam,
		"Ahom":                   unicode.Ahom,
		"Anatolian_Hieroglyphs":  unicode.Anatolian_Hieroglyphs,
		"Arabic":                 unicode.Arabic,
		"Armenian":               unicode.Armenian,
		"Avestan":                unicode.Avestan,
		"Balinese":               unicode.Balinese,
		"Bamum":                  unicode.Bamum,
		"Bassa_Vah":              unicode.Bassa_Vah,
		"Batak":                  unicode.Batak,
		"Bengali":                unicode.Bengali,
		"Bhaiksuki":              unicode.Bhaiksuki,
		"Bopomofo":               unicode.Bopomofo,
		"Brahmi":                 unicode.Brahmi,
		"Braille":                unicode.Braille,
		"Buginese":               unicode.Buginese,
		"Buhid":                  unicode.Buhid,
		"Canadian_Aboriginal":    unicode.Canadian_Aboriginal,
		"Carian":                 unicode.Carian,
		"Caucasian_Albanian":     unicode.Caucasian_Albanian,
		"Chakma":                 unicode.Chakma,
		"Cham":                   unicode.Cham,
		"Egyptian_Hieroglyphs":   unicode.Egyptian_Hieroglyphs,
		"Elbasan":                unicode.Elbasan,
		"Ethiopic":               unicode.Ethiopic,
		"Georgian":               unicode.Georgian,
		"Glagolitic":             unicode.Glagolitic,
		"Gothic":                 unicode.Gothic,
		"Grantha":                unicode.Grantha,
		"Greek":                  unicode.Greek,
		"Gujarati":               unicode.Gujarati,
		"Gurmukhi":               unicode.Gurmukhi,
		"Han":                    unicode.Han,
		"Hangul":                 unicode.Hangul,
		"Hanunoo":                unicode.Hanunoo,
		"Hatran":                 unicode.Hatran,
		"Hebrew":                 unicode.Hebrew,
		"Hiragana":               unicode.Hiragana,
		"Imperial_Aramaic":       unicode.Imperial_Aramaic,
		"Inherited":              unicode.Inherited,
		"Inscriptional_Pahlavi":  unicode.Inscriptional_Pahlavi,
		"Inscriptional_Parthian": unicode.Inscriptional_Parthian,
		"Javanese":               unicode.Javanese,
		"Kaithi":                 unicode.Kaithi,
		"Kannada":                unicode.Kannada,
		"Katakana":               unicode.Katakana,
		"Kayah_Li":               unicode.Kayah_Li,
		"Kharoshthi":             unicode.Kharoshthi,
		"Khmer":                  unicode.Khmer,
		"Khojki":                 unicode.Khojki,
		"Khudawadi":              unicode.Khudawadi,
		"Lao":                    unicode.Lao,
		"Latin":                  unicode.Latin,
		"Lepcha":                 unicode.Lepcha,
		"Limbu":                  unicode.Limbu,
		"Lisu":                   unicode.Lisu,
		"Malayalam":              unicode.Malayalam,
		"Mandaic":                unicode.Mandaic,
		"Manichaean":             unicode.Manichaean,
		"Miao":                   unicode.Miao,
		"Modi":                   unicode.Modi,
		"Mongolian":              unicode.Mongolian,
		"Mro":                    unicode.Mro,
		"Multani":                unicode.Multani,
		"Myanmar":                unicode.Myanmar,
		"Nabataean":              unicode.Nabataean,
		"New_Tai_Lue":            unicode.New_Tai_Lue,
		"Newa":                   unicode.Newa,
		"Nko":                    unicode.Nko,
		"Nushu":                  unicode.Nushu,
		"Ogham":                  unicode.Ogham,
		"Oriya":                  unicode.Oriya,
		"Rejang":                 unicode.Rejang,
		"Runic":                  unicode.Runic,
		"Syloti_Nagri":           unicode.Syloti_Nagri,
		"Syriac":                 unicode.Syriac,
		"Thai":                   unicode.Thai,
		"Yi":                     unicode.Yi,
	}
)

type RedisConfig struct {
	Prefix       string
	UserKey      string
	UserCountKey string
	UserEmailKey string
}

type AppConfig struct {
	Production string
}

func init() {
	prefix := "gopherparty:"
	Redis = RedisConfig{
		Prefix:       prefix,
		UserKey:      prefix + "users:",
		UserCountKey: prefix + "usercount:",
		UserEmailKey: prefix + "useremails:",
	}
	App = AppConfig{
		Production: "production",
	}
}
