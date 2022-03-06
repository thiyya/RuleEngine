package ScoreCalculationMainConcept

import (
	"RuleEngine/utils"
	"errors"
	"fmt"
	"github.com/leekchan/accounting"
	"sort"
	"strconv"
)

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
	ScoringCalculationHashMap     map[string]interface{}
	MaxCriterionScoreOfModel      float64
	MinCriterionScoreOfModel      float64
	MaxModelSensitiveValue        float64
	MinModelSensitiveValue        float64
	ModelWeightDistributionTypeId utils.DistributeModelWeightType
}

func (model *Model) FindModelCriterions() map[string]*Criterion {
	if model.ModelCriterions == nil || len(model.ModelCriterions) == 0 {
		model.ModelCriterions = map[string]*Criterion{}
		mainCriterion := []*Criterion{}
		mainCriterion = append(mainCriterion, model.ModelTree)
		model.addCriterionToModelCriterions(mainCriterion)
	}
	return model.ModelCriterions
}

func (model *Model) addCriterionToModelCriterions(criterionList []*Criterion) {
	for _, criterion := range criterionList {
		model.ModelCriterions[criterion.CriterionId] = criterion
		if !criterion.IsLeaf() {
			model.addCriterionToModelCriterions(criterion.ChildCriterion)
		}
	}
}

func (model *Model) CalculateModelScore() float64 {
	var normalizedSensitiveValue float64 = model.ModelSensitiveValue
	if model.MaxModelSensitiveValue != 0 && model.MinModelSensitiveValue != 0 {
		if model.ModelSensitiveValue > model.MaxModelSensitiveValue {
			normalizedSensitiveValue = model.MaxModelSensitiveValue
		}
		if model.ModelSensitiveValue < model.MinModelSensitiveValue {
			normalizedSensitiveValue = model.MinModelSensitiveValue
		}
	}

	modelScore := model.ModelTree.CriterionScore + normalizedSensitiveValue
	if model.MaxCriterionScoreOfModel != 0 && model.MinCriterionScoreOfModel != 0 {
		if modelScore > model.MaxCriterionScoreOfModel {
			modelScore = model.MaxCriterionScoreOfModel
		}
		if modelScore < model.MinCriterionScoreOfModel {
			modelScore = model.MinCriterionScoreOfModel
		}
	}

	return modelScore
}

func (model *Model) CalculateScore() {
	model.ModelTree.CriterionScore = model.calculateParentScore(model.ModelTree.ChildCriterion)
}

func (model *Model) calculateParentScore(childCriterionList []*Criterion) float64 {
	var childCriterionTotalScore float64 = 0
	var childCriterionTotalWeight float64 = 0
	for _, childCriterion := range childCriterionList {
		if childCriterion.IsLeaf() {
			childCriterionTotalScore += childCriterion.CriterionWeight * childCriterion.CriterionNormalizedScore
		} else {
			childCriterion.CriterionScore = model.calculateParentScore(childCriterion.ChildCriterion)
			childCriterion.CriterionNormalizedScore = childCriterion.CriterionScore
			childCriterionTotalScore += childCriterion.CriterionWeight * childCriterion.CriterionScore

		}
		childCriterionTotalWeight += childCriterion.CriterionWeight
	}

	if childCriterionTotalWeight == 0 {
		return 0
	}
	return childCriterionTotalScore / model.ModelWeight
}

func (model *Model) DistributeModelWeight() {
	if model.ModelWeightDistributionTypeId == utils.DistributeType_FDS {
		modelCriterions := model.FindModelCriterions()
		for _, criterion := range modelCriterions {
			if !criterion.IsLeaf() {
				if criterion.calculateTotalChildCriterionWeight() > 0 && criterion.calculateTotalChildCriterionWeight() != model.ModelWeight {
					totalChildCriterionsWeight := criterion.calculateTotalChildCriterionWeight()
					for _, childCriterion := range criterion.ChildCriterion {
						childCriterion.CriterionWeight = childCriterion.CriterionWeight * model.ModelWeight / totalChildCriterionsWeight
					}
				}
			}
		}
	} else if model.ModelWeightDistributionTypeId == utils.DistributeType_General {
		var totalLeafWeight float64 = 0
		modelCriterions := model.FindModelCriterions()
		for _, criterion := range modelCriterions {
			if criterion.IsLeaf() {
				totalLeafWeight += criterion.CriterionWeight
			} else {
				criterion.CriterionWeight = model.ModelWeight
			}
		}
		if totalLeafWeight != model.ModelWeight {
			for _, criterion := range modelCriterions {
				if criterion.IsLeaf() {
					criterion.CriterionWeight = criterion.CriterionWeight * model.ModelWeight / totalLeafWeight
				}
			}
		}
	}
}

