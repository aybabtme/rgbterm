// Package rgbterm colorizes bytes and strings using RGB colors, for a
// full range of pretty terminal strings.
//
// Beyond the traditional boring 16 colors of your terminal lie an
// extended set of 256 pretty colors waiting to be used. However, they
// are weirdly encoded; simply asking for an RGB color is much more
// convenient!
//
// RGB <-> HSL helpers were shamelessly taken from gorilla color, MIT
// licensed:
//    https://code.google.com/p/gorilla/source/browse/?r=ef489f63418265a7249b1d53bdc358b09a4a2ea0#hg%2Fcolor
package rgbterm

var (
	reset  = []byte("\033[0;00m")
	colors = termRGB[16:232]
)

// String colorizes the input with the terminal color that matches
// the closest the RGB color.
//
// This is simply a helper for Bytes.
func String(in string, r, b, g uint8) string {
	return string(Bytes([]byte(in), r, g, b))
}

// Bytes colorizes the input with the terminal color that matches
// the closest the RGB color.
func Bytes(in []byte, r, g, b uint8) []byte {
	return append(append(rgb(r, g, b), in...), reset...)
}

// Byte colorizes the input with the terminal color that matches
// the closest the RGB color.
func Byte(in byte, r, g, b uint8) []byte {
	return append(append(rgb(r, g, b), in), reset...)
}

func rgb(r, g, b uint8) []byte {
	// if all colors are equal, it might be in the grayscale range
	if r == g && g == b {
		color, ok := grayscale(r)
		if ok {
			return color
		}
	}

	// the general case approximates RGB by using the closest color.
	r6 := ((uint16(r) * 5) / 255)
	g6 := ((uint16(g) * 5) / 255)
	b6 := ((uint16(b) * 5) / 255)
	i := 36*r6 + 6*g6 + b6
	return colors[i]
}

func grayscale(scale uint8) ([]byte, bool) {
	switch scale {
	case 0x08:
		return termRGB[232], true
	case 0x12:
		return termRGB[233], true
	case 0x1c:
		return termRGB[234], true
	case 0x26:
		return termRGB[235], true
	case 0x30:
		return termRGB[236], true
	case 0x3a:
		return termRGB[237], true
	case 0x44:
		return termRGB[238], true
	case 0x4e:
		return termRGB[239], true
	case 0x58:
		return termRGB[240], true
	case 0x62:
		return termRGB[241], true
	case 0x6c:
		return termRGB[242], true
	case 0x76:
		return termRGB[243], true
	case 0x80:
		return termRGB[244], true
	case 0x8a:
		return termRGB[245], true
	case 0x94:
		return termRGB[246], true
	case 0x9e:
		return termRGB[247], true
	case 0xa8:
		return termRGB[248], true
	case 0xb2:
		return termRGB[249], true
	case 0xbc:
		return termRGB[250], true
	case 0xc6:
		return termRGB[251], true
	case 0xd0:
		return termRGB[252], true
	case 0xda:
		return termRGB[253], true
	case 0xe4:
		return termRGB[254], true
	case 0xee:
		return termRGB[255], true
	}
	return nil, false
}
