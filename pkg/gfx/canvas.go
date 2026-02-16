package gfx

import (
	"fmt"
	"math"
	"strings"

	. "github.com/novelalex/soft-raytracer/pkg/nmath"
	"github.com/novelalex/soft-raytracer/pkg/nutil"
)

type Canvas struct {
	width  uint
	height uint
	buffer []Color
}

func NewCanvas(width, height uint) Canvas {
	buffer_size := width * height
	buffer := make([]Color, buffer_size)
	return Canvas{width, height, buffer}
}

func (c Canvas) Width() uint {
	return c.width
}

func (c Canvas) Height() uint {
	return c.height
}

func (c Canvas) PixelAt(x, y uint) Color {
	return c.buffer[y*c.width+x]
}

func (c *Canvas) WritePixel(x, y uint, color Color) {
	c.buffer[y*c.width+x] = color
}

func (c Canvas) constructPPMHeader() string {
	return fmt.Sprintf("P3\n%d %d\n255\n", c.width, c.height)
}

func (c Canvas) constructP6PPMHeader() []byte {
	return []byte(fmt.Sprintf("P6\n%d %d\n255\n", c.width, c.height))
}

func (c Canvas) constructP6PPMBody() []byte {
	result := []byte{}
	for _, color := range c.buffer {
		for i := range 3 {
			cc := uint8(math.Max(0, math.Min(math.Round(color.At(i)*255), 255)))
			result = append(result, byte(cc))
		}
	}

	return result
}

func (c Canvas) AsP6PPM() []byte {
	result := []byte{}
	result = append(result, c.constructP6PPMHeader()...)
	result = append(result, c.constructP6PPMBody()...)
	return result
}

func (c Canvas) constructPPMBody() string {
	var sb strings.Builder
	width_counter := 0
	ln_start := true

	for _, color := range c.buffer {
		for i := range 3 {
			cc := uint8(math.Max(0, math.Min(math.Round(color.At(i)*255), 255)))

			written_digits := fmt.Sprintf("%d", cc)
			width_needed := len(written_digits) + nutil.IntFromBool(!ln_start)

			if width_counter+width_needed > 70 {
				sb.WriteString("\n")
				width_counter = 0
				ln_start = true
			}

			if !ln_start {
				sb.WriteString(" ")
			}

			sb.WriteString(written_digits)
			width_counter += width_needed
			ln_start = false
		}
	}
	sb.WriteString("\n\n")

	return sb.String()
}

func (c Canvas) AsPPM() string {
	ppm_header := c.constructPPMHeader()
	ppm_body := c.constructPPMBody()
	return ppm_header + ppm_body
}
