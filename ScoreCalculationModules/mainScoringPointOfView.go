package ScoreCalculationModules

import (
	"RuleEngine/ScoreCalculationMainConcept"
	"RuleEngine/utils"
	"errors"
	"fmt"
)

type MainScoringPointOfView struct {
	ScoringCalculationHashMap map[string]interface{}
	ModelStructureInfo        []utils.ModelStructure
	ScaleInfo                 []utils.ScaleStructure
	InteractionInfo           []utils.InteractionStructure
	SegmentationInfo          []utils.SegmentationStructure
	ModelAdditionalInfo       utils.ModelAdditionalInfoStructure
}

func (mainScoringPointOfView MainScoringPointOfView) ScoreCalculationSteps(sessionInfo map[string]interface{}) []utils.ScoringCalculationStep {
	defer func() {
		if capturedError := recover(); capturedError != nil {
			err := errors.New(fmt.Sprintf("Score calculation step error : \n %s", capturedError.(error).Error()))
			sessionInfo["ErrorList"] = err
		}
	}()

	steps := []utils.ScoringCalculationStep{}

	steps = append(steps, utils.Step_BuildingModelTree)
	steps = append(steps, utils.Step_SetCriterionValueByFormula)
	steps = append(steps, utils.Step_ScaleCriterionScore)
	steps = append(steps, utils.Step_ExecuteBeforeCalculationInteractions)
	steps = append(steps, utils.Step_CalculateDemographicalScore)
	steps = append(steps, utils.Step_ExecuteInteractions)
	steps = append(steps, utils.Step_CalculateDemographicalScore)
	steps = append(steps, utils.Step_ScaleDemographicalScoreToSegmentScore)
	steps = append(steps, utils.Step_CreateModelAndInteractionTrees)
	sessionInfo["ScoringCalculationSteps"] = &steps
	return steps
}

func (mainScoringPointOfView MainScoringPointOfView) BuildingModelTree(sessionInfo map[string]interface{}) {
	defer func() {
		if capturedError := recover(); capturedError != nil {
			err := errors.New(fmt.Sprintf("Building. model tree step error : \n %s", capturedError.(error).Error()))
			sessionInfo["ErrorList"] = err
		}
	}()

	model := &ScoreCalculationMainConcept.Model{ModelId: mainScoringPointOfView.ModelAdditionalInfo.ModelId}

	model.MaxCriterionScoreOfModel = mainScoringPointOfView.ModelAdditionalInfo.MaxCriterionScoreOfModel
	model.MinCriterionScoreOfModel = mainScoringPointOfView.ModelAdditionalInfo.MinCriterionScoreOfModel

	model.MaxModelSensitiveValue = mainScoringPointOfView.ModelAdditionalInfo.MaxModelSensitiveValue
	model.MinModelSensitiveValue = mainScoringPointOfView.ModelAdditionalInfo.MinModelSensitiveValue

	model.ModelWeightDistributionTypeId = mainScoringPointOfView.ModelAdditionalInfo.ModelWeightDistributionTypeId

	model.FindCriterionBelongToModel(mainScoringPointOfView.ModelStructureInfo)
	model.SetScoringCalculationHashMap(mainScoringPointOfView.ScoringCalculationHashMap)
	model.SetCriterionValuesBelongToModel()

	sessionInfo["Model"] = model
}

func (mainScoringPointOfView MainScoringPointOfView) SetCriterionValueByFormula(sessionInfo map[string]interface{}) {
	defer func() {
		if capturedError := recover(); capturedError != nil {
			err := errors.New(fmt.Sprintf("Set Criterion Value By Formula step error : \n %s", capturedError.(error).Error()))
			sessionInfo["ErrorList"] = err
		}
	}()
	model := sessionInfo["Model"].(*ScoreCalculationMainConcept.Model)
	criterionWhoseFormulaCouldntExecuted, errList := model.FindCriterionValuesByFormula(model.FindModelCriterions())
	for i := 0; i < 10; i++ {
		if len(criterionWhoseFormulaCouldntExecuted) > 0 {
			criterionWhoseFormulaCouldntExecuted, errList = model.FindCriterionValuesByFormula(criterionWhoseFormulaCouldntExecuted)
		} else {
			errList = []error{}
			break
		}
	}
	if len(errList) > 0 {
		errString := ""
		for _, err := range errList {
			errString += err.Error() + "\n"
		}
		err := errors.New(fmt.Sprintf("Set Criterion Value By Formula step error : \n %s", errString))
		sessionInfo["ErrorList"] = err
	} else {
		_, err := model.ControlWhetherCriterionValueIsFound()
		if err != nil {
			sessionInfo["ErrorList"] = errors.New(fmt.Sprintf("Set Criterion Value By Formula step error : \n %s", err.Error()))
		}
	}
}

