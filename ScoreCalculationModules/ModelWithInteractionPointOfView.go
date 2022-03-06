package ScoreCalculationModules

import "engine/utils"

type ModelWithInteractionPointOfView struct {
	ScoringCalculationHashMap map[string]interface{}
	ModelStructureInfo        []utils.ModelStructure
	InteractionInfo           []utils.InteractionStructure
	ModelAdditionalInfo       utils.ModelAdditionalInfoStructure
	RoundDecimalPlace         int
}
