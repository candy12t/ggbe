package ppu

// PPU mode
type mode int

const (
	HBlank  mode = iota // period of waiting HSYNC signal. period depends on sprite to be rendered etc.
	VBlank              // period of waiting VSYNC signal.
	OAMScan             // period of scanning OAM. period is 80 T-cycles.
	Drawing             // period of drawing pixel data and output to LCD. period depends on sprite to be rendered etc.
)

// each bit of LCDC register
const (
	BGWindowEnable         = 1 << 0
	SpriteEnable           = 1 << 1
	SpriteSize             = 1 << 2
	BGTileMap              = 1 << 3
	TileDataAddressingMode = 1 << 4
	WindowEnable           = 1 << 5
	WindowTileMap          = 1 << 6
	PPUEnable              = 1 << 7
)

// each bit of STAT register
const (
	LYC_EQ_LY     = 1 << 2
	HBlankInt     = 1 << 3
	VBlankInt     = 1 << 4
	OAMScanInt    = 1 << 5
	LYC_EQ_LY_INT = 1 << 6
)

// PPU is Picture Processing Unit.
// PPU perform rendering proccess to outout graphics to LCD display.
// PPU render `bg`, `window` and `sprite`.
type PPU struct {
	mode mode
	lcdc uint8         // 0xFF40, LCDC register
	stat uint8         // 0xFF41, STAT register
	scy  uint8         // 0xFF42, SCY register specify how much to scroll `bg`.
	scx  uint8         // 0xFF43, SCX register specify how much to scroll `bg`.
	ly   uint8         // 0xFF44, LY register hold which line of the LCD is currently being renderd.
	lyc  uint8         // 0xFF45, LYC register
	bgp  uint8         // 0xFF47, BGP register hold pallet data used when rendering `bg` and `window`.
	obp0 uint8         // 0xFF48, OBP0 register hold pallet data used when rendering `sprite`.
	obp1 uint8         // 0xFF49, OBP1 register hold pallet data used when rendering `sprite`.
	wy   uint8         // 0xFF4A, WY register specify the upper left position of the window.
	wx   uint8         // 0xFF4B, WX register specify the upper left position of the window.
	vram [0x2000]uint8 // 0x8000 ~ 0x9FFF, VRAM is Video RAM. the size is 8 KiB.
	oam  [0xA0]uint8   // 0xFE00 ~ 0xFE9F, OAM is Object Attribute Memory. the size is 160 B.
}

func New() *PPU {
	return &PPU{
		mode: OAMScan,
	}
}

func (p *PPU) Read(addr uint16) uint8 {
	switch {
	case 0x8000 <= addr && addr <= 0x9FFF:
		if p.mode == Drawing {
			return 0xFF
		}
		return p.vram[addr&0x1FFF]
	case 0xFE00 <= addr && addr <= 0xFE9F:
		if p.mode == Drawing || p.mode == OAMScan {
			return 0xFF
		}
		return p.oam[addr&0xFF]
	default:
		switch addr {
		case 0xFF40:
			return p.lcdc
		case 0xFF41:
			return 0x80 | p.stat | uint8(p.mode) // 7 bit is always 1.
		case 0xFF42:
			return p.scy
		case 0xFF43:
			return p.scx
		case 0xFF44:
			return p.ly
		case 0xFF45:
			return p.lyc
		case 0xFF47:
			return p.bgp
		case 0xFF48:
			return p.obp0
		case 0xFF49:
			return p.obp1
		case 0xFF4A:
			return p.wy
		case 0xFF4B:
			return p.wx
		default:
			panic("unreachable")
		}
	}
}

func (p *PPU) Write(addr uint16, val uint8) {
	switch {
	case 0x8000 <= addr && addr <= 0x9FFF:
		if p.mode != Drawing {
			p.vram[addr&0x1FFF] = val
		}
	case 0xFE00 <= addr && addr <= 0xFE9F:
		if p.mode != Drawing && p.mode != OAMScan {
			p.oam[addr&0xFF] = val
		}
	default:
		switch addr {
		case 0xFF40:
			p.lcdc = val
		case 0xFF41:
			p.stat = (p.stat & LYC_EQ_LY) | (val & 0xF8) // 0 ~ 2 bit is not writable.
		case 0xFF42:
			p.scy = val
		case 0xFF43:
			p.scx = val
		case 0xFF44: // LY register is not writable.
		case 0xFF45:
			p.lyc = val
		case 0xFF47:
			p.bgp = val
		case 0xFF48:
			p.obp0 = val
		case 0xFF49:
			p.obp1 = val
		case 0xFF4A:
			p.wy = val
		case 0xFF4B:
			p.wx = val
		default:
			panic("unreachable")
		}
	}
}