func (model *Model) CalculateLeafScores(scaleInfo []utils.ScaleStructure) error {
	modelCriterions := model.FindModelCriterions()
	for _, criterion := range modelCriterions {
		if criterion.IsLeaf() && criterion.CriterionFormula.Formula != "NOTCALCULABLE" && criterion.CriterionValue != "-99999" {
			criterion.SetScaleInfo(scaleInfo)
			err := criterion.CalculateCriterionScore()
			if utils.IsAnErrorOccured(err) {
				return err
			}
		}
	}
	return nil
}

func (model *Model) SetScoringCalculationHashMap(scoringCalculationHashMap map[string]interface{}) {
	model.ScoringCalculationHashMap = scoringCalculationHashMap
}

func (model *Model) CreateModelTreeDataStructure() {
	rootCriterion := model.ModelTree
	var root = &utils.Node{rootCriterion.CriterionId, "", rootCriterion.CriterionName, []*utils.Node{}}
	var data []*utils.Node
	for _, parentChildRelation := range model.ParentChildRelationship {
		criterion := model.FindModelCriterions()[parentChildRelation.Child]
		criterionScoreBeforeInteraction := ""
		criterionValue := ""
		if len(criterion.CriterionScoreLogInfo) > 0 {
			f, err := strconv.ParseFloat(criterion.CriterionScoreLogInfo[0].CriterionScore, 64)
			if err == nil {
				criterionScoreBeforeInteraction = " Criterion Score Before Interaction : " + accounting.FormatNumber(f, 2, ".", ",")
			} else {
				criterionScoreBeforeInteraction = " Criterion Score Before Interaction : " + criterion.CriterionScoreLogInfo[0].CriterionScore
			}
		}
		f, err := strconv.ParseFloat(criterion.CriterionValue, 64)
		if err == nil {
			criterionValue = " Criterion Value : " + accounting.FormatNumber(f, 2, ".", ",")
		} else {
			criterionValue = " Criterion Value : " + criterion.CriterionValue
		}
		criterionWeight := " Criterion Weight : " + accounting.FormatNumber(criterion.calculateTotalChildCriterionWeight(), 2, ".", ",")
		criterionScore := " Criterion Score : " + accounting.FormatNumber(criterion.CriterionScore, 2, ".", ",")
		criterionNormalizedScore := " Criterion Normalized Score : " + accounting.FormatNumber(criterion.CriterionNormalizedScore, 2, ".", ",")
		criterionWeightedScore := " CriterionWeightedScore : " + accounting.FormatNumber(criterion.CalculateCriterionWeightedScore(model.ModelWeight), 2, ".", ",")
		var criterionInfo = criterion.CriterionId + " -> " + criterion.CriterionName + " " + criterionValue + " " + criterionWeight + " " + criterionScore + " " + criterionScoreBeforeInteraction + " " + criterionNormalizedScore + "	" + criterionNormalizedScore + " " + criterionWeightedScore
		data = append(data, &utils.Node{Id: parentChildRelation.Child, ParentId: parentChildRelation.Parent, CriterionInfo: criterion.CriterionName, ChildCriterion: []*utils.Node{}})

		data = append(data, &utils.Node{Id: parentChildRelation.Child + "Ack1", ParentId: parentChildRelation.Child, CriterionInfo: " Criterion Id : " + criterion.CriterionId, ChildCriterion: []*utils.Node{}})
		data = append(data, &utils.Node{Id: parentChildRelation.Child + "Ack2", ParentId: parentChildRelation.Child, CriterionInfo: criterionValue, ChildCriterion: []*utils.Node{}})
		data = append(data, &utils.Node{Id: parentChildRelation.Child + "Ack3", ParentId: parentChildRelation.Child, CriterionInfo: criterionWeight, ChildCriterion: []*utils.Node{}})
		data = append(data, &utils.Node{Id: parentChildRelation.Child + "Ack4", ParentId: parentChildRelation.Child, CriterionInfo: criterionScoreBeforeInteraction, ChildCriterion: []*utils.Node{}})
		data = append(data, &utils.Node{Id: parentChildRelation.Child + "Ack5", ParentId: parentChildRelation.Child, CriterionInfo: criterionNormalizedScore, ChildCriterion: []*utils.Node{}})
		data = append(data, &utils.Node{Id: parentChildRelation.Child + "Ack6", ParentId: parentChildRelation.Child, CriterionInfo: criterionWeightedScore, ChildCriterion: []*utils.Node{}})

		emptyFormula := utils.FormulaStructure{}
		if criterion.CriterionFormula != emptyFormula {
			criterionInfo += " Criterion Formula Log : -->   "
			criterionFormula := " Criterion Formula :  " + criterion.CriterionFormulaLogInfo.Formula
			isCriterionFormulaExecuted := " Criterion Formula is not executed ! " + criterionFormula
			criterionFormulaLastState := ""
			criterionFormulaResult := ""
			if criterion.CriterionFormulaLogInfo.IsFormulaExecuted {
				isCriterionFormulaExecuted = " Criterion Formula executed "
				criterionFormulaLastState = " Criterion Formula value : " + criterion.CriterionFormulaLogInfo.FormulaLastState
				f, err := strconv.ParseFloat(criterion.CriterionFormulaLogInfo.FormulaResult, 64)
				if err == nil {
					criterionFormulaResult = " Criterion Formula Result :  " + accounting.FormatNumber(f, 2, ".", ",")
				} else {
					criterionFormulaResult = " Criterion Formula Result :  " + criterion.CriterionFormulaLogInfo.FormulaResult
				}
			}
			criterionInfo += " " + isCriterionFormulaExecuted + " " + criterionFormula + " " + criterionFormulaLastState + " " + criterionFormulaResult
			data = append(data, &utils.Node{Id: parentChildRelation.Child + "Ack7", ParentId: parentChildRelation.Child, CriterionInfo: isCriterionFormulaExecuted, ChildCriterion: []*utils.Node{}})
			data = append(data, &utils.Node{Id: parentChildRelation.Child + "Ack8", ParentId: parentChildRelation.Child, CriterionInfo: criterionFormula, ChildCriterion: []*utils.Node{}})
			data = append(data, &utils.Node{Id: parentChildRelation.Child + "Ack9", ParentId: parentChildRelation.Child, CriterionInfo: criterionFormulaLastState, ChildCriterion: []*utils.Node{}})
			data = append(data, &utils.Node{Id: parentChildRelation.Child + "Ack10", ParentId: parentChildRelation.Child, CriterionInfo: criterionFormulaResult, ChildCriterion: []*utils.Node{}})
		}
	}
	root.AddModelTreeInfo(data...)
	model.ModelTreeDataStructure = root
}

