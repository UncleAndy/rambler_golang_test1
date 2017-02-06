package bit_counter

const (
	TwoBitsMask = 0x5555555555555555
	FourBitsMask = 0x3333333333333333
	EightBitsMask = 0x0F0F0F0F0F0F0F0F
	SixteenBitsMask = 0x00FF00FF00FF00FF
	ThirtyTwoBitsMask = 0x0000FFFF0000FFFF
	FinalBitsMask = 0x00000000FFFFFFFF
)

func BitsCount(val uint64) uint8 {
	tmp := (val&TwoBitsMask) + ((val>>1)&TwoBitsMask)
	tmp = (tmp&FourBitsMask) + ((tmp>>2)&FourBitsMask)
	tmp = (tmp&EightBitsMask) + ((tmp>>4)&EightBitsMask)
	tmp = (tmp&SixteenBitsMask) + ((tmp>>8)&SixteenBitsMask)
	tmp = (tmp&ThirtyTwoBitsMask) + ((tmp>>16)&ThirtyTwoBitsMask)
	tmp = (tmp&FinalBitsMask) + ((tmp>>32)&FinalBitsMask)
	return uint8(tmp)
}
