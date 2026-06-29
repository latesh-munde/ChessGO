package ai

// this defines bot strength

type BotConfig struct {
	Rating             int
	SearchDepth        int
	ThinkTimeMS        float64
	EvaluationNoise    float64
	MistakeProbability float64
}
