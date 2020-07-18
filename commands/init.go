package commands

func init() {
	InstructionReverseMap = make(map[string]uint8)
	for k, v := range InstructionForwardMap {
		InstructionReverseMap[v] = k
	}
}