func (model *Model) CreateInteractionTreeDataStructure() {
	var root = &utils.Node{"Interaction List", "", "Interaction List", []*utils.Node{}}
	var data []*utils.Node
	for _, interaction := range model.InteractionList {
		interactionId := " Interaction Id : " + strconv.Itoa(interaction.InteractionId)
		data = append(data, &utils.Node{Id: strconv.Itoa(interaction.InteractionId), ParentId: "Interaction List", CriterionInfo: interactionId, ChildCriterion: []*utils.Node{}})
	}

	for _, interaction := range model.InteractionList {
		var interactionFormulaLogData []*utils.Node
		for _, interactionFormulaLog := range interaction.InteractionLog.InteractionFormulasLog {
			partFormulaOrderNumber := " partFormulaOrderNumber : " + strconv.Itoa(interactionFormulaLog.FormulaOrderNum)
			partFormula := " partFormula : " + interactionFormulaLog.Formula
			partFormulaResult := " partFormulaResult : " + interactionFormulaLog.FormulaResult
			isPartFormulaExecuted := " Part Formula is not executed!"
			partFormulaLastState := ""
			if interactionFormulaLog.IsFormulaExecuted {
				isPartFormulaExecuted = " Part Formula executed."
				partFormulaLastState = " Part Formula Last State : " + interactionFormulaLog.FormulaLastState
			}
			interactionFormulaLogDataChild := []*utils.Node{}
			interactionFormulaLogDataChild = append(interactionFormulaLogDataChild, &utils.Node{Id: strconv.Itoa(interaction.InteractionId) + " _ " + partFormula, ParentId: "Interaction Formula Info _" + strconv.Itoa(interaction.InteractionId) + partFormulaOrderNumber, CriterionInfo: partFormula, ChildCriterion: []*utils.Node{}})
			interactionFormulaLogDataChild = append(interactionFormulaLogDataChild, &utils.Node{Id: strconv.Itoa(interaction.InteractionId) + " _ " + partFormulaResult, ParentId: "Interaction Formula Info _" + strconv.Itoa(interaction.InteractionId) + partFormulaOrderNumber, CriterionInfo: partFormulaResult, ChildCriterion: []*utils.Node{}})
			interactionFormulaLogDataChild = append(interactionFormulaLogDataChild, &utils.Node{Id: strconv.Itoa(interaction.InteractionId) + " _ " + isPartFormulaExecuted, ParentId: "Interaction Formula Info _" + strconv.Itoa(interaction.InteractionId) + partFormulaOrderNumber, CriterionInfo: isPartFormulaExecuted, ChildCriterion: []*utils.Node{}})
			interactionFormulaLogDataChild = append(interactionFormulaLogDataChild, &utils.Node{Id: strconv.Itoa(interaction.InteractionId) + " _ " + partFormulaLastState, ParentId: "Interaction Formula Info _" + strconv.Itoa(interaction.InteractionId) + partFormulaOrderNumber, CriterionInfo: partFormulaLastState, ChildCriterion: []*utils.Node{}})

			interactionFormulaLogData = append(interactionFormulaLogData, &utils.Node{Id: "Interaction Formula Info _ " + strconv.Itoa(interaction.InteractionId) + partFormulaOrderNumber, ParentId: "Interaction Formula Info _" + strconv.Itoa(interaction.InteractionId), CriterionInfo: partFormulaOrderNumber, ChildCriterion: interactionFormulaLogDataChild})

		}

		effectedCriterionId := ""
		if interaction.EffectedCriterion != nil {
			effectedCriterionId = " effectedCriterionId : " + interaction.EffectedCriterion.CriterionId
		}
		interactionBehaviour := " interactionBehaviour : "
		switch interaction.InteractionBehaviour {
		case utils.Behaviour_IncreaseCriterionScore:
			interactionBehaviour += "Behaviour_IncreaseCriterionScore "
		case utils.Behaviour_DecreaseCriterionScore:
			interactionBehaviour += "Behaviour_DecreaseCriterionScore "
		case utils.Behaviour_OverrideCriterionScore:
			interactionBehaviour += "Behaviour_OverrideCriterionScore "
		case utils.Behaviour_OverrideCriterionValue:
			interactionBehaviour += "Behaviour_OverrideCriterionValue "
		case utils.Behaviour_ChangeSensitivityValueOfModel:
			interactionBehaviour += "Behaviour_ChangeSensitivityValueOfModel "
		default:
			interactionBehaviour += "???"
		}
		interactionOrder := " interactionOrder : " + strconv.Itoa(interaction.InteractionOrder)
		interactionType := " interactionType : "
		switch interaction.InteractionType {
		case utils.Interaction_General:
			interactionType += "Interaction_General "
		case utils.Interaction_Sensitive:
			interactionType += "Interaction_Sensitive "
		case utils.Interaction_BeforeCalculation:
			interactionType += "Interaction_BeforeCalculation "
		default:
			interactionType += "???"
		}
		isInteractionApplied := " Interaction is not applied "
		if interaction.InteractionLog.IsInteractionPerformed {
			isInteractionApplied = " Interaction applied "
		}

		data = append(data, &utils.Node{Id: strconv.Itoa(interaction.InteractionId) + " _ " + interactionOrder, ParentId: strconv.Itoa(interaction.InteractionId), CriterionInfo: interactionOrder, ChildCriterion: []*utils.Node{}})
		data = append(data, &utils.Node{Id: strconv.Itoa(interaction.InteractionId) + " _ " + interactionBehaviour, ParentId: strconv.Itoa(interaction.InteractionId), CriterionInfo: interactionBehaviour, ChildCriterion: []*utils.Node{}})
		data = append(data, &utils.Node{Id: strconv.Itoa(interaction.InteractionId) + " _ " + effectedCriterionId, ParentId: strconv.Itoa(interaction.InteractionId), CriterionInfo: effectedCriterionId, ChildCriterion: []*utils.Node{}})
		data = append(data, &utils.Node{Id: strconv.Itoa(interaction.InteractionId) + " _ " + interactionType, ParentId: strconv.Itoa(interaction.InteractionId), CriterionInfo: interactionType, ChildCriterion: []*utils.Node{}})
		data = append(data, &utils.Node{Id: strconv.Itoa(interaction.InteractionId) + " _ " + isInteractionApplied, ParentId: strconv.Itoa(interaction.InteractionId), CriterionInfo: isInteractionApplied, ChildCriterion: []*utils.Node{}})
		data = append(data, &utils.Node{Id: "Interaction Formula Info _ " + strconv.Itoa(interaction.InteractionId), ParentId: strconv.Itoa(interaction.InteractionId), CriterionInfo: "Interaction Formula Info", ChildCriterion: interactionFormulaLogData})
	}
	root.AddModelTreeInfo(data...)
	model.InteractionTreeDataStructure = root
}