func (mainScoringPointOfView MainScoringPointOfView) ScaleCriterionScore(sessionInfo map[string]interface{}) {
	defer func() {
		if capturedError := recover(); capturedError != nil {
			err := errors.New(fmt.Sprintf("Scale Criterion Score step error : \n %s", capturedError.(error).Error()))
			sessionInfo["ErrorList"] = err
		}
	}()
	model := sessionInfo["Model"].(*ScoreCalculationMainConcept.Model)
	err := model.CalculateLeafScores(mainScoringPointOfView.ScaleInfo)
	if utils.IsAnErrorOccured(err) {
		sessionInfo["ErrorList"] = errors.New(fmt.Sprintf("Scale Criterion Score step error : \n %s", err.Error()))
	}

}

func (mainScoringPointOfView MainScoringPointOfView) ExecuteBeforeCalculationInteractions(sessionInfo map[string]interface{}) {
	defer func() {
		if capturedError := recover(); capturedError != nil {
			err := errors.New(fmt.Sprintf("Execute Before Calculation Interactions step error : \n %s", capturedError.(error).Error()))
			sessionInfo["ErrorList"] = err
		}
	}()
	model := sessionInfo["Model"].(*ScoreCalculationMainConcept.Model)
	beforeCalculationInteractions := []utils.InteractionStructure{}
	filteredInteractions := utils.Filter(mainScoringPointOfView.InteractionInfo, func(val interface{}) bool {
		return val.(utils.InteractionStructure).InteractionType == int(utils.Interaction_BeforeCalculation)
	}).([]interface{})
	if len(filteredInteractions) > 0 {
		for _, filteredInteraction := range filteredInteractions {
			beforeCalculationInteractions = append(beforeCalculationInteractions, filteredInteraction.(utils.InteractionStructure))
		}
		model.CreateInteractions(beforeCalculationInteractions)
		err := model.ExecuteInteractions()
		if err != nil {
			sessionInfo["ErrorList"] = errors.New(fmt.Sprintf("Execute Before Calculation Interactions step error : \n %s", err.Error()))
		}
	}
}

func (mainScoringPointOfView MainScoringPointOfView) CalculateDemographicalScore(sessionInfo map[string]interface{}) float64 {
	defer func() {
		if capturedError := recover(); capturedError != nil {
			err := errors.New(fmt.Sprintf("Calculate Demographical Score step error : \n %s", capturedError.(error).Error()))
			sessionInfo["ErrorList"] = err
		}
	}()
	model := sessionInfo["Model"].(*ScoreCalculationMainConcept.Model)
	modelCriterions := model.FindModelCriterions()
	for _, criterion := range modelCriterions {
		if criterion.CriterionFormula.Formula == "NOTCALCULABLE" && criterion.CriterionScore == -88888 {
			sessionInfo["ErrorList"] = errors.New(fmt.Sprintf("Calculate Demographical Score step error : \n %s", criterion.CriterionName+" criteria's score could not be calculated."))
			return 0
		}
	}

	model.ModelWeight = mainScoringPointOfView.ModelAdditionalInfo.ModelWeight
	model.DistributeModelWeight()
	model.CalculateScore()
	demographicalScore := model.CalculateModelScore()
	sessionInfo["DemographicalScore"] = demographicalScore
	return demographicalScore
}

