package ScoreCalculationModules

import "engine/utils"

type MainScoringPointOfView struct {
	ScoringCalculationHashMap map[string]interface{}
	ModelStructureInfo        []utils.ModelStructure
	ScaleInfo                 []utils.ScaleStructure
	InteractionInfo           []utils.InteractionStructure
	SegmentationInfo          []utils.SegmentationStructure
	ModelAdditionalInfo       utils.ModelAdditionalInfoStructure
}