func (model *Model) CreateScoreLog() string {
	result := ""
	criterionOrdered := []*Criterion{}
	for _, criterion := range model.ModelCriterions {
		criterionOrdered = append(criterionOrdered, criterion)
	}
	sort.Sort(CriterionList(criterionOrdered))
	for _, criterion := range criterionOrdered {
		result = result + criterion.CriterionId + "#" + strconv.FormatFloat(criterion.CriterionScore, 'f', -1, 64) + "/"
	}

	for _, interaction := range model.InteractionList {
		result = result + strconv.Itoa(interaction.InteractionId)
		if interaction.InteractionLog.IsInteractionPerformed {
			result = result + "+"
			theLastFormulaOrderNum := ""
			for _, interactionFormulaLog := range interaction.InteractionLog.InteractionFormulasLog {
				if interactionFormulaLog.IsFormulaExecuted {
					theLastFormulaOrderNum = ":" + strconv.Itoa(interactionFormulaLog.FormulaOrderNum)
				}
			}
			result = result + theLastFormulaOrderNum
		} else {
			result = result + "-"
		}
		result = result + "/"
	}

	return result
}

func (model *Model) createChildCriterions(criterion Criterion, modelStructureInfo []utils.ModelStructure) []*Criterion {
	var childCriterionList []*Criterion
	filteredChildCriterionList := utils.Filter(model.ParentChildRelationship, func(val interface{}) bool {
		return val.(utils.ParentChildRelationship).Parent == criterion.CriterionId
	}).([]interface{})

	for _, filteredChildCriterion := range filteredChildCriterionList {
		childCriterionParentInfo := filteredChildCriterion.(utils.ParentChildRelationship)
		childCriterionInfo := utils.Filter(modelStructureInfo, func(val interface{}) bool {
			return val.(utils.ModelStructure).CriterionId == childCriterionParentInfo.Child
		}).([]interface{})[0].(utils.ModelStructure)

		var criterionFormulaLogInfo *PartFormulaLogStructure = nil
		if childCriterionInfo.CriterionFormula.Formula != "" {
			criterionFormulaLogInfo = CreatePartFormulaLogInfo(childCriterionInfo.CriterionFormula.FormulaPriorityNum, childCriterionInfo.CriterionFormula.Formula, childCriterionInfo.CriterionFormula.Result)
		}
		childCriterion := Criterion{CriterionId: childCriterionParentInfo.Child, CriterionName: childCriterionInfo.Name, CriterionWeight: childCriterionInfo.Weight, CriterionScaleLimit: childCriterionInfo.CriterionScaleLimit, CriterionFormula: childCriterionInfo.CriterionFormula, CriterionFormulaLogInfo: criterionFormulaLogInfo}

		doesCriterionHasParent := len(utils.Filter(model.ParentChildRelationship, func(val interface{}) bool {
			return val.(utils.ParentChildRelationship).Parent == childCriterionParentInfo.Child
		}).([]interface{})) > 0
		if doesCriterionHasParent {
			childCriterion.SetChildCriterion(model.createChildCriterions(childCriterion, modelStructureInfo))
		}

		childCriterionList = append(childCriterionList, &childCriterion)
	}
	return childCriterionList
}

