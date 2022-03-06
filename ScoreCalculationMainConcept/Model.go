package ScoreCalculationMainConcept

import "engine/utils"

type Model struct {
	ModelId                       int
	ModelTree                     *Criterion
	ModelWeight                   float64
	ModelCriterions               map[string]*Criterion
	ModelSensitiveValue           float64
	ParentChildRelationship       []utils.ParentChildRelationship
	ModelTreeDataStructure        *utils.Node
	InteractionTreeDataStructure  *utils.Node
	InteractionList               []*Interaction
	ScoreCalculationHashMap       map[string]interface{}
	MaxCriterionScoreOfModel      float64
	MinCriterionScoreOfModel      float64
	MaxModelSensitiveValue        float64
	MinModelSensitiveValue        float64
	ModelWeightDistributionTypeId utils.DistributeModelWeightType
}