func (mainScoringPointOfView MainScoringPointOfView) CreateModelAndInteractionTrees(sessionInfo map[string]interface{}) {
	defer func() {
		if capturedError := recover(); capturedError != nil {
			err := errors.New(fmt.Sprintf("Create Model And Interaction Trees step error : \n %s", capturedError.(error).Error()))
			sessionInfo["ErrorList"] = err
		}
	}()
	model := sessionInfo["Model"].(*ScoreCalculationMainConcept.Model)
	model.CreateModelTreeDataStructure()
	model.CreateInteractionTreeDataStructure()

}

func (mainScoringPointOfView MainScoringPointOfView) ExecuteInteractions(sessionInfo map[string]interface{}) {
	defer func() {
		if capturedError := recover(); capturedError != nil {
			err := errors.New(fmt.Sprintf("Execute Interactions step error : \n %s", capturedError.(error).Error()))
			sessionInfo["ErrorList"] = err
		}
	}()
	model := sessionInfo["Model"].(*ScoreCalculationMainConcept.Model)
	interactions := []utils.InteractionStructure{}
	filteredInteractions := utils.Filter(mainScoringPointOfView.InteractionInfo, func(val interface{}) bool {
		return val.(utils.InteractionStructure).InteractionType != int(utils.Interaction_BeforeCalculation)
	}).([]interface{})
	if len(filteredInteractions) > 0 {
		for _, filteredInteraction := range filteredInteractions {
			interactions = append(interactions, filteredInteraction.(utils.InteractionStructure))
		}
		model.CreateInteractions(interactions)
		err := model.ExecuteInteractions()
		for _, criterion := range model.FindModelCriterions() {
			criterion.CriterionScore = criterion.CriterionNormalizedScore
		}
		if err != nil {
			sessionInfo["ErrorList"] = errors.New(fmt.Sprintf("Execute Interactions step error : \n %s", err.Error()))
		}
	}
}

func (mainScoringPointOfView MainScoringPointOfView) ScaleDemographicalScoreToSegmentScore(sessionInfo map[string]interface{}) string {
	defer func() {
		if capturedError := recover(); capturedError != nil {
			err := errors.New(fmt.Sprintf("Scale Demographical Score To SegmentScore step error : \n %s", capturedError.(error).Error()))
			sessionInfo["ErrorList"] = err
		}
	}()
	demographicalScore := sessionInfo["DemographicalScore"].(float64)
	scoreSegmentation := ""
	for _, segmentationInfo := range mainScoringPointOfView.SegmentationInfo {
		if mainScoringPointOfView.ModelAdditionalInfo.SegmentationLimit == "A" {
			if demographicalScore >= segmentationInfo.MinimumPoint && demographicalScore < segmentationInfo.MaximumPoint {
				scoreSegmentation = segmentationInfo.ScoreSegmentation
				break
			}
		} else {
			if demographicalScore > segmentationInfo.MinimumPoint && demographicalScore <= segmentationInfo.MaximumPoint {
				scoreSegmentation = segmentationInfo.ScoreSegmentation
				break
			}
		}
	}
	sessionInfo["SegmentScore"] = scoreSegmentation
	return scoreSegmentation
}

func (mainScoringPointOfView MainScoringPointOfView) Validation(sessionInfo map[string]interface{}) {
	defer func() {
		if capturedError := recover(); capturedError != nil {
			err := errors.New(fmt.Sprintf("Validation step error : \n %s", capturedError.(error).Error()))
			sessionInfo["ErrorList"] = err
		}
	}()
	sessionInfo["ErrorList"] = errors.New(fmt.Sprintf("Validation step error : \n %s", "Not implemented."))
}

func (mainScoringPointOfView MainScoringPointOfView) ExecuteRules(sessionInfo map[string]interface{}) {
	defer func() {
		if capturedError := recover(); capturedError != nil {
			err := errors.New(fmt.Sprintf("Execute Rules step error : \n %s", capturedError.(error).Error()))
			sessionInfo["ErrorList"] = err
		}
	}()
	sessionInfo["ErrorList"] = errors.New(fmt.Sprintf("Execute Rules step error : \n %s", "Not implemented."))

}
