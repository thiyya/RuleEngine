package ScoreCalculationMainConcept

import (
	"RuleEngine/utils"
	"errors"
	"fmt"
	"strconv"
)

type CriterionScoreLogStructure struct {
	LogType                string
	CriterionValue         string
	CriterionScore         string
	CriterionWeight        float64
	CriterionWeightedScore string
}

type Criterion struct {
	CriterionId              string
	CriterionName            string
	CriterionWeight          float64
	CriterionExplanation     string
	CriterionScore           float64
	CriterionNormalizedScore float64
	ChildCriterion           []*Criterion
	ParentCriterion          *Criterion
	CriterionValue           string
	CriterionScoreLogInfo    []*CriterionScoreLogStructure
	CriterionScaleInfo       []utils.ScaleStructure
	CriterionScaleLimit      string
	CriterionFormula         utils.FormulaStructure
	CriterionFormulaLogInfo  *PartFormulaLogStructure
}

func (criterion *Criterion) SetChildCriterion(childCriterions []*Criterion) {
	criterion.ChildCriterion = childCriterions
	for _, c := range childCriterions {
		c.ParentCriterion = criterion
	}
}

func (criterion *Criterion) AddChildCriterion(childCriterion *Criterion) {
	criterion.ChildCriterion = append(criterion.ChildCriterion, childCriterion)
	childCriterion.ParentCriterion = criterion
}

func (criterion *Criterion) CalculateCriterionWeightedScore(modelWeight float64) float64 {
	if criterion.CriterionWeight == 0 {
		return 0
	}
	return (criterion.CriterionNormalizedScore * criterion.CriterionWeight) / modelWeight
}

func (criterion *Criterion) calculateTotalChildCriterionWeight() float64 {
	if !criterion.IsLeaf() {
		var totalChildCriterionWeight float64
		for _, c := range criterion.ChildCriterion {
			totalChildCriterionWeight += c.calculateTotalChildCriterionWeight()
		}
		return totalChildCriterionWeight
	} else {
		return criterion.CriterionWeight
	}
}

func (criterion *Criterion) IsLeaf() bool {
	return criterion.ChildCriterion == nil || len(criterion.ChildCriterion) == 0
}

func (criterion *Criterion) SetScaleInfo(scaleInfo []utils.ScaleStructure) {
	scaleInfoOfCriterion := utils.Filter(scaleInfo, func(val interface{}) bool {
		return val.(utils.ScaleStructure).CriterionId == criterion.CriterionId
	})
	for _, scale := range scaleInfoOfCriterion.([]interface{}) {
		tempScale := scale.(utils.ScaleStructure)
		tempScale.FirstValueOfScale = utils.ConvertCommaToDot(tempScale.FirstValueOfScale)
		tempScale.LastValueOfScale = utils.ConvertCommaToDot(tempScale.LastValueOfScale)

		theLowestScore, _ := strconv.ParseFloat(utils.ConvertCommaToDot(strconv.FormatFloat(tempScale.LowerScore, 'f', -1, 64)), 64)
		tempScale.LowerScore = theLowestScore

		theHighestScore, _ := strconv.ParseFloat(utils.ConvertCommaToDot(strconv.FormatFloat(tempScale.HigherScore, 'f', -1, 64)), 64)
		tempScale.HigherScore = theHighestScore

		criterion.CriterionScaleInfo = append(criterion.CriterionScaleInfo, tempScale)
	}
}

func (criterion *Criterion) CalculateCriterionScore() error {
	isBelongedScaleOfCriterionFoundAndSet := false
	for _, criterionScaleInfo := range criterion.CriterionScaleInfo {
		firstValueOfScale := criterionScaleInfo.FirstValueOfScale
		lastValueOfScale := criterionScaleInfo.LastValueOfScale
		tempCriterionLookupValue, errCriterionLookupValue := strconv.ParseFloat(criterion.CriterionValue, 64)
		tempFirstValueOfScale, errFirstValueOfScale := strconv.ParseFloat(firstValueOfScale, 64)
		tempLastValueOfScale, errLastValueOfScale := strconv.ParseFloat(lastValueOfScale, 64)
		if errCriterionLookupValue == nil && errFirstValueOfScale == nil && errLastValueOfScale == nil {
			if criterion.CriterionScaleLimit == "A" {
				if tempCriterionLookupValue >= tempFirstValueOfScale && tempCriterionLookupValue < tempLastValueOfScale {
					criterion.CriterionScore = (tempCriterionLookupValue-tempFirstValueOfScale)*(criterionScaleInfo.HigherScore-criterionScaleInfo.LowerScore)/(tempLastValueOfScale-tempFirstValueOfScale) + criterionScaleInfo.LowerScore
					isBelongedScaleOfCriterionFoundAndSet = true
					break
				}
			} else {
				if tempCriterionLookupValue > tempFirstValueOfScale && tempCriterionLookupValue <= tempLastValueOfScale {
					criterion.CriterionScore = (tempCriterionLookupValue-tempFirstValueOfScale)*(criterionScaleInfo.HigherScore-criterionScaleInfo.LowerScore)/(tempLastValueOfScale-tempFirstValueOfScale) + criterionScaleInfo.LowerScore
					isBelongedScaleOfCriterionFoundAndSet = true
					break
				}
			}
		} else {
			if criterion.CriterionValue == firstValueOfScale {
				criterion.CriterionScore = criterionScaleInfo.LowerScore
				isBelongedScaleOfCriterionFoundAndSet = true
				break
			}
		}
	}
	if !isBelongedScaleOfCriterionFoundAndSet && len(criterion.CriterionScaleInfo) > 0 {
		return errors.New(fmt.Sprintf("There is no scale for %s value of the %s criteria", criterion.CriterionName, criterion.CriterionValue))
	} else {
		criterion.CriterionNormalizedScore = criterion.CriterionScore
	}
	return nil
}

func (criterion *Criterion) HasFormula() bool {
	return criterion.CriterionFormula.Formula != ""
}