func (model *Model) createParentChildRelationList(modelStructureInfo []utils.ModelStructure) {
	var parentChildRelationship []utils.ParentChildRelationship
	for _, modelStructure := range modelStructureInfo {
		structure := modelStructure.Structure
		childCriterionId := modelStructure.CriterionId
		parent := structure[0 : len(structure)-4]
		var parentCriterionId string = "0"
		if parent != "" {
			parentCriterion := utils.Filter(modelStructureInfo, func(val interface{}) bool {
				return val.(utils.ModelStructure).Structure == parent
			})

			parentCriterionId = parentCriterion.([]interface{})[0].(utils.ModelStructure).CriterionId
		}
		parentChildRelationship = append(parentChildRelationship, utils.ParentChildRelationship{parentCriterionId, childCriterionId})
	}
	model.ParentChildRelationship = parentChildRelationship
}

func (model *Model) FindCriterionBelongToModel(modelStructureInfo []utils.ModelStructure) {
	sort.Sort(ModelStructureList(modelStructureInfo))
	model.createParentChildRelationList(modelStructureInfo)
	mainCriterion := Criterion{CriterionId: "0", CriterionName: "Main Criterion", CriterionWeight: 1, CriterionValue: ""}
	mainCriterion.SetChildCriterion(model.createChildCriterions(mainCriterion, modelStructureInfo))
	model.ModelTree = &mainCriterion
}

func (model *Model) SetCriterionValuesBelongToModel() {
	modelCriterion := model.FindModelCriterions()
	for _, criterion := range modelCriterion {
		if criterion.IsLeaf() {
			if model.ScoringCalculationHashMap["D_"+criterion.CriterionId] != nil && !criterion.HasFormula() {
				criterion.CriterionValue = model.ScoringCalculationHashMap["D_"+criterion.CriterionId].(string)
			} else {
				delete(model.ScoringCalculationHashMap, "D_"+criterion.CriterionId)
			}
		}
	}
}

