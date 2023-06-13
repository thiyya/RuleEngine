package main

import (
	"RuleEngine/ScoreCalculation"
	"RuleEngine/ScoreCalculationMainConcept"
	"RuleEngine/ScoreCalculationModules"
	"RuleEngine/utils"
	"encoding/json"
	"fmt"
)

func main() {
	BKRSkorHesaplamaVerisi := map[string]interface{}{
		"D_YAS":        "30",
		"D_2":          "200",
		"D_3":          "500",
		"D_6":          "1.5",
		"[Degisken_1]": "5",
	}

	BireyselModelYapiVerisi := []utils.ModelStructure{
		{Name: "Kriter_9", CriterionId: "9", Structure: "...1", Weight: 0, CriterionScaleLimit: "U", CriterionFormula: utils.FormulaStructure{}},
		{Name: "Kriter_8", CriterionId: "8", Structure: "...1...1", Weight: 15, CriterionScaleLimit: "A", CriterionFormula: utils.FormulaStructure{FormulaPriorityNum: 1, Formula: "(3 + 4) * ORT(2,3,4)"}},
		{Name: "Kriter_7", CriterionId: "7", Structure: "...1...2", Weight: 0, CriterionScaleLimit: "U", CriterionFormula: utils.FormulaStructure{}},
		{Name: "Kriter_6", CriterionId: "6", Structure: "...1...2...1", Weight: 5, CriterionScaleLimit: "A", CriterionFormula: utils.FormulaStructure{}},
		{Name: "Kriter_5", CriterionId: "5", Structure: "...1...2...2", Weight: 0, CriterionScaleLimit: "U", CriterionFormula: utils.FormulaStructure{}},
		{Name: "Kriter_4", CriterionId: "4", Structure: "...1...2...2...1", Weight: 0, CriterionScaleLimit: "U", CriterionFormula: utils.FormulaStructure{}},
		{Name: "Kriter_3", CriterionId: "3", Structure: "...1...2...2...1...3", Weight: 15, CriterionScaleLimit: "U", CriterionFormula: utils.FormulaStructure{}},
		{Name: "Kriter_2", CriterionId: "2", Structure: "...1...2...2...1...2", Weight: 10, CriterionScaleLimit: "U", CriterionFormula: utils.FormulaStructure{}},
		{Name: "Kriter_1", CriterionId: "YAS", Structure: "...1...2...2...1...1", Weight: 5, CriterionScaleLimit: "A", CriterionFormula: utils.FormulaStructure{}},
	}

	BKROlcekSecenekVerisi := []utils.ScaleStructure{
		{CriterionId: "YAS", FirstValueOfScale: "0", LastValueOfScale: "20", LowerScore: 1, HigherScore: 1, Explanation: "0-200 arası"},
		{CriterionId: "YAS", FirstValueOfScale: "20", LastValueOfScale: "30", LowerScore: 2, HigherScore: 2, Explanation: "200-400 arası"},
		{CriterionId: "YAS", FirstValueOfScale: "30", LastValueOfScale: "40", LowerScore: 3, HigherScore: 3, Explanation: "400-600 arası"},
		{CriterionId: "2", FirstValueOfScale: "0", LastValueOfScale: "200", LowerScore: 10, HigherScore: 10, Explanation: "0-200 arası"},
		{CriterionId: "2", FirstValueOfScale: "200", LastValueOfScale: "400", LowerScore: 20, HigherScore: 20, Explanation: "201-400 arası"},
		{CriterionId: "2", FirstValueOfScale: "400", LastValueOfScale: "600", LowerScore: 30, HigherScore: 30, Explanation: "401-600 arası"},
		{CriterionId: "3", FirstValueOfScale: "0", LastValueOfScale: "200", LowerScore: 1, HigherScore: 1, Explanation: "0-200 arası"},
		{CriterionId: "3", FirstValueOfScale: "200", LastValueOfScale: "400", LowerScore: 2, HigherScore: 2, Explanation: "200-400 arası"},
		{CriterionId: "3", FirstValueOfScale: "400", LastValueOfScale: "600", LowerScore: 3, HigherScore: 3, Explanation: "400-600 arası"},
		{CriterionId: "6", FirstValueOfScale: "0", LastValueOfScale: "1", LowerScore: 1, HigherScore: 1, Explanation: "0-200 arası"},
		{CriterionId: "6", FirstValueOfScale: "1", LastValueOfScale: "1,50", LowerScore: 2, HigherScore: 2, Explanation: "200-400 arası"},
		{CriterionId: "6", FirstValueOfScale: "1,50", LastValueOfScale: "5", LowerScore: 3, HigherScore: 3, Explanation: "400-600 arası"},
		{CriterionId: "6", FirstValueOfScale: "600", LastValueOfScale: "800", LowerScore: 4, HigherScore: 4, Explanation: "600-800 arası"},
		{CriterionId: "8", FirstValueOfScale: "0", LastValueOfScale: "200", LowerScore: 1, HigherScore: 1, Explanation: "0-200 arası"},
		{CriterionId: "8", FirstValueOfScale: "200", LastValueOfScale: "400", LowerScore: 2, HigherScore: 2, Explanation: "200-400 arası"},
		{CriterionId: "8", FirstValueOfScale: "400", LastValueOfScale: "600", LowerScore: 3, HigherScore: 3, Explanation: "400-600 arası"},
		{CriterionId: "8", FirstValueOfScale: "600", LastValueOfScale: "800", LowerScore: 4, HigherScore: 4, Explanation: "400-600 arası"},
	}
	BKREtkilesimVerisi := []utils.InteractionStructure{
		{InteractionId: 67104, PriorityOrderNum: 9, InteractionType: 4, EffectedCriterionId: "YAS", InteractionBehaviourType: 3, FormulaList: []utils.FormulaStructure{utils.FormulaStructure{FormulaPriorityNum: 1, Formula: "76 > 75", Result: "100"}}},
		{InteractionId: 67105, PriorityOrderNum: 4, InteractionType: 4, EffectedCriterionId: "2", InteractionBehaviourType: 2, FormulaList: []utils.FormulaStructure{utils.FormulaStructure{FormulaPriorityNum: 1, Formula: "([D_YAS] + 4) * -ORT(2,3,4) == -12", Result: "20"}}},
		{InteractionId: 67106, PriorityOrderNum: 1, InteractionType: 4, EffectedCriterionId: "3", InteractionBehaviourType: 1, FormulaList: []utils.FormulaStructure{utils.FormulaStructure{FormulaPriorityNum: 1, Formula: "(30 + 4) * ORT(2,3,4,5) == 3", Result: "25"}}},
		{InteractionId: 67107, PriorityOrderNum: 2, InteractionType: 4, EffectedCriterionId: "6", InteractionBehaviourType: 1, FormulaList: []utils.FormulaStructure{utils.FormulaStructure{FormulaPriorityNum: 1, Formula: "(-30 + 4) * ORT(2,3,4,5)== 4", Result: "30"}}},
		{InteractionId: 65696, PriorityOrderNum: 3, InteractionType: 3, EffectedCriterionId: "", InteractionBehaviourType: 5, FormulaList: []utils.FormulaStructure{utils.FormulaStructure{FormulaPriorityNum: 1, Formula: "(30 - 4) * ORT(2,3,4,5) == 5", Result: "35"}}},
		{InteractionId: 65202, PriorityOrderNum: 6, InteractionType: 3, EffectedCriterionId: "", InteractionBehaviourType: 5, FormulaList: []utils.FormulaStructure{utils.FormulaStructure{FormulaPriorityNum: 1, Formula: "(30 + 4) * ORT(2,3,7,5) == 15", Result: "8"}}},
	}
	BKRModelVerisi := utils.ModelAdditionalInfoStructure{
		ModelId:                       45,
		SegmentationLimit:             "U",
		ModelWeight:                   100,
		MaxCriterionScoreOfModel:      100,
		MinCriterionScoreOfModel:      0,
		MaxModelSensitiveValue:        10,
		MinModelSensitiveValue:        -10,
		ModelWeightDistributionTypeId: utils.DistributeType_General,
	}

	SegmentasyonVerisi10lu := []utils.SegmentationStructure{
		{ScoreSegmentation: "AAA", MinimumPoint: 90, MaximumPoint: 100},
		{ScoreSegmentation: "AA", MinimumPoint: 80, MaximumPoint: 90},
		{ScoreSegmentation: "A", MinimumPoint: 70, MaximumPoint: 80},
		{ScoreSegmentation: "BBB", MinimumPoint: 60, MaximumPoint: 70},
		{ScoreSegmentation: "BB", MinimumPoint: 50, MaximumPoint: 60},
		{ScoreSegmentation: "B", MinimumPoint: 40, MaximumPoint: 50},
		{ScoreSegmentation: "CCC", MinimumPoint: 30, MaximumPoint: 40},
		{ScoreSegmentation: "CC", MinimumPoint: 20, MaximumPoint: 30},
		{ScoreSegmentation: "C", MinimumPoint: 10, MaximumPoint: 20},
		{ScoreSegmentation: "D", MinimumPoint: 0, MaximumPoint: 10},
	}
	SegmentasyonVerisi5li := []utils.SegmentationStructure{
		{ScoreSegmentation: "*****", MinimumPoint: 80, MaximumPoint: 100},
		{ScoreSegmentation: "****", MinimumPoint: 60, MaximumPoint: 80},
		{ScoreSegmentation: "***", MinimumPoint: 40, MaximumPoint: 60},
		{ScoreSegmentation: "**", MinimumPoint: 20, MaximumPoint: 40},
		{ScoreSegmentation: "*", MinimumPoint: 0, MaximumPoint: 20},
	}
	SegmentasyonVerisiRenk := []utils.SegmentationStructure{
		{ScoreSegmentation: "Beyaz", MinimumPoint: 70, MaximumPoint: 100},
		{ScoreSegmentation: "Gri", MinimumPoint: 35, MaximumPoint: 70},
		{ScoreSegmentation: "Siyah", MinimumPoint: 0, MaximumPoint: 35},
	}
	segmentasyonTipi := utils.SegmentationType_10
	var SegmentasyonVerisi []utils.SegmentationStructure
	switch segmentasyonTipi {
	case utils.SegmentationType_10:
		SegmentasyonVerisi = SegmentasyonVerisi10lu
	case utils.SegmentationType_5:
		SegmentasyonVerisi = SegmentasyonVerisi5li
	case utils.SegmentationType_Renk:
		SegmentasyonVerisi = SegmentasyonVerisiRenk
	}

	bireysel := ScoreCalculationModules.MainScoringPointOfView{ScoringCalculationHashMap: BKRSkorHesaplamaVerisi, ModelStructureInfo: BireyselModelYapiVerisi, ScaleInfo: BKROlcekSecenekVerisi, InteractionInfo: BKREtkilesimVerisi, SegmentationInfo: SegmentasyonVerisi, ModelAdditionalInfo: BKRModelVerisi}

	sonuc, err := ScoreCalculation.CalculateScore(bireysel)
	if err == nil {
		model := sonuc["Model"].(*ScoreCalculationMainConcept.Model)
		type SkorHesaplamaOutput struct {
			DemographicalScore           float64
			SegmentScore                 string
			ModelTreeDataStructure       *utils.Node
			InteractionTreeDataStructure *utils.Node
		}
		skorHesaplamaOutput := SkorHesaplamaOutput{DemographicalScore: sonuc["DemographicalScore"].(float64), SegmentScore: sonuc["SegmentScore"].(string), ModelTreeDataStructure: model.ModelTreeDataStructure, InteractionTreeDataStructure: model.InteractionTreeDataStructure}
		mapB, _ := json.Marshal(skorHesaplamaOutput)
		fmt.Println(string(mapB))
	} else {
		fmt.Println(err)
	}
}

