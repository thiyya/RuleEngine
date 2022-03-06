package ScoreCalculationModules

import "RuleEngine/utils"

type ModelWithInteractionPointOfView struct {
	ScoringCalculationHashMap map[string]interface{}
	ModelStructureInfo        []utils.ModelStructure
	InteractionInfo           []utils.InteractionStructure
	ModelAdditionalInfo       utils.ModelAdditionalInfoStructure
	RoundDecimalPlace         int
}
