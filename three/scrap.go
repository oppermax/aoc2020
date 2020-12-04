if n == 4 {
	sLayers := []string{}
	for i, layer := range layers {
		if i%2 == 0 {
			sLayers = append(sLayers, layer)
		}
	}
	layers = sLayers
	log.Infof("For this round layers has only %v entries: %v to the right", len(layers), slopes[n])
}

// 94 186 281 368 410
// 94 99 214 91