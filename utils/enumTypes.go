package utils

type InteractionBehaviour int

const (
	Behaviour_None InteractionBehaviour = iota
	Behaviour_IncreaseCriterionScore
	Behaviour_DecreaseCriterionScore
	Behaviour_OverrideCriterionScore
	Behaviour_OverrideCriterionValue
	Behaviour_ChangeSensitivityValueOfModel
)

type ScoringCalculationStep int

const (
	Step_BuildingModelTree ScoringCalculationStep = iota
	Step_Validation
	Step_SetCriterionValueByFormula
	Step_ScaleCriterionScore
	Step_CalculateDemographicalScore
	Step_ExecuteInteractions
	Step_ExecuteBeforeCalculationInteractions
	Step_ScaleDemographicalScoreToSegmentScore
	Step_ExecuteRules
	Step_CreateModelAndInteractionTrees
)

type InteractionType int

const (
	Interaction_Rota InteractionType = iota + 1
	Interaction_EtkilesimliRota
	Interaction_Hassas
	Interaction_GenelEtkilesim
	Interaction_YoneticiEtkisi
	Interaction_MizanEtkisi
	Interaction_MaliKayitKalitesi
	Interaction_RiskAnalizi
	Interaction_BeforeCalculation
)

type DistributeModelWeightType int

const (
	DistributeType_FDS DistributeModelWeightType = iota
	DistributeType_General
	DistributeType_DontDistribute
)

type SegmentationType int

const (
	SegmentationType_10 SegmentationType = iota + 1
	SegmentationType_5
	SegmentationType_Renk
)

type ModelStructure struct {
	Name                string
	CriterionId         string
	Structure           string
	Weight              float64
	CriterionScaleLimit string
	CriterionFormula    FormulaStructure
}

type ScaleStructure struct {
	CriterionId       string
	FirstValueOfScale string
	LastValueOfScale  string
	LowerScore        float64
	HigherScore       float64
	Explanation       string
}

type InteractionStructure struct {
	InteractionId            int
	PriorityOrderNum         int
	InteractionType          int
	EffectedCriterionId      string
	InteractionBehaviourType int
	FormulaList              []FormulaStructure
}

type FormulaStructure struct {
	FormulaPriorityNum int
	Formula            string
	Result             string
}

type ParentChildRelationship struct {
	Parent string
	Child  string
}

type SegmentationStructure struct {
	ScoreSegmentation string
	MinimumPoint      float64
	MaximumPoint      float64
}

type ModelAdditionalInfoStructure struct {
	ModelId                       int
	SegmentationLimit             string
	ModelWeight                   float64
	MaxCriterionScoreOfModel      float64
	MinCriterionScoreOfModel      float64
	MaxModelSensitiveValue        float64
	MinModelSensitiveValue        float64
	ModelWeightDistributionTypeId DistributeModelWeightType
	SegmentationType              int
	ModelName                     string
}

type RelationType int

const (
	RelationType_Scale                  RelationType = 9
	RelationType_Rate                                = 10
	RelationType_Interaction                         = 11
	RelationType_Option                              = 13
	RelationType_Statement                           = 14
	RelationType_Option_Scale_Statement              = 51
)

type ModelAdditionalInfoResponse struct {
	ModelAdditionalInfoStructure ModelAdditionalInfoStructure
}
