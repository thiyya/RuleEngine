package ScoreCalculation

import (
	"RuleEngine/ScoreCalculationMainConcept"
	"RuleEngine/utils"
)

type ScoreCalculationType interface {
	ScoreCalculationSteps(sessionInfo map[string]interface{}) []utils.ScoringCalculationStep
	BuildingModelTree(sessionInfo map[string]interface{})
	Validation(sessionInfo map[string]interface{})
	SetCriterionValueByFormula(sessionInfo map[string]interface{})
	ScaleCriterionScore(sessionInfo map[string]interface{})
	ExecuteBeforeCalculationInteractions(sessionInfo map[string]interface{})
	CalculateDemographicalScore(sessionInfo map[string]interface{}) float64
	ExecuteInteractions(sessionInfo map[string]interface{})
	ScaleDemographicalScoreToSegmentScore(sessionInfo map[string]interface{}) string
	ExecuteRules(sessionInfo map[string]interface{})
	CreateModelAndInteractionTrees(sessionInfo map[string]interface{})
}

func CalculateScore(scoreCalculationType ScoreCalculationType) (map[string]interface{}, error) {
	sessionInfo := map[string]interface{}{}
	steps := scoreCalculationType.ScoreCalculationSteps(sessionInfo)
	for _, step := range steps {
		if _, ok := sessionInfo["ErrorList"]; ok {
			return map[string]interface{}{}, sessionInfo["ErrorList"].(error)
		}
		switch step {
		case utils.Step_BuildingModelTree:
			scoreCalculationType.BuildingModelTree(sessionInfo)
			break
		case utils.Step_Validation:
			scoreCalculationType.Validation(sessionInfo)
			break
		case utils.Step_SetCriterionValueByFormula:
			scoreCalculationType.SetCriterionValueByFormula(sessionInfo)
			break
		case utils.Step_ExecuteBeforeCalculationInteractions:
			scoreCalculationType.ExecuteBeforeCalculationInteractions(sessionInfo)
			break
		case utils.Step_ScaleCriterionScore:
			scoreCalculationType.ScaleCriterionScore(sessionInfo)
			break
		case utils.Step_CalculateDemographicalScore:
			scoreCalculationType.CalculateDemographicalScore(sessionInfo)
			break
		case utils.Step_ExecuteInteractions:
			model := sessionInfo["Model"].(*ScoreCalculationMainConcept.Model)
			model.CreateCriterionScoreLogInfo("1")
			scoreCalculationType.ExecuteInteractions(sessionInfo)
			break
		case utils.Step_ScaleDemographicalScoreToSegmentScore:
			scoreCalculationType.ScaleDemographicalScoreToSegmentScore(sessionInfo)
			break
		case utils.Step_ExecuteRules:
			scoreCalculationType.ExecuteRules(sessionInfo)
			break
		case utils.Step_CreateModelAndInteractionTrees:
			scoreCalculationType.CreateModelAndInteractionTrees(sessionInfo)
			break
		default:
			break
		}
	}
	return sessionInfo, nil
}