func (model *Model) CreateCriterionScoreLogInfo(title string) {
	modelCriterions := model.FindModelCriterions()
	for _, criterion := range modelCriterions {
		criterionScoreLogInfoWithTheSameTitle := utils.Filter(criterion.CriterionScoreLogInfo, func(val interface{}) bool {
			return val.(*CriterionScoreLogStructure).LogType == title
		}).([]interface{})
		if len(criterionScoreLogInfoWithTheSameTitle) > 0 {
			criterionScoreLogInfoWithTheSameTitle[0].(*CriterionScoreLogStructure).CriterionValue = criterion.CriterionValue
			criterionScoreLogInfoWithTheSameTitle[0].(*CriterionScoreLogStructure).CriterionScore = strconv.FormatFloat(criterion.CriterionScore, 'f', -1, 64)
			criterionScoreLogInfoWithTheSameTitle[0].(*CriterionScoreLogStructure).CriterionWeight = criterion.CriterionWeight
			criterionScoreLogInfoWithTheSameTitle[0].(*CriterionScoreLogStructure).CriterionWeightedScore = strconv.FormatFloat(criterion.CalculateCriterionWeightedScore(model.ModelWeight), 'f', -1, 64)

		} else {
			criterionScoreLogInfo := CriterionScoreLogStructure{LogType: title, CriterionValue: criterion.CriterionValue, CriterionScore: strconv.FormatFloat(criterion.CriterionScore, 'f', -1, 64), CriterionWeight: criterion.CriterionWeight, CriterionWeightedScore: strconv.FormatFloat(criterion.CalculateCriterionWeightedScore(model.ModelWeight), 'f', -1, 64)}
			criterion.CriterionScoreLogInfo = append(criterion.CriterionScoreLogInfo, &criterionScoreLogInfo)
		}
	}
}

func (model *Model) CreateModelLogTree() *utils.Node {
	rootCriterion := model.ModelTree
	var root = &utils.Node{rootCriterion.CriterionId, "", rootCriterion.CriterionName, []*utils.Node{}}
	var data []*utils.Node
	for _, parentChildRelation := range model.ParentChildRelationship {
		criterion := model.FindModelCriterions()[parentChildRelation.Child]
		criterionScoreBeforeInteraction := ""
		if len(criterion.CriterionScoreLogInfo) > 0 {
			criterionScoreBeforeInteraction = " criterionScoreBeforeInteraction : " + criterion.CriterionScoreLogInfo[0].CriterionScore
		}
		criterionValue := " criterionValue: " + criterion.CriterionValue
		criterionWeight := " criterionWeight : " + strconv.FormatFloat(criterion.calculateTotalChildCriterionWeight(), 'f', -1, 64)
		criterionScore := " criterionScore : " + strconv.FormatFloat(criterion.CriterionScore, 'f', -1, 64)
		criterionNormalizedScore := " criterionNormalizedScore : " + strconv.FormatFloat(criterion.CriterionNormalizedScore, 'f', -1, 64)
		criterionWeightedScore := " criterionWeightedScore : " + strconv.FormatFloat(criterion.CalculateCriterionWeightedScore(model.ModelWeight), 'f', -1, 64)
		var criterionInfo = criterion.CriterionId + " -> " + criterion.CriterionName + criterionValue + criterionWeight + criterionScore + criterionScoreBeforeInteraction + criterionNormalizedScore + criterionWeightedScore
		emptyFormula := utils.FormulaStructure{}
		if criterion.CriterionFormula != emptyFormula {
			criterionInfo += " Criterion Formula Log : -->   ("
			criterionFormula := " CriterionFormula :  " + criterion.CriterionFormulaLogInfo.Formula
			isCriterionFormulaExecuted := " Criterion Formula is not executed! " + criterionFormula
			criterionFormulaLastState := ""
			criterionFormulaResult := ""
			if criterion.CriterionFormulaLogInfo.IsFormulaExecuted {
				isCriterionFormulaExecuted = " Criterion Formula is executed "
				criterionFormulaLastState = " criterionFormulaLastState : " + criterion.CriterionFormulaLogInfo.FormulaLastState
				criterionFormulaResult = " criterionFormulaResult :  " + criterion.CriterionFormulaLogInfo.FormulaResult
			}
			criterionInfo += criterionFormula + isCriterionFormulaExecuted + criterionFormulaLastState + criterionFormulaResult + " )"
		}
		data = append(data, &utils.Node{Id: parentChildRelation.Child, ParentId: parentChildRelation.Parent, CriterionInfo: criterionInfo, ChildCriterion: []*utils.Node{}})
	}
	root.AddModelTreeInfo(data...)
	return root
}

