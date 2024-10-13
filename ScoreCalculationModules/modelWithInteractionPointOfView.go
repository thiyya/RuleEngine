package ScoreCalculationModules

import (
	"RuleEngine/ScoreCalculationMainConcept"
	"RuleEngine/utils"
	"errors"
	"fmt"
)

type ModelWithInteractionPointOfView struct {
	ScoringCalculationHashMap map[string]interface{}
	ModelStructureInfo        []utils.ModelStructure
	InteractionInfo           []utils.InteractionStructure
	ModelAdditionalInfo       utils.ModelAdditionalInfoStructure
	RoundDecimalPlace         int
}

func (m ModelWithInteractionPointOfView) ScoreCalculationSteps(sessionInfo map[string]interface{}) []utils.ScoringCalculationStep {
	defer func() {
		if capturedError := recover(); capturedError != nil {
			err := errors.New(fmt.Sprintf("Score calculation step error : \n %s", capturedError.(error).Error()))
			sessionInfo["ErrorList"] = err
		}
	}()

	steps := []utils.ScoringCalculationStep{}

	steps = append(steps, utils.Step_BuildingModelTree)
	steps = append(steps, utils.Step_ExecuteBeforeCalculationInteractions)
	steps = append(steps, utils.Step_ExecuteInteractions)
	steps = append(steps, utils.Step_CalculateDemographicalScore)
	steps = append(steps, utils.Step_CreateModelAndInteractionTrees)
	sessionInfo["ScoringCalculationSteps"] = &steps
	return steps
}

func (m ModelWithInteractionPointOfView) BuildingModelTree(sessionInfo map[string]interface{}) {
	defer func() {
		if capturedError := recover(); capturedError != nil {
			err := errors.New(fmt.Sprintf("Building model tree step error : \n %s", capturedError.(error).Error()))
			sessionInfo["ErrorList"] = err
		}
	}()

	model := &ScoreCalculationMainConcept.Model{ModelId: m.ModelAdditionalInfo.ModelId}

	model.MaxCriterionScoreOfModel = m.ModelAdditionalInfo.MaxCriterionScoreOfModel
	model.MinCriterionScoreOfModel = m.ModelAdditionalInfo.MinCriterionScoreOfModel

	model.MaxModelSensitiveValue = m.ModelAdditionalInfo.MaxModelSensitiveValue
	model.MinModelSensitiveValue = m.ModelAdditionalInfo.MinModelSensitiveValue

	model.ModelWeightDistributionTypeId = m.ModelAdditionalInfo.ModelWeightDistributionTypeId

	model.FindCriterionBelongToModel(m.ModelStructureInfo)
	model.SetScoringCalculationHashMap(m.ScoringCalculationHashMap)
	model.SetCriterionValuesBelongToModel()

	sessionInfo["Model"] = model
}

func (m ModelWithInteractionPointOfView) SetCriterionValueByFormula(sessionInfo map[string]interface{}) {
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

func (m ModelWithInteractionPointOfView) ScaleCriterionScore(sessionInfo map[string]interface{}) {
	defer func() {
		if capturedError := recover(); capturedError != nil {
			err := errors.New(fmt.Sprintf("ScaleCriterionScore step error : \n %s", capturedError.(error).Error()))
			sessionInfo["ErrorList"] = err
		}
	}()
	sessionInfo["ErrorList"] = errors.New(fmt.Sprintf("ScaleCriterionScore step error : \n %s", "Not implemented."))
}

func (m ModelWithInteractionPointOfView) ExecuteBeforeCalculationInteractions(sessionInfo map[string]interface{}) {
	defer func() {
		if capturedError := recover(); capturedError != nil {
			err := errors.New(fmt.Sprintf("Execute Before Calculation Interactions step error : \n %s", capturedError.(error).Error()))
			sessionInfo["ErrorList"] = err
		}
	}()
	model := sessionInfo["Model"].(*ScoreCalculationMainConcept.Model)
	beforeCalculationInteractions := []utils.InteractionStructure{}
	filteredInteractions := utils.Filter(m.InteractionInfo, func(val interface{}) bool {
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

func (m ModelWithInteractionPointOfView) CalculateDemographicalScore(sessionInfo map[string]interface{}) float64 {
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

	model.ModelWeight = m.ModelAdditionalInfo.ModelWeight
	model.DistributeModelWeight()
	model.CalculateScore()
	demographicalScore := model.CalculateModelScore()
	sessionInfo["DemographicalScore"] = demographicalScore
	return demographicalScore
}

func (m ModelWithInteractionPointOfView) CreateModelAndInteractionTrees(sessionInfo map[string]interface{}) {
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

func (m ModelWithInteractionPointOfView) ExecuteInteractions(sessionInfo map[string]interface{}) {
	defer func() {
		if capturedError := recover(); capturedError != nil {
			err := errors.New(fmt.Sprintf("Execute Interactions step error : \n %s", capturedError.(error).Error()))
			sessionInfo["ErrorList"] = err
		}
	}()
	model := sessionInfo["Model"].(*ScoreCalculationMainConcept.Model)
	interactions := []utils.InteractionStructure{}
	filteredInteractions := utils.Filter(m.InteractionInfo, func(val interface{}) bool {
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

func (m ModelWithInteractionPointOfView) ScaleDemographicalScoreToSegmentScore(sessionInfo map[string]interface{}) string {
	defer func() {
		if capturedError := recover(); capturedError != nil {
			err := errors.New(fmt.Sprintf("ScaleDemographicalScoreToSegmentScore step error : \n %s", capturedError.(error).Error()))
			sessionInfo["ErrorList"] = err
		}
	}()
	sessionInfo["ErrorList"] = errors.New(fmt.Sprintf("ScaleDemographicalScoreToSegmentScore step error : \n %s", "Not implemented."))
	return ""
}

func (m ModelWithInteractionPointOfView) Validation(sessionInfo map[string]interface{}) {
	defer func() {
		if capturedError := recover(); capturedError != nil {
			err := errors.New(fmt.Sprintf("Validation step error : \n %s", capturedError.(error).Error()))
			sessionInfo["ErrorList"] = err
		}
	}()
	sessionInfo["ErrorList"] = errors.New(fmt.Sprintf("Validation step error : \n %s", "Not implemented."))
}

func (m ModelWithInteractionPointOfView) ExecuteRules(sessionInfo map[string]interface{}) {
	defer func() {
		if capturedError := recover(); capturedError != nil {
			err := errors.New(fmt.Sprintf("Execute Rules step error : \n %s", capturedError.(error).Error()))
			sessionInfo["ErrorList"] = err
		}
	}()
	sessionInfo["ErrorList"] = errors.New(fmt.Sprintf("Execute Rules step error : \n %s", "Not implemented."))

}
