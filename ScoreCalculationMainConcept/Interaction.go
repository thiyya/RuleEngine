package ScoreCalculationMainConcept

import (
	"engine/utils"
	"errors"
	"fmt"
	"strconv"
)

type InteractionLogStructure struct {
	InteractionId          int
	InteractionFormulasLog []*PartFormulaLogStructure
	EffectedCriterion      *Criterion
	InteractionBehaviour   utils.InteractionBehaviour
	InteractionType        utils.InteractionType
	InteractionOrder       int
	IsInteractionPerformed bool
}

type Interaction struct {
	InteractionId        int
	InteractionFormulas  []*PartFormula
	EffectedCriterion    *Criterion
	InteractionBehaviour utils.InteractionBehaviour
	InteractionType      utils.InteractionType
	InteractionResult    string
	InteractionOrder     int
	InteractionModel     *Model
	InteractionLog       *InteractionLogStructure
}

func CreateInteraction(model *Model, inteactionId int, effectedCriterion *Criterion, interactionBehaviour utils.InteractionBehaviour, interactionType utils.InteractionType, interactionOrder int) *Interaction {
	tempInteractionLog := InteractionLogStructure{
		InteractionId:          inteactionId,
		EffectedCriterion:      effectedCriterion,
		InteractionBehaviour:   interactionBehaviour,
		InteractionType:        interactionType,
		InteractionOrder:       interactionOrder,
		IsInteractionPerformed: false,
	}
	tempInteraction := Interaction{
		InteractionId:        inteactionId,
		EffectedCriterion:    effectedCriterion,
		InteractionBehaviour: interactionBehaviour,
		InteractionType:      interactionType,
		InteractionOrder:     interactionOrder,
		InteractionModel:     model,
		InteractionLog:       &tempInteractionLog,
	}
	return &tempInteraction

}

func (interaction *Interaction) SetInteractionFormulas(interactionFormulas []*PartFormula) {
	interaction.InteractionFormulas = interactionFormulas
	var partFormulaLog []*PartFormulaLogStructure
	for _, formula := range interactionFormulas {
		partFormulaLog = append(partFormulaLog, CreatePartFormulaLogInfo(formula.FormulaOrderNum, formula.Formula, formula.FormulaResult))
	}
	interaction.InteractionLog.InteractionFormulasLog = partFormulaLog
}

func (interaction *Interaction) ExecuteFormula() error {
	for _, partFormula := range interaction.InteractionFormulas {
		result, finalStateOfFormula, variablesOfFormula, err := partFormula.ExecutePartFormula()
		partFormulaLogInfo := utils.Filter(interaction.InteractionLog.InteractionFormulasLog, func(val interface{}) bool {
			return val.(*PartFormulaLogStructure).FormulaOrderNum == partFormula.FormulaOrderNum
		}).([]interface{})[0].(*PartFormulaLogStructure)
		partFormulaLogInfo.FormulaLastState = finalStateOfFormula
		partFormulaLogInfo.VariablesOfFormula = variablesOfFormula
		if !utils.IsAnErrorOccured(err) {
			partFormulaLogInfo.IsFormulaExecuted = true
			if result.(bool) {
				interaction.InteractionResult = partFormula.FormulaResult
				return nil
			}
		} else {
			return err
		}
	}
	return nil
}

func (interaction *Interaction) ReflectTheBehaviour() error {
	if interaction.InteractionResult == "" {
		return nil
	}
	switch interaction.InteractionBehaviour {
	case utils.Behaviour_IncreaseCriterionScore:
		{
			newScore, err := strconv.ParseFloat(interaction.InteractionResult, 64)
			if err != nil {
				interaction.EffectedCriterion.CriterionScore += newScore
			} else {
				return errors.New(fmt.Sprintf("Interaction Result can not be parsed. %s", err.Error()))
			}
		}
	case utils.Behaviour_DecreaseCriterionScore:
		{
			newScore, err := strconv.ParseFloat(interaction.InteractionResult, 64)
			if err != nil {
				interaction.EffectedCriterion.CriterionScore -= newScore
			} else {
				return errors.New(fmt.Sprintf("Interaction Result can not be parsed. %s", err.Error()))
			}
		}
	case utils.Behaviour_OverrideCriterionScore:
		{
			newScore, err := strconv.ParseFloat(interaction.InteractionResult, 64)
			if err != nil {
				interaction.EffectedCriterion.CriterionScore = newScore
			} else {
				return errors.New(fmt.Sprintf("Interaction Result can not be parsed. %s", err.Error()))
			}
		}
	case utils.Behaviour_OverrideCriterionValue:
		{
			interaction.EffectedCriterion.CriterionValue = interaction.InteractionResult
			if interaction.InteractionResult == "-99999" {
				interaction.EffectedCriterion.CriterionWeight = 0
				interaction.EffectedCriterion.CriterionScore = 0
				break
			}
			err := interaction.EffectedCriterion.CalculateCriterionScore()
			if utils.IsAnErrorOccured(err) {
				return err
			}
			break
		}
	case utils.Behaviour_ChangeSensitivityValueOfModel:
		{
			interaction.EffectedCriterion.CriterionValue = interaction.InteractionResult
			if interaction.InteractionResult == "-99999" {
				interaction.EffectedCriterion.CriterionWeight = 0
				interaction.EffectedCriterion.CriterionScore = 0
				break
			}
			err := interaction.EffectedCriterion.CalculateCriterionScore()
			if utils.IsAnErrorOccured(err) {
				return err
			}
			break
		}
	default:
		break
	}
	if interaction.EffectedCriterion != nil {
		if interaction.InteractionModel.MaxCriterionScoreOfModel != 0 && interaction.InteractionModel.MinCriterionScoreOfModel != 0 {
			if interaction.EffectedCriterion.CriterionScore > interaction.InteractionModel.MaxCriterionScoreOfModel {
				interaction.EffectedCriterion.CriterionNormalizedScore = interaction.InteractionModel.MaxCriterionScoreOfModel
			} else if interaction.EffectedCriterion.CriterionScore < interaction.InteractionModel.MinCriterionScoreOfModel {
				interaction.EffectedCriterion.CriterionNormalizedScore = interaction.InteractionModel.MinCriterionScoreOfModel
			}
		} else {
			interaction.EffectedCriterion.CriterionNormalizedScore = interaction.EffectedCriterion.CriterionScore
		}
	}
	interaction.InteractionLog.IsInteractionPerformed = true
	return nil
}

type InteractionList []*Interaction

func (s InteractionList) Len() int {
	return len(s)
}
func (s InteractionList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s InteractionList) Less(i, j int) bool {
	return s[i].InteractionOrder < s[j].InteractionOrder
}