func (model *Model) CreateInteractionLogTree() *utils.Node {
	var root = &utils.Node{"Interaction List", "Interaction List", "Interaction List", []*utils.Node{}}
	var data []*utils.Node
	for _, interaction := range model.InteractionList {
		var interactionFormulaLogData []*utils.Node
		for _, interactionFormulaLog := range interaction.InteractionLog.InteractionFormulasLog {
			partFormulaOrderNumber := " partFormulaOrderNumber : " + strconv.Itoa(interactionFormulaLog.FormulaOrderNum)
			partFormula := " partFormula : " + interactionFormulaLog.Formula
			partFormulaResult := " partFormulaResult : " + interactionFormulaLog.FormulaResult
			isPartFormulaExecuted := " Part Formula is not executed"
			partFormulaLastState := ""
			if interactionFormulaLog.IsFormulaExecuted {
				isPartFormulaExecuted = " Part Formula is executed"
				partFormulaLastState = " partFormulaLastState : " + interactionFormulaLog.FormulaLastState
			}
			interactionLogData := partFormulaOrderNumber + partFormula + partFormulaResult + isPartFormulaExecuted + partFormulaLastState
			interactionFormulaLogData = append(interactionFormulaLogData, &utils.Node{Id: strconv.Itoa(interactionFormulaLog.FormulaOrderNum), ParentId: strconv.Itoa(interaction.InteractionId), CriterionInfo: interactionLogData, ChildCriterion: []*utils.Node{}})
		}

		interactionId := " interactionId : " + strconv.Itoa(interaction.InteractionId)
		effectedCriterionId := ""
		if interaction.EffectedCriterion != nil {
			effectedCriterionId = " effectedCriterionId : " + interaction.EffectedCriterion.CriterionId
		}
		interactionBehaviour := " interactionBehaviour : "
		switch interaction.InteractionBehaviour {
		case utils.Behaviour_IncreaseCriterionScore:
			interactionBehaviour += "Behaviour_IncreaseCriterionScore "
		case utils.Behaviour_DecreaseCriterionScore:
			interactionBehaviour += "Behaviour_DecreaseCriterionScore "
		case utils.Behaviour_OverrideCriterionScore:
			interactionBehaviour += "Behaviour_OverrideCriterionScore "
		case utils.Behaviour_OverrideCriterionValue:
			interactionBehaviour += "Behaviour_OverrideCriterionValue "
		case utils.Behaviour_ChangeSensitivityValueOfModel:
			interactionBehaviour += "Behaviour_ChangeSensitivityValueOfModel "
		default:
			interactionBehaviour += "???"
		}
		interactionOrder := " interactionOrder : " + strconv.Itoa(interaction.InteractionOrder)
		interactionType := " interactionType : "
		switch interaction.InteractionType {
		case utils.Interaction_General:
			interactionType += "Interaction_General "
		case utils.Interaction_Sensitive:
			interactionType += "Interaction_Sensitive "
		case utils.Interaction_BeforeCalculation:
			interactionType += "Interaction_BeforeCalculation "
		default:
			interactionType += "???"
		}
		isInteractionApplied := " Interaction is not applied "
		if interaction.InteractionLog.IsInteractionPerformed {
			isInteractionApplied = " Interaction is applied "
		}
		interactionInfo := interactionId + interactionOrder + interactionBehaviour + effectedCriterionId + interactionType + isInteractionApplied
		data = append(data, &utils.Node{Id: strconv.Itoa(interaction.InteractionId), ParentId: "Interaction List", CriterionInfo: interactionInfo, ChildCriterion: interactionFormulaLogData})
	}
	root.AddModelTreeInfo(data...)
	return root
}

func (model *Model) CreateInteractions(interactionInfoList []utils.InteractionStructure) {
	modelCriterion := model.FindModelCriterions()
	for _, interactionInfo := range interactionInfoList {
		interaction := CreateInteraction(model, interactionInfo.InteractionId, modelCriterion[interactionInfo.EffectedCriterionId], utils.InteractionBehaviour(interactionInfo.InteractionBehaviourType), utils.InteractionType(interactionInfo.InteractionType), interactionInfo.PriorityOrderNum)
		var partFormula []*PartFormula
		formulaList := interactionInfo.FormulaList
		isFormulaResultCalculable := true
		for _, formula := range formulaList {
			partFormula = append(partFormula, CreatePartFormula(formula.FormulaPriorityNum, formula.Formula, isFormulaResultCalculable, formula.Result, model))
		}
		sort.Sort(PartFormulaList(partFormula))
		interaction.SetInteractionFormulas(partFormula)
		model.InteractionList = append(model.InteractionList, interaction)
	}
	sort.Sort(InteractionList(model.InteractionList))
}

func (model *Model) ExecuteInteractions() error {
	for _, interaction := range model.InteractionList {
		errExecuteFormula := interaction.ExecuteFormula()
		if !utils.IsAnErrorOccured(errExecuteFormula) {
			errReflectTheBehaviour := interaction.ReflectTheBehaviour()
			if utils.IsAnErrorOccured(errReflectTheBehaviour) {
				return errReflectTheBehaviour
			}
		} else {
			errExecuteFormula = errors.New(strconv.Itoa(interaction.InteractionId) + " - " + errExecuteFormula.Error())
			return errExecuteFormula
		}
	}
	return nil
}

func (model *Model) ControlWhetherCriterionValueIsFound() (bool, error) {
	var criterionWhoseValueIsNotFoundList []*Criterion
	for _, criterion := range model.FindModelCriterions() {
		if criterion.IsLeaf() && criterion.CriterionValue == "" {
			criterionWhoseValueIsNotFoundList = append(criterionWhoseValueIsNotFoundList, criterion)
		}
	}
	var resultOfControl string = ""
	for _, criterionWhoseValueIsNotFound := range criterionWhoseValueIsNotFoundList {
		criterionFormulaLastState := ""
		if criterionWhoseValueIsNotFound.HasFormula() {
			criterionFormulaLastState = " ( criterionFormulaLastState --> " + criterionWhoseValueIsNotFound.CriterionFormulaLogInfo.FormulaLastState + ")"
		}
		resultOfControl = resultOfControl + criterionWhoseValueIsNotFound.CriterionName + criterionFormulaLastState + " ,"
	}

	if resultOfControl != "" {
		resultOfControl = resultOfControl[0 : len(resultOfControl)-1]
		if len(criterionWhoseValueIsNotFoundList) > 1 {
			return true, errors.New(fmt.Sprintf("'%s' values of criterions could not be found.", resultOfControl))
		} else {
			return true, errors.New(fmt.Sprintf("'%s' value of criterion could not be found.", resultOfControl))
		}
	} else {
		return false, nil
	}
}

func (model *Model) FindCriterionValuesByFormula(criterionWhoseFormulaWillExecute map[string]*Criterion) (map[string]*Criterion, []error) {
	var criterionWhoseFormulaCouldNotExecutedList = map[string]*Criterion{}
	var errorList []error
	for _, criterion := range criterionWhoseFormulaWillExecute {
		if criterion.HasFormula() {
			partFormula := CreatePartFormula(1, criterion.CriterionFormula.Formula, true, criterion.CriterionFormula.Result, model)
			if partFormula.Formula == "NOTCALCULABLE" {
				criterion.CriterionValue = "-88888"
				criterion.CriterionScore = -88888
				criterion.CriterionFormulaLogInfo.FormulaLastState = partFormula.Formula
				criterion.CriterionFormulaLogInfo.IsFormulaExecuted = true
				criterion.CriterionFormulaLogInfo.FormulaResult = "-88888"
				model.ScoringCalculationHashMap["D_"+criterion.CriterionId] = "-88888"
			} else {
				result, formulaLastState, valuesOfFormula, err := partFormula.ExecutePartFormula()
				criterion.CriterionFormulaLogInfo.FormulaLastState = formulaLastState
				criterion.CriterionFormulaLogInfo.VariablesOfFormula = valuesOfFormula

				if !utils.IsAnErrorOccured(err) {
					switch result.(type) {
					case string:
						model.ScoringCalculationHashMap["D_"+criterion.CriterionId] = result.(string)
						criterion.CriterionValue = result.(string)
						criterion.CriterionFormulaLogInfo.FormulaResult = result.(string)
					case float64:
						if result.(float64) == -99999 {
							criterion.CriterionWeight = 0
						}
						model.ScoringCalculationHashMap["D_"+criterion.CriterionId] = strconv.FormatFloat(result.(float64), 'f', -1, 64)
						criterion.CriterionValue = strconv.FormatFloat(result.(float64), 'f', -1, 64)
						criterion.CriterionFormulaLogInfo.FormulaResult = strconv.FormatFloat(result.(float64), 'f', -1, 64)
					default:
					}
					criterion.CriterionFormulaLogInfo.IsFormulaExecuted = true
				} else {
					criterionWhoseFormulaCouldNotExecutedList[criterion.CriterionId] = criterion
					errCriterion := errors.New(fmt.Sprintf("An error occured for '%s' criterion : %s", criterion.CriterionName, err.Error()))
					errorList = append(errorList, errCriterion)
				}
			}
		}
	}
	return criterionWhoseFormulaCouldNotExecutedList, errorList
}

func (model *Model) GetCriterionScore(criterionId string) float64 {
	return model.FindModelCriterions()[criterionId].CriterionScore
}

func (model *Model) GetInputValue(value string) string {
	result := model.ScoringCalculationHashMap[value]
	if result == nil {
		return value
	}
	return result.(string)
}

type ModelStructureList []utils.ModelStructure

func (s ModelStructureList) Len() int {
	return len(s)
}
func (s ModelStructureList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ModelStructureList) Less(i, j int) bool {
	return s[i].Structure < s[j].Structure
}

type CriterionList []*Criterion

func (s CriterionList) Len() int {
	return len(s)
}
func (s CriterionList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s CriterionList) Less(i, j int) bool {
	return s[i].CriterionId < s[j].CriterionId
}
