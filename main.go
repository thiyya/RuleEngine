package main

/*
// Vadeli Fiyatlama Notu
import (
	"gitlab.zfu.fintek.local/zfu/ratings-api/engine/utils"
	"gitlab.zfu.fintek.local/zfu/ratings-api/engine/ScoreCalculationModules"
	"gitlab.zfu.fintek.local/zfu/ratings-api/engine/ScoreCalculation"
	"encoding/json"
	"fmt"
	"gitlab.zfu.fintek.local/zfu/ratings-api/engine/ScoreCalculationMainConcept"
)

func main() {

	scoringCalculationHashMap := map[string]interface{}{
		"[Temdit_Adedi]" : "3",
		"[Temdit_1_Orani]" : "10.50",
		"[Temdit_2_Orani]" : "12.25",
		"[Temdit_3_Orani]" : "11.50",
		"[Temdit_4_Orani]" : "0",
		"[Temdit_1_Max_Yetki]" : "11",
		"[Temdit_2_Max_Yetki]" : "12.50",
		"[Temdit_3_Max_Yetki]" : "12",
		"[Temdit_4_Max_Yetki]" : "1",
		"[Temdit_1_Tabela]" : "5",
		"[Temdit_2_Tabela]" : "4.75",
		"[Temdit_3_Tabela]" : "5.50",
		"[Temdit_4_Tabela]" : "0",
		"[Urun_Tabela]" : "4.5",
		"[Sube_Yetkisi]" : "11.47",
		"[Karlılık_Tutari]" : "135",
		"[Segment]" : "GKTA",
		"[U_1]" : "S",
		"[U_2]" : "A",
		"[LokasyonEkFaizi]" : "1.5",
		"[DOVIZ]" : "TRY",
	}




	modelStructureInfo := []utils.ModelStructure{

		{Name: "Duyarlılık", CriterionId: "3", Structure: "...3", Weight: 100, CriterionScaleLimit: "A", CriterionFormula :utils.FormulaStructure{FormulaPriorityNum:1, Formula:"NOTCALCULABLE"}},
		{Name: "Duyarlılık_Katsayısı", CriterionId: "2", Structure: "...2", Weight: 0, CriterionScaleLimit: "U", CriterionFormula :utils.FormulaStructure{FormulaPriorityNum:1, Formula:"NOTCALCULABLE"}},
		{Name: "Ortalama Duyarlılık", CriterionId: "1", Structure: "...1", Weight: 0, CriterionScaleLimit: "A", CriterionFormula :utils.FormulaStructure{}},
		{Name: "Temdit_1", CriterionId: "11", Structure: "...1...1", Weight: 100, CriterionScaleLimit: "A", CriterionFormula :utils.FormulaStructure{FormulaPriorityNum:1, Formula:"[Temdit_Adedi] != 1 ? -99999 : (([Temdit_1_Orani] > [Temdit_1_Max_Yetki] ? [Temdit_1_Max_Yetki] : [Temdit_1_Orani]) - [Temdit_1_Tabela]) / ([Temdit_1_Max_Yetki] - [Temdit_1_Tabela]) * 100 "}},
		{Name: "Temdit_2_1", CriterionId: "12", Structure: "...1...2", Weight: 50, CriterionScaleLimit: "A", CriterionFormula :utils.FormulaStructure{FormulaPriorityNum:1, Formula:"[Temdit_Adedi] != 2 ? -99999 : (([Temdit_1_Orani] > [Temdit_1_Max_Yetki] ? [Temdit_1_Max_Yetki] : [Temdit_1_Orani]) - [Temdit_1_Tabela]) / ([Temdit_1_Max_Yetki] - [Temdit_1_Tabela]) * 100 "}},
		{Name: "Temdit_2_2", CriterionId: "13", Structure: "...1...3", Weight: 50, CriterionScaleLimit: "A", CriterionFormula :utils.FormulaStructure{FormulaPriorityNum:1, Formula:"[Temdit_Adedi] != 2 ? -99999 : (([Temdit_2_Orani] > [Temdit_2_Max_Yetki] ? [Temdit_2_Max_Yetki] : [Temdit_2_Orani]) - [Temdit_2_Tabela]) / ([Temdit_2_Max_Yetki] - [Temdit_2_Tabela]) * 100 "}},
		{Name: "Temdit_3_1", CriterionId: "14", Structure: "...1...4", Weight: 34, CriterionScaleLimit: "A", CriterionFormula :utils.FormulaStructure{FormulaPriorityNum:1, Formula:"[Temdit_Adedi] != 3 ? -99999 : (([Temdit_1_Orani] > [Temdit_1_Max_Yetki] ? [Temdit_1_Max_Yetki] : [Temdit_1_Orani]) - [Temdit_1_Tabela]) / ([Temdit_1_Max_Yetki] - [Temdit_1_Tabela]) * 100 "}},
		{Name: "Temdit_3_2", CriterionId: "15", Structure: "...1...5", Weight: 33, CriterionScaleLimit: "A", CriterionFormula :utils.FormulaStructure{FormulaPriorityNum:1, Formula:"[Temdit_Adedi] != 3 ? -99999 : (([Temdit_2_Orani] > [Temdit_2_Max_Yetki] ? [Temdit_2_Max_Yetki] : [Temdit_2_Orani]) - [Temdit_2_Tabela]) / ([Temdit_2_Max_Yetki] - [Temdit_2_Tabela]) * 100 "}},
		{Name: "Temdit_3_3", CriterionId: "16", Structure: "...1...6", Weight: 33, CriterionScaleLimit: "A", CriterionFormula :utils.FormulaStructure{FormulaPriorityNum:1, Formula:"[Temdit_Adedi] != 3 ? -99999 : (([Temdit_3_Orani] > [Temdit_3_Max_Yetki] ? [Temdit_3_Max_Yetki] : [Temdit_3_Orani]) - [Temdit_3_Tabela]) / ([Temdit_3_Max_Yetki] - [Temdit_3_Tabela]) * 100 "}},
		{Name: "Temdit_4_1", CriterionId: "17", Structure: "...1...7", Weight: 25, CriterionScaleLimit: "A", CriterionFormula :utils.FormulaStructure{FormulaPriorityNum:1, Formula:"[Temdit_Adedi] != 4 ? -99999 : (([Temdit_1_Orani] > [Temdit_1_Max_Yetki] ? [Temdit_1_Max_Yetki] : [Temdit_1_Orani]) - [Temdit_1_Tabela]) / ([Temdit_1_Max_Yetki] - [Temdit_1_Tabela]) * 100 "}},
		{Name: "Temdit_4_2", CriterionId: "18", Structure: "...1...8", Weight: 25, CriterionScaleLimit: "A", CriterionFormula :utils.FormulaStructure{FormulaPriorityNum:1, Formula:"[Temdit_Adedi] != 4 ? -99999 : (([Temdit_2_Orani] > [Temdit_2_Max_Yetki] ? [Temdit_2_Max_Yetki] : [Temdit_2_Orani]) - [Temdit_2_Tabela]) / ([Temdit_2_Max_Yetki] - [Temdit_2_Tabela]) * 100 "}},
		{Name: "Temdit_4_3", CriterionId: "19", Structure: "...1...9", Weight: 25, CriterionScaleLimit: "A", CriterionFormula :utils.FormulaStructure{FormulaPriorityNum:1, Formula:"[Temdit_Adedi] != 4 ? -99999 : (([Temdit_3_Orani] > [Temdit_3_Max_Yetki] ? [Temdit_3_Max_Yetki] : [Temdit_3_Orani]) - [Temdit_3_Tabela]) / ([Temdit_3_Max_Yetki] - [Temdit_3_Tabela]) * 100 "}},
		{Name: "Temdit_4_4", CriterionId: "20", Structure: "...1..10", Weight: 25, CriterionScaleLimit: "A", CriterionFormula :utils.FormulaStructure{FormulaPriorityNum:1, Formula:"[Temdit_Adedi] != 4 ? -99999 : (([Temdit_4_Orani] > [Temdit_4_Max_Yetki] ? [Temdit_4_Max_Yetki] : [Temdit_4_Orani]) - [Temdit_4_Tabela]) / ([Temdit_4_Max_Yetki] - [Temdit_4_Tabela]) * 100 "}},

		{Name: "Varlık-Karlılık", CriterionId: "4", Structure: "...4", Weight: 100, CriterionScaleLimit: "A", CriterionFormula :utils.FormulaStructure{FormulaPriorityNum:1, Formula:"NOTCALCULABLE"}},
		{Name: "Varlık", CriterionId: "5", Structure: "...5", Weight: 0, CriterionScaleLimit: "A", CriterionFormula :utils.FormulaStructure{FormulaPriorityNum:1, Formula:"[Segment]"}},
		{Name: "Karlılık", CriterionId: "6", Structure: "...6", Weight: 0, CriterionScaleLimit: "U", CriterionFormula :utils.FormulaStructure{FormulaPriorityNum:1, Formula:"[Karlılık_Tutari]"}},

		{Name: "Sahiplik", CriterionId: "7", Structure: "...7", Weight: 0, CriterionScaleLimit: "A", CriterionFormula :utils.FormulaStructure{}},
		{Name: "Sahiplik_Urun_1", CriterionId: "71", Structure: "...7...1", Weight: 100, CriterionScaleLimit: "A", CriterionFormula :utils.FormulaStructure{FormulaPriorityNum:1, Formula:"NVL([U_1],'H') == 'S' || NVL([U_1],'H') == 'A' ? 'Sahip' : 'Degil'"}},
		{Name: "Sahiplik_Urun_2", CriterionId: "72", Structure: "...7...2", Weight: 100, CriterionScaleLimit: "A", CriterionFormula :utils.FormulaStructure{FormulaPriorityNum:1, Formula:"NVL([U_2],'H') == 'S' || NVL([U_2],'H') == 'A' ? 'Sahip' : 'Degil'"}},

		{Name: "Aktiflik", CriterionId: "8", Structure: "...8", Weight: 0, CriterionScaleLimit: "A", CriterionFormula :utils.FormulaStructure{}},
		{Name: "Aktiflik_Urun_1", CriterionId: "81", Structure: "...8...1", Weight: 100, CriterionScaleLimit: "A", CriterionFormula :utils.FormulaStructure{FormulaPriorityNum:1, Formula:"NVL([U_1],'H') == 'A' ? 'Aktif' : 'Degil'"}},
		{Name: "Aktiflik_Urun_2", CriterionId: "82", Structure: "...8...2", Weight: 100, CriterionScaleLimit: "A", CriterionFormula :utils.FormulaStructure{FormulaPriorityNum:1, Formula:"NVL([U_2],'H') == 'A' ? 'Aktif' : 'Degil'"}},

		{Name: "Lokasyon", CriterionId: "9", Structure: "...9", Weight: 100, CriterionScaleLimit: "A", CriterionFormula :utils.FormulaStructure{FormulaPriorityNum:1, Formula:"[LokasyonEkFaizi]" }},
		{Name: "Sahiplik - Aktiflik", CriterionId: "10", Structure: "..10", Weight: 100, CriterionScaleLimit: "A", CriterionFormula :utils.FormulaStructure{FormulaPriorityNum:1, Formula:"NOTCALCULABLE"}},

	}

	ScaleInfo := []utils.ScaleStructure{

		{CriterionId: "9", FirstValueOfScale: "-99", LastValueOfScale: "99", LowerScore: -99, HigherScore: 99, Explanation: "Lokasyon"},

		{CriterionId: "71", FirstValueOfScale: "Sahip", LowerScore: 2.2, Explanation: "Internet Bankacılığı"},
		{CriterionId: "71", FirstValueOfScale: "Degil", LowerScore: 0, Explanation: "Internet Bankacılığı"},
		{CriterionId: "72", FirstValueOfScale: "Sahip", LowerScore: 2.3, Explanation: "OÖT Sahiplik"},
		{CriterionId: "72", FirstValueOfScale: "Degil", LowerScore: 0, Explanation: "OÖT Sahiplik"},
		{CriterionId: "81", FirstValueOfScale: "Aktif", LowerScore: 2.4, Explanation: "Internet Bankacılığı"},
		{CriterionId: "81", FirstValueOfScale: "Degil", LowerScore: 0, Explanation: "Internet Bankacılığı"},
		{CriterionId: "82", FirstValueOfScale: "Aktif", LowerScore: 2.6, Explanation: "OÖT Aktiflik"},
		{CriterionId: "82", FirstValueOfScale: "Degil", LowerScore: 0, Explanation: "OÖT Aktiflik"},

		{CriterionId: "11", FirstValueOfScale: "-99999", LastValueOfScale: "99999", LowerScore: -99999, HigherScore: 99999, Explanation: "1. Temdit"},
		{CriterionId: "12", FirstValueOfScale: "-99999", LastValueOfScale: "99999", LowerScore: -99999, HigherScore: 99999, Explanation: "1. Temdit"},
		{CriterionId: "13", FirstValueOfScale: "-99999", LastValueOfScale: "99999", LowerScore: -99999, HigherScore: 99999, Explanation: "2. Temdit"},
		{CriterionId: "14", FirstValueOfScale: "-99999", LastValueOfScale: "99999", LowerScore: -99999, HigherScore: 99999, Explanation: "1. Temdit"},
		{CriterionId: "15", FirstValueOfScale: "-99999", LastValueOfScale: "99999", LowerScore: -99999, HigherScore: 99999, Explanation: "2. Temdit"},
		{CriterionId: "16", FirstValueOfScale: "-99999", LastValueOfScale: "99999", LowerScore: -99999, HigherScore: 99999, Explanation: "3. Temdit"},
		{CriterionId: "17", FirstValueOfScale: "-99999", LastValueOfScale: "99999", LowerScore: -99999, HigherScore: 99999, Explanation: "1. Temdit"},
		{CriterionId: "18", FirstValueOfScale: "-99999", LastValueOfScale: "99999", LowerScore: -99999, HigherScore: 99999, Explanation: "2. Temdit"},
		{CriterionId: "19", FirstValueOfScale: "-99999", LastValueOfScale: "99999", LowerScore: -99999, HigherScore: 99999, Explanation: "3. Temdit"},
		{CriterionId: "20", FirstValueOfScale: "-99999", LastValueOfScale: "99999", LowerScore: -99999, HigherScore: 99999, Explanation: "4. Temdit"},

		{CriterionId: "2", FirstValueOfScale: "0", LastValueOfScale: "20", LowerScore: 3.75, HigherScore: 2.50, Explanation: "Duyarlılık_Katsayısı"},
		{CriterionId: "2", FirstValueOfScale: "20", LastValueOfScale: "40", LowerScore: 2.50, HigherScore: 2, Explanation: "Duyarlılık_Katsayısı"},
		{CriterionId: "2", FirstValueOfScale: "40", LastValueOfScale: "60", LowerScore: 2, HigherScore: 1.50, Explanation: "Duyarlılık_Katsayısı"},
		{CriterionId: "2", FirstValueOfScale: "60", LastValueOfScale: "75", LowerScore: 1.50, HigherScore: 1.25, Explanation: "Duyarlılık_Katsayısı"},
		{CriterionId: "2", FirstValueOfScale: "75", LastValueOfScale: "90", LowerScore: 1.25, HigherScore: 1, Explanation: "Duyarlılık_Katsayısı"},
		{CriterionId: "2", FirstValueOfScale: "90", LastValueOfScale: "100", LowerScore: 1, HigherScore: 0.98, Explanation: "Duyarlılık_Katsayısı"},

		{CriterionId: "6", FirstValueOfScale: "-9999999", LastValueOfScale: "80", LowerScore: 1, HigherScore: 1, Explanation: "Zararlı"},
		{CriterionId: "6", FirstValueOfScale: "80", LastValueOfScale: "130", LowerScore: 2, HigherScore: 2, Explanation: "Çok Az Karlı"},
		{CriterionId: "6", FirstValueOfScale: "130", LastValueOfScale: "185", LowerScore: 3, HigherScore: 3, Explanation: "Az Karlı"},
		{CriterionId: "6", FirstValueOfScale: "185", LastValueOfScale: "260", LowerScore: 4, HigherScore: 4, Explanation: "Orta Karlı"},
		{CriterionId: "6", FirstValueOfScale: "260", LastValueOfScale: "375", LowerScore: 5, HigherScore: 5, Explanation: "Karlı"},
		{CriterionId: "6", FirstValueOfScale: "375", LastValueOfScale: "625", LowerScore: 6, HigherScore: 6, Explanation: "Çok Karlı"},
		{CriterionId: "6", FirstValueOfScale: "625", LastValueOfScale: "9999999", LowerScore: 7, HigherScore: 7, Explanation: "Yüksek Karlı"},

		{CriterionId: "5", FirstValueOfScale: "GM", LowerScore: 0, Explanation: "Bireysel Diğer"},
		{CriterionId: "5", FirstValueOfScale: "GKTA", LowerScore: 1, Explanation: "Bireysel Kitle"},
		{CriterionId: "5", FirstValueOfScale: "BKV", LowerScore: 2, Explanation: "Bireysel Kitle Artı"},
		{CriterionId: "5", FirstValueOfScale: "BV", LowerScore: 3, Explanation: "Bireysel Kitle Varlıklı"},
		{CriterionId: "5", FirstValueOfScale: "GKTB", LowerScore: 4, Explanation: "Bireysel Özel"},
		{CriterionId: "5", FirstValueOfScale: "BD", LowerScore: 5, Explanation: "Bireysel Varlıklı"},
		{CriterionId: "5", FirstValueOfScale: "BO", LowerScore: 6, Explanation: "Girişimci Kitle A"},
		{CriterionId: "5", FirstValueOfScale: "BK", LowerScore: 7, Explanation: "Girişimci Kitle B"},
		{CriterionId: "5", FirstValueOfScale: "BKA", LowerScore: 8, Explanation: "Girişimci Kitle C"},
		{CriterionId: "5", FirstValueOfScale: "GKTC", LowerScore: 9, Explanation: "Girişimci Mikro"},

		{CriterionId: "4", FirstValueOfScale: "10", LowerScore: 0, Explanation: "Zararlı - Bireysel Diğer"},
		{CriterionId: "4", FirstValueOfScale: "11", LowerScore: -0.20, Explanation: "Zararlı - Bireysel Kitle"},
		{CriterionId: "4", FirstValueOfScale: "12", LowerScore: -0.40, Explanation: "Zararlı - Bireysel Kitle Artı"},
		{CriterionId: "4", FirstValueOfScale: "13", LowerScore: -0.40, Explanation: "Zararlı - Bireysel Kitle Varlıklı"},
		{CriterionId: "4", FirstValueOfScale: "14", LowerScore: -0.20, Explanation: "Zararlı - Bireysel Özel"},
		{CriterionId: "4", FirstValueOfScale: "15", LowerScore: -0.40, Explanation: "Zararlı - Bireysel Varlıklı"},
		{CriterionId: "4", FirstValueOfScale: "16", LowerScore: -0.40, Explanation: "Zararlı - Girişimci Kitle A"},
		{CriterionId: "4", FirstValueOfScale: "17", LowerScore: -0.40, Explanation: "Zararlı - Girişimci Kitle B"},
		{CriterionId: "4", FirstValueOfScale: "18", LowerScore: -0.40, Explanation: "Zararlı - Girişimci Kitle C"},
		{CriterionId: "4", FirstValueOfScale: "19", LowerScore: -0.40, Explanation: "Zararlı - Girişimci Mikro"},

		{CriterionId: "4", FirstValueOfScale: "20", LowerScore: 0, Explanation: "Çok Az Karlı - Bireysel Diğer"},
		{CriterionId: "4", FirstValueOfScale: "21", LowerScore: -0.10, Explanation: "Çok Az Karlı - Bireysel Kitle"},
		{CriterionId: "4", FirstValueOfScale: "22", LowerScore: -0.20, Explanation: "Çok Az Karlı - Bireysel Kitle Artı"},
		{CriterionId: "4", FirstValueOfScale: "23", LowerScore: -0.20, Explanation: "Çok Az Karlı - Bireysel Kitle Varlıklı"},
		{CriterionId: "4", FirstValueOfScale: "24", LowerScore: -0.10, Explanation: "Çok Az Karlı - Bireysel Özel"},
		{CriterionId: "4", FirstValueOfScale: "25", LowerScore: -0.30, Explanation: "Çok Az Karlı - Bireysel Varlıklı"},
		{CriterionId: "4", FirstValueOfScale: "26", LowerScore: -0.20, Explanation: "Çok Az Karlı - Girişimci Kitle A"},
		{CriterionId: "4", FirstValueOfScale: "27", LowerScore: -0.20, Explanation: "Çok Az Karlı - Girişimci Kitle B"},
		{CriterionId: "4", FirstValueOfScale: "28", LowerScore: -0.20, Explanation: "Çok Az Karlı - Girişimci Kitle C"},
		{CriterionId: "4", FirstValueOfScale: "29", LowerScore: -0.20, Explanation: "Çok Az Karlı - Girişimci Mikro"},

		{CriterionId: "4", FirstValueOfScale: "30", LowerScore: 0, Explanation: "Az Karlı - Bireysel Diğer"},
		{CriterionId: "4", FirstValueOfScale: "31", LowerScore: 0.10, Explanation: "Az Karlı - Bireysel Kitle"},
		{CriterionId: "4", FirstValueOfScale: "32", LowerScore: -0.10, Explanation: "Az Karlı - Bireysel Kitle Artı"},
		{CriterionId: "4", FirstValueOfScale: "33", LowerScore: -0.10, Explanation: "Az Karlı - Bireysel Kitle Varlıklı"},
		{CriterionId: "4", FirstValueOfScale: "34", LowerScore: 0, Explanation: "Az Karlı - Bireysel Özel"},
		{CriterionId: "4", FirstValueOfScale: "35", LowerScore: -0.20, Explanation: "Az Karlı - Bireysel Varlıklı"},
		{CriterionId: "4", FirstValueOfScale: "36", LowerScore: 0, Explanation: "Az Karlı - Girişimci Kitle A"},
		{CriterionId: "4", FirstValueOfScale: "37", LowerScore: 0, Explanation: "Az Karlı - Girişimci Kitle B"},
		{CriterionId: "4", FirstValueOfScale: "38", LowerScore: 0, Explanation: "Az Karlı - Girişimci Kitle C"},
		{CriterionId: "4", FirstValueOfScale: "39", LowerScore: 0, Explanation: "Az Karlı - Girişimci Mikro"},

		{CriterionId: "4", FirstValueOfScale: "40", LowerScore: 0, Explanation: "Orta Karlı - Bireysel Diğer"},
		{CriterionId: "4", FirstValueOfScale: "41", LowerScore: 0.30, Explanation: "Orta Karlı - Bireysel Kitle"},
		{CriterionId: "4", FirstValueOfScale: "42", LowerScore: 0, Explanation: "Orta Karlı - Bireysel Kitle Artı"},
		{CriterionId: "4", FirstValueOfScale: "43", LowerScore: 0, Explanation: "Orta Karlı - Bireysel Kitle Varlıklı"},
		{CriterionId: "4", FirstValueOfScale: "44", LowerScore: 0.10, Explanation: "Orta Karlı - Bireysel Özel"},
		{CriterionId: "4", FirstValueOfScale: "45", LowerScore: -0.10, Explanation: "Orta Karlı - Bireysel Varlıklı"},
		{CriterionId: "4", FirstValueOfScale: "46", LowerScore: 0.20, Explanation: "Orta Karlı - Girişimci Kitle A"},
		{CriterionId: "4", FirstValueOfScale: "47", LowerScore: 0.20, Explanation: "Orta Karlı - Girişimci Kitle B"},
		{CriterionId: "4", FirstValueOfScale: "48", LowerScore: 0.20, Explanation: "Orta Karlı - Girişimci Kitle C"},
		{CriterionId: "4", FirstValueOfScale: "49", LowerScore: 0.20, Explanation: "Orta Karlı - Girişimci Mikro"},

		{CriterionId: "4", FirstValueOfScale: "50", LowerScore: 0, Explanation: "Karlı - Bireysel Diğer"},
		{CriterionId: "4", FirstValueOfScale: "51", LowerScore: 0.40, Explanation: "Karlı - Bireysel Kitle"},
		{CriterionId: "4", FirstValueOfScale: "52", LowerScore: 0.10, Explanation: "Karlı - Bireysel Kitle Artı"},
		{CriterionId: "4", FirstValueOfScale: "53", LowerScore: 0.10, Explanation: "Karlı - Bireysel Kitle Varlıklı"},
		{CriterionId: "4", FirstValueOfScale: "54", LowerScore: 0.20, Explanation: "Karlı - Bireysel Özel"},
		{CriterionId: "4", FirstValueOfScale: "55", LowerScore: 0.10, Explanation: "Karlı - Bireysel Varlıklı"},
		{CriterionId: "4", FirstValueOfScale: "56", LowerScore: 0.40, Explanation: "Karlı - Girişimci Kitle A"},
		{CriterionId: "4", FirstValueOfScale: "57", LowerScore: 0.40, Explanation: "Karlı - Girişimci Kitle B"},
		{CriterionId: "4", FirstValueOfScale: "58", LowerScore: 0.40, Explanation: "Karlı - Girişimci Kitle C"},
		{CriterionId: "4", FirstValueOfScale: "59", LowerScore: 0.40, Explanation: "Karlı - Girişimci Mikro"},

		{CriterionId: "4", FirstValueOfScale: "60", LowerScore: 0, Explanation: "Çok Karlı - Bireysel Diğer"},
		{CriterionId: "4", FirstValueOfScale: "61", LowerScore: 0.50, Explanation: "Çok Karlı - Bireysel Kitle"},
		{CriterionId: "4", FirstValueOfScale: "62", LowerScore: 0.30, Explanation: "Çok Karlı - Bireysel Kitle Artı"},
		{CriterionId: "4", FirstValueOfScale: "63", LowerScore: 0.30, Explanation: "Çok Karlı - Bireysel Kitle Varlıklı"},
		{CriterionId: "4", FirstValueOfScale: "64", LowerScore: 0.30, Explanation: "Çok Karlı - Bireysel Özel"},
		{CriterionId: "4", FirstValueOfScale: "65", LowerScore: 0.30, Explanation: "Çok Karlı - Bireysel Varlıklı"},
		{CriterionId: "4", FirstValueOfScale: "66", LowerScore: 0.50, Explanation: "Çok Karlı - Girişimci Kitle A"},
		{CriterionId: "4", FirstValueOfScale: "67", LowerScore: 0.50, Explanation: "Çok Karlı - Girişimci Kitle B"},
		{CriterionId: "4", FirstValueOfScale: "68", LowerScore: 0.50, Explanation: "Çok Karlı - Girişimci Kitle C"},
		{CriterionId: "4", FirstValueOfScale: "69", LowerScore: 0.50, Explanation: "Çok Karlı - Girişimci Mikro"},

		{CriterionId: "4", FirstValueOfScale: "70", LowerScore: 0, Explanation: "Yüksek Karlı - Bireysel Diğer"},
		{CriterionId: "4", FirstValueOfScale: "71", LowerScore: 0.60, Explanation: "Yüksek Karlı - Bireysel Kitle"},
		{CriterionId: "4", FirstValueOfScale: "72", LowerScore: 0.50, Explanation: "Yüksek Karlı - Bireysel Kitle Artı"},
		{CriterionId: "4", FirstValueOfScale: "73", LowerScore: 0.50, Explanation: "Yüksek Karlı - Bireysel Kitle Varlıklı"},
		{CriterionId: "4", FirstValueOfScale: "74", LowerScore: 0.40, Explanation: "Yüksek Karlı - Bireysel Özel"},
		{CriterionId: "4", FirstValueOfScale: "75", LowerScore: 0.50, Explanation: "Yüksek Karlı - Bireysel Varlıklı"},
		{CriterionId: "4", FirstValueOfScale: "76", LowerScore: 0.60, Explanation: "Yüksek Karlı - Girişimci Kitle A"},
		{CriterionId: "4", FirstValueOfScale: "77", LowerScore: 0.60, Explanation: "Yüksek Karlı - Girişimci Kitle B"},
		{CriterionId: "4", FirstValueOfScale: "78", LowerScore: 0.60, Explanation: "Yüksek Karlı - Girişimci Kitle C"},
		{CriterionId: "4", FirstValueOfScale: "79", LowerScore: 0.60, Explanation: "Yüksek Karlı - Girişimci Mikro"},

	}

	InteractionInfo := []utils.InteractionStructure{
		{InteractionId:1, PriorityOrderNum:1, InteractionType:4, EffectedCriterionId:"2", InteractionBehaviourType:4, FormulaList:[]utils.FormulaStructure{utils.FormulaStructure{FormulaPriorityNum:1, Formula:"1 > 0", Result:"[P_1]"}}},
		{InteractionId:2, PriorityOrderNum:2, InteractionType:4, EffectedCriterionId:"3", InteractionBehaviourType:3, FormulaList:[]utils.FormulaStructure{utils.FormulaStructure{FormulaPriorityNum:1, Formula:"1 > 0", Result:"[Urun_Tabela] + ([Sube_Yetkisi] - [Urun_Tabela]) * [P_1] * [P_2] / 100"}}},
		{InteractionId:3, PriorityOrderNum:3, InteractionType:4, EffectedCriterionId:"4", InteractionBehaviourType:4, FormulaList:[]utils.FormulaStructure{utils.FormulaStructure{FormulaPriorityNum:1, Formula:"1 > 0", Result:"[P_6] * 10 + [P_5]"}}},

		{InteractionId:4, PriorityOrderNum:4, InteractionType:4, EffectedCriterionId:"7", InteractionBehaviourType:3, FormulaList:[]utils.FormulaStructure{utils.FormulaStructure{FormulaPriorityNum:1, Formula:"[P_7] > ([DOVIZ] == 'TRY' ? 1.5 : ([DOVIZ] == 'USD' ? 0.6 : ([DOVIZ] == 'EUR' ? 0.4 : 99999)))  ", Result:"[DOVIZ] == 'TRY' ? 1.5 : ([DOVIZ] == 'USD' ? 0.6 : ([DOVIZ] == 'EUR' ? 0.4 : 99999))"}}},
		{InteractionId:5, PriorityOrderNum:5, InteractionType:4, EffectedCriterionId:"8", InteractionBehaviourType:3, FormulaList:[]utils.FormulaStructure{utils.FormulaStructure{FormulaPriorityNum:1, Formula:"[P_8] > ([DOVIZ] == 'TRY' ? 1.5 : ([DOVIZ] == 'USD' ? 0.6 : ([DOVIZ] == 'EUR' ? 0.4 : 99999)))  ", Result:"[DOVIZ] == 'TRY' ? 1.5 : ([DOVIZ] == 'USD' ? 0.6 : ([DOVIZ] == 'EUR' ? 0.4 : 99999))"}}},
		{InteractionId:6, PriorityOrderNum:6, InteractionType:4, EffectedCriterionId:"10", InteractionBehaviourType:3, FormulaList:[]utils.FormulaStructure{utils.FormulaStructure{FormulaPriorityNum:1, Formula:"1 > 0", Result:"[P_7] + [P_8]"}}},

	}

	modelAdditionalInfo := utils.ModelAdditionalInfoStructure{
		ModelId: 1500,
		ModelWeight: 100,
		ModelWeightDistributionTypeId : utils.DistributeType_DontDistribute,
	}

	deposit := ScoreCalculationModules.DepositProposedRateCalculation{ScoringCalculationHashMap: scoringCalculationHashMap, ModelStructureInfo: modelStructureInfo, ScaleInfo: ScaleInfo, InteractionInfo : InteractionInfo, ModelAdditionalInfo: modelAdditionalInfo }
	sonuc, err := ScoreCalculation.CalculateScore(deposit)

	if err == nil {
		model := sonuc["Model"].(*ScoreCalculationMainConcept.Model)
		type SkorHesaplamaOutput struct {
			DemographicalScore           float64
			SegmentScore              string
			ModelTreeDataStructure     *utils.Node
			InteractionTreeDataStructure *utils.Node
		}
		skorHesaplamaOutput := SkorHesaplamaOutput{DemographicalScore : sonuc["DemographicalScore"].(float64), ModelTreeDataStructure: model.ModelTreeDataStructure, InteractionTreeDataStructure: model.InteractionTreeDataStructure}
		mapB, _ := json.Marshal(skorHesaplamaOutput)
		fmt.Println(string(mapB))
	} else {
		fmt.Println(err)
	}
}
*/
/*
// Vadeli indirim yetkileri
import (
	"gitlab.zfu.fintek.local/zfu/ratings-api/engine/utils"
	"gitlab.zfu.fintek.local/zfu/ratings-api/engine/ScoreCalculationModules"
	"gitlab.zfu.fintek.local/zfu/ratings-api/engine/ScoreCalculation"
	"encoding/json"
	"fmt"
	"gitlab.zfu.fintek.local/zfu/ratings-api/engine/ScoreCalculationMainConcept"
)
func main() {

	BKRSkorHesaplamaVerisi := map[string]interface{}{
		"[Max_Faiz]" : "5",
		"[Donem_Sonu_Ek_Faiz]" : "10",
		"[Yeni_Mus_Ek_Faiz]" : "5",
		"[Lokasyon_Ek_Faiz]" : "5",
		"[AcikHesaplarinMinAcilisTar]" : "",
		"[KapaliHesaplarinMaxKapanisTar]" : "",
		"[Gun_Tarihi]" : "2017-04-11",
		"[Vade_Sonu_Tarihi]" : "2017-09-27",
		"[Gun_Sayisi]" : "100",
		"[Doviz]" : "88",
		"[Tutar]" : "4000",
	}

	BireyselModelYapiVerisi := []utils.ModelStructure{
		{Name: "Max Faiz", CriterionId: "1", Structure: "...1", Weight: 100, CriterionScaleLimit: "U", CriterionFormula : utils.FormulaStructure{}},
		{Name: "Dönem Sonu Ek Faiz", CriterionId: "2", Structure: "...2", Weight: 100, CriterionScaleLimit: "U", CriterionFormula : utils.FormulaStructure{}},
		{Name: "Yeni Müşteri Ek Faiz", CriterionId: "3", Structure: "...3", Weight: 100, CriterionScaleLimit: "U", CriterionFormula : utils.FormulaStructure{}},
		{Name: "Lokasyon Ek Faiz", CriterionId: "4", Structure: "...4", Weight: 100, CriterionScaleLimit: "U", CriterionFormula : utils.FormulaStructure{}},
	}

	BKREtkilesimVerisi := []utils.InteractionStructure{
		{InteractionId:1, PriorityOrderNum:1, InteractionType:9, EffectedCriterionId:"1", InteractionBehaviourType:3, FormulaList:[]utils.FormulaStructure{utils.FormulaStructure{FormulaPriorityNum:1, Formula:"1>0", Result:"[Max_Faiz]"}}},
		{InteractionId:2, PriorityOrderNum:2, InteractionType:9, EffectedCriterionId:"2", InteractionBehaviourType:3, FormulaList:[]utils.FormulaStructure{utils.FormulaStructure{FormulaPriorityNum:1, Formula:"ADD([Vade_Sonu_Tarihi],5) >= EOQ([Vade_Sonu_Tarihi]) && [Vade_Sonu_Tarihi] <= EOQ([Vade_Sonu_Tarihi])", Result:"[Donem_Sonu_Ek_Faiz]"}}},
		{InteractionId:3, PriorityOrderNum:3, InteractionType:9, EffectedCriterionId:"3", InteractionBehaviourType:3, FormulaList:[]utils.FormulaStructure{utils.FormulaStructure{FormulaPriorityNum:1, Formula:"[AcikHesaplarinMinAcilisTar] != '' && [AcikHesaplarinMinAcilisTar] <= ADD([Gun_Tarihi], -120)", Result:"0"}, utils.FormulaStructure{FormulaPriorityNum:2, Formula:"[KapaliHesaplarinMaxKapanisTar] == ''", Result:"[Yeni_Mus_Ek_Faiz]"}, utils.FormulaStructure{FormulaPriorityNum:3, Formula:"[KapaliHesaplarinMaxKapanisTar] < ADD([Gun_Tarihi], -5)", Result:"[Yeni_Mus_Ek_Faiz]"}}},
		{InteractionId:4, PriorityOrderNum:4, InteractionType:9, EffectedCriterionId:"4", InteractionBehaviourType:3, FormulaList:[]utils.FormulaStructure{utils.FormulaStructure{FormulaPriorityNum:1, Formula:"[Gun_Sayisi] > 400 && [Doviz] != 88", Result:"0"}, utils.FormulaStructure{FormulaPriorityNum:2, Formula:"1>0", Result:"[Lokasyon_Ek_Faiz]"}}},
		{InteractionId:5, PriorityOrderNum:99, InteractionType:4, EffectedCriterionId:"2", InteractionBehaviourType:3, FormulaList:[]utils.FormulaStructure{utils.FormulaStructure{FormulaPriorityNum:1, Formula:"([Tutar] < 5000 && [Doviz] == 88) || [Gun_Sayisi] < 30", Result:"0"}}},
		{InteractionId:6, PriorityOrderNum:99, InteractionType:4, EffectedCriterionId:"3", InteractionBehaviourType:3, FormulaList:[]utils.FormulaStructure{utils.FormulaStructure{FormulaPriorityNum:1, Formula:"([Tutar] < 5000 && [Doviz] == 88) || [Gun_Sayisi] < 30", Result:"0"}}},
		{InteractionId:7, PriorityOrderNum:99, InteractionType:4, EffectedCriterionId:"4", InteractionBehaviourType:3, FormulaList:[]utils.FormulaStructure{utils.FormulaStructure{FormulaPriorityNum:1, Formula:"([Tutar] < 5000 && [Doviz] == 88) || [Gun_Sayisi] < 30", Result:"0"}}},
	}
	BKRModelVerisi := utils.ModelAdditionalInfoStructure{
		ModelId: 45,
		SegmentLimit: "U",
		ModelWeight: 100,
		MaxCriterionScoreOfModel : 100,
		MinCriterionScoreOfModel : 0,
		MaxModelSensitiveValue : 10,
		MinModelSensitiveValue : -10,
		ModelWeightDistributionTypeId : utils.DistributeType_DontDistribute,
	}

	bireysel := ScoreCalculationModules.ModelWithInteractionPointOfView{ScoringCalculationHashMap: BKRSkorHesaplamaVerisi, ModelStructureInfo: BireyselModelYapiVerisi, InteractionInfo : BKREtkilesimVerisi, ModelAdditionalInfo: BKRModelVerisi }

	sonuc, err := ScoreCalculation.CalculateScore(bireysel)
	if err == nil {
		model := sonuc["Model"].(*ScoreCalculationMainConcept.Model)
		type SkorHesaplamaOutput struct {
			DemographicalScore           float64
			SegmentScore              string
			ModelTreeDataStructure     *utils.Node
			InteractionTreeDataStructure *utils.Node
		}
		skorHesaplamaOutput := SkorHesaplamaOutput{DemographicalScore : sonuc["DemographicalScore"].(float64), SegmentScore : "YOK", ModelTreeDataStructure: model.ModelTreeDataStructure, InteractionTreeDataStructure: model.InteractionTreeDataStructure}
		mapB, _ := json.Marshal(skorHesaplamaOutput)
		fmt.Println(string(mapB))
	} else {
		fmt.Println(err)
	}

}
*/

// Main Scoring Bakış açısı
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

/*
// Vadeli önerilen Bulunacak mi
import (
	"gitlab.zfu.fintek.local/zfu/ratings-api/engine/utils"
	"gitlab.zfu.fintek.local/zfu/ratings-api/engine/ScoreCalculationModules"
	"gitlab.zfu.fintek.local/zfu/ratings-api/engine/ScoreCalculation"
	"gitlab.zfu.fintek.local/zfu/ratings-api/engine/ScoreCalculationMainConcept"
	"encoding/json"
	"fmt"
)

func main() {
	scoringCalculationHashMap := map[string]interface{}{
		"[Musteri_Personel_Mi]":           "H",
		"[Musteri_Istirak_Personeli_Mi]":  "H",
		"[AcikHesaplarinMinAcilisTar]":    "",
		"[KapaliHesaplarinMaxKapanisTar]": "",
		"[Gun_Tarihi]":                    "2017-04-11",
		"[Doviz]":                         "TRY",
		"[Tutar]":                         "15000",
		"[Musteri_Kamu_Musterisi_Mi]":     "H",
		"[Segment]":                       "1006",
		"[Vade_Suresi]":                   "30",
		"[Mevcut_Tabela]":                 "7",
		"[Sube_Yetki_Limiti]":             "5",
	}

	modelStructureInfo := []utils.ModelStructure{
		{Name: "Önerilen Bulunacak Mı ?", CriterionId: "1", Structure: "...1", Weight: 100, CriterionScaleLimit: "U", CriterionFormula: utils.FormulaStructure{}},
		{Name: "Personel Kontrolü", CriterionId: "2", Structure: "...1...1", Weight: 100, CriterionScaleLimit: "U", CriterionFormula: utils.FormulaStructure{}},
		{Name: "İştirak Personeli Kontrolü", CriterionId: "3", Structure: "...1...2", Weight: 100, CriterionScaleLimit: "U", CriterionFormula: utils.FormulaStructure{}},
		{Name: "Alt-Üst Limit Kontrolü", CriterionId: "4", Structure: "...1...3", Weight: 100, CriterionScaleLimit: "U", CriterionFormula: utils.FormulaStructure{}},
		{Name: "Segment Kontrolu", CriterionId: "5", Structure: "...1...4", Weight: 100, CriterionScaleLimit: "U", CriterionFormula: utils.FormulaStructure{}},
		{Name: "Kamu Müşterisi Kontrolü", CriterionId: "6", Structure: "...1...5", Weight: 100, CriterionScaleLimit: "U", CriterionFormula: utils.FormulaStructure{}},
		{Name: "Vade Süresi Kontrolü", CriterionId: "7", Structure: "...1...6", Weight: 100, CriterionScaleLimit: "U", CriterionFormula: utils.FormulaStructure{}},
		{Name: "Yeni Müşteri Kontrolü", CriterionId: "8", Structure: "...1...7", Weight: 100, CriterionScaleLimit: "U", CriterionFormula: utils.FormulaStructure{}},
		{Name: "Model Çalışma Kontrolü", CriterionId: "9", Structure: "...1...8", Weight: 100, CriterionScaleLimit: "U", CriterionFormula: utils.FormulaStructure{}},
		{Name: "Şube Yetkisi Tabela İle Aynı Mı Kontrolü", CriterionId: "10", Structure: "...1...9", Weight: 100, CriterionScaleLimit: "U", CriterionFormula: utils.FormulaStructure{}},
	}

	interactionInfo := []utils.InteractionStructure{
		{InteractionId: 1, PriorityOrderNum: 1, InteractionType: 9, EffectedCriterionId: "2", InteractionBehaviourType: 3, FormulaList: []utils.FormulaStructure{utils.FormulaStructure{FormulaPriorityNum: 1, Formula: "[Musteri_Personel_Mi] == 'E'", Result: "-1"}}},
		{InteractionId: 2, PriorityOrderNum: 2, InteractionType: 9, EffectedCriterionId: "3", InteractionBehaviourType: 3, FormulaList: []utils.FormulaStructure{utils.FormulaStructure{FormulaPriorityNum: 1, Formula: "[Musteri_Istirak_Personeli_Mi] == 'E'", Result: "-1"}}},
		{InteractionId: 3, PriorityOrderNum: 3, InteractionType: 9, EffectedCriterionId: "4", InteractionBehaviourType: 3, FormulaList: []utils.FormulaStructure{utils.FormulaStructure{FormulaPriorityNum: 1, Formula: "[Doviz] == 'TRY' && ([Tutar] < 10000 || [Tutar] > 20000) ", Result: "-1"}, utils.FormulaStructure{FormulaPriorityNum: 2, Formula: "[Doviz] == 'USD' && ([Tutar] < 500 || [Tutar] > 1000) ", Result: "-1"}, utils.FormulaStructure{FormulaPriorityNum: 3, Formula: "[Doviz] == 'EUR' && ([Tutar] < 1000 || [Tutar] > 2000) ", Result: "-1"}}},
		{InteractionId: 4, PriorityOrderNum: 4, InteractionType: 9, EffectedCriterionId: "5", InteractionBehaviourType: 3, FormulaList: []utils.FormulaStructure{utils.FormulaStructure{FormulaPriorityNum: 1, Formula: "[Segment] != 1009 && [Segment] != 10091 && [Segment] != 10092 && [Segment] != 10093 && [Segment] != 1006 && [Segment] != 1007 && [Segment] != 1008", Result: "-1"}}},
		{InteractionId: 5, PriorityOrderNum: 5, InteractionType: 9, EffectedCriterionId: "6", InteractionBehaviourType: 3, FormulaList: []utils.FormulaStructure{utils.FormulaStructure{FormulaPriorityNum: 1, Formula: "[Musteri_Kamu_Musterisi_Mi] == 'E'", Result: "-1"}}},
		{InteractionId: 6, PriorityOrderNum: 6, InteractionType: 9, EffectedCriterionId: "7", InteractionBehaviourType: 3, FormulaList: []utils.FormulaStructure{utils.FormulaStructure{FormulaPriorityNum: 1, Formula: "[Vade_Suresi] <= 7", Result: "-1"}}},
		{InteractionId:7, PriorityOrderNum:7, InteractionType:9, EffectedCriterionId:"8", InteractionBehaviourType:3, FormulaList:[]utils.FormulaStructure{utils.FormulaStructure{FormulaPriorityNum:1, Formula:"[AcikHesaplarinMinAcilisTar] != '' && [AcikHesaplarinMinAcilisTar] <= ADD([Gun_Tarihi], -120)", Result:"0"}, utils.FormulaStructure{FormulaPriorityNum:2, Formula:"[KapaliHesaplarinMaxKapanisTar] == ''", Result:"-1"}, utils.FormulaStructure{FormulaPriorityNum:3, Formula:"[KapaliHesaplarinMaxKapanisTar] < ADD([Gun_Tarihi], -5)", Result:"-1"}}},
		{InteractionId: 8, PriorityOrderNum: 8, InteractionType: 9, EffectedCriterionId: "9", InteractionBehaviourType: 3, FormulaList: []utils.FormulaStructure{utils.FormulaStructure{FormulaPriorityNum: 1, Formula: "false", Result: "-1"}}},
		{InteractionId: 9, PriorityOrderNum: 9, InteractionType: 9, EffectedCriterionId: "10", InteractionBehaviourType: 3, FormulaList: []utils.FormulaStructure{utils.FormulaStructure{FormulaPriorityNum: 1, Formula: "[Mevcut_Tabela] == [Sube_Yetki_Limiti]", Result: "-1"}}},
	}
	modelAdditionalInfo := utils.ModelAdditionalInfoStructure{
		ModelId:                        45,
		SegmentLimit:                   "U",
		ModelWeight:                    100,
		MaxCriterionScoreOfModel:       100,
		MinCriterionScoreOfModel:       -100,
		MaxModelSensitiveValue:         10,
		MinModelSensitiveValue:         -10,
		ModelWeightDistributionTypeId: utils.DistributeType_DontDistribute,
	}

	scoring := ScoreCalculationModules.ModelWithInteractionPointOfView{ScoringCalculationHashMap: scoringCalculationHashMap, ModelStructureInfo: modelStructureInfo, InteractionInfo: interactionInfo, ModelAdditionalInfo: modelAdditionalInfo }

	sonuc, err := ScoreCalculation.CalculateScore(scoring)
	if err == nil {
		model := sonuc["Model"].(*ScoreCalculationMainConcept.Model)
		type SkorHesaplamaOutput struct {
			DemographicalScore           float64
			SegmentScore              string
			ModelTreeDataStructure     *utils.Node
			InteractionTreeDataStructure *utils.Node
		}
		skorHesaplamaOutput := SkorHesaplamaOutput{DemographicalScore: sonuc["DemographicalScore"].(float64), SegmentScore: "YOK", ModelTreeDataStructure: model.CreateModelLogTree(), InteractionTreeDataStructure: model.CreateInteractionLogTree()}
		mapB, _ := json.Marshal(skorHesaplamaOutput)
		fmt.Println(string(mapB))
	} else {
		fmt.Println(err)
	}

}
*/
/*
// Vadeli önerilen
import (
	"gitlab.zfu.fintek.local/zfu/ratings-api/engine/utils"
	"gitlab.zfu.fintek.local/zfu/ratings-api/engine/ScoreCalculationModules"
	"gitlab.zfu.fintek.local/zfu/ratings-api/engine/ScoreCalculation"
	"gitlab.zfu.fintek.local/zfu/ratings-api/engine/ScoreCalculationMainConcept"
	"encoding/json"
	"fmt"
)

func main() {
	scoringCalculationHashMap := map[string]interface{}{
		"[Lokasyon_Onerilen_Ek_Faizi]": "1.5",
		"[Fiyatlama_Notu]":             "7.5",
		"[Mevcut_Tabela]":              "7",
		"[Doviz]":                      "TRY",
		"[Max_Faiz]":                   "5",
		"[Donem_Sonu_Ek_Faiz]":         "2",
		"[Lokasyon_Ek_Faiz]":           "3",
		"[Musteri_Vadeli_Bakiyesi]":    "1000001",
		"[Tutar]":                      "100",
	}

	modelStructureInfo := []utils.ModelStructure{
		{Name: "Önerilen Faiz Oranı", CriterionId: "1", Structure: "...1", Weight: 100, CriterionScaleLimit: "U", CriterionFormula: utils.FormulaStructure{}},
		{Name: "Minimum Önerilebilir Faiz Oranı", CriterionId: "2", Structure: "...2", Weight: 0, CriterionScaleLimit: "U", CriterionFormula: utils.FormulaStructure{}},
		{Name: "Max Önerilebilir Faiz Oranı", CriterionId: "3", Structure: "...3", Weight: 0, CriterionScaleLimit: "U", CriterionFormula: utils.FormulaStructure{}},

		{Name: "Min Önerilebilir TRY Faiz Oranı", CriterionId: "4", Structure: "...4", Weight: 0, CriterionScaleLimit: "U", CriterionFormula: utils.FormulaStructure{}},
		{Name: "Min Önerilebilir EUR Faiz Oranı", CriterionId: "5", Structure: "...5", Weight: 0, CriterionScaleLimit: "U", CriterionFormula: utils.FormulaStructure{}},
		{Name: "Min Önerilebilir USD Faiz Oranı", CriterionId: "6", Structure: "...6", Weight: 0, CriterionScaleLimit: "U", CriterionFormula: utils.FormulaStructure{}},

	}

	interactionInfo := []utils.InteractionStructure{
		{InteractionId: 1, PriorityOrderNum: 4, InteractionType: 9, EffectedCriterionId: "1", InteractionBehaviourType: 3, FormulaList: []utils.FormulaStructure{utils.FormulaStructure{FormulaPriorityNum: 1, Formula: "true", Result: "[Fiyatlama_Notu] + [Lokasyon_Onerilen_Ek_Faizi]"}}},
		{InteractionId: 2, PriorityOrderNum: 5, InteractionType: 9, EffectedCriterionId: "2", InteractionBehaviourType: 3, FormulaList: []utils.FormulaStructure{utils.FormulaStructure{FormulaPriorityNum: 1, Formula: "true", Result: "[Mevcut_Tabela] + ([Doviz] == 'TRY' ?  [P_4] : ([Doviz] == 'USD' ?  [P_5] : ([Doviz] == 'EUR' ?  [P_6] : 0)))"}}},
		{InteractionId: 3, PriorityOrderNum: 6, InteractionType: 9, EffectedCriterionId: "3", InteractionBehaviourType: 3, FormulaList: []utils.FormulaStructure{utils.FormulaStructure{FormulaPriorityNum: 1, Formula: "true", Result: "[Max_Faiz] + [Donem_Sonu_Ek_Faiz] + [Lokasyon_Ek_Faiz]"}}},

		{InteractionId: 4, PriorityOrderNum: 1, InteractionType: 9, EffectedCriterionId: "4", InteractionBehaviourType: 3, FormulaList: []utils.FormulaStructure{
			utils.FormulaStructure{FormulaPriorityNum: 1, Formula: "([Musteri_Vadeli_Bakiyesi] + [Tutar]) >= 1000 && ([Musteri_Vadeli_Bakiyesi] + [Tutar]) < 5000", Result: "0"},
			utils.FormulaStructure{FormulaPriorityNum: 2, Formula: "([Musteri_Vadeli_Bakiyesi] + [Tutar]) >= 5000 && ([Musteri_Vadeli_Bakiyesi] + [Tutar]) < 50000", Result: "2.5"},
			utils.FormulaStructure{FormulaPriorityNum: 3, Formula: "([Musteri_Vadeli_Bakiyesi] + [Tutar]) >= 50000 && ([Musteri_Vadeli_Bakiyesi] + [Tutar]) < 100000", Result: "3"},
			utils.FormulaStructure{FormulaPriorityNum: 4, Formula: "([Musteri_Vadeli_Bakiyesi] + [Tutar]) >= 100000 && ([Musteri_Vadeli_Bakiyesi] + [Tutar]) < 250000", Result: "3.5"},
			utils.FormulaStructure{FormulaPriorityNum: 5, Formula: "([Musteri_Vadeli_Bakiyesi] + [Tutar]) >= 250000 && ([Musteri_Vadeli_Bakiyesi] + [Tutar]) < 500000", Result: "4"},
			utils.FormulaStructure{FormulaPriorityNum: 6, Formula: "([Musteri_Vadeli_Bakiyesi] + [Tutar]) >= 500000 && ([Musteri_Vadeli_Bakiyesi] + [Tutar]) < 1000000", Result: "4.5"},
			utils.FormulaStructure{FormulaPriorityNum: 7, Formula: "([Musteri_Vadeli_Bakiyesi] + [Tutar]) >= 1000000 && ([Musteri_Vadeli_Bakiyesi] + [Tutar]) < 20000000", Result: "5"},
		}},

		{InteractionId: 5, PriorityOrderNum: 2, InteractionType: 9, EffectedCriterionId: "5", InteractionBehaviourType: 3, FormulaList: []utils.FormulaStructure{
			utils.FormulaStructure{FormulaPriorityNum: 1, Formula: "([Musteri_Vadeli_Bakiyesi] + [Tutar]) >= 1000 && ([Musteri_Vadeli_Bakiyesi] + [Tutar]) < 5000", Result: "0"},
			utils.FormulaStructure{FormulaPriorityNum: 2, Formula: "([Musteri_Vadeli_Bakiyesi] + [Tutar]) >= 5000 && ([Musteri_Vadeli_Bakiyesi] + [Tutar]) < 50000", Result: "0.25"},
			utils.FormulaStructure{FormulaPriorityNum: 3, Formula: "([Musteri_Vadeli_Bakiyesi] + [Tutar]) >= 50000 && ([Musteri_Vadeli_Bakiyesi] + [Tutar]) < 100000", Result: "0.50"},
		}},
		{InteractionId: 6, PriorityOrderNum: 3, InteractionType: 9, EffectedCriterionId: "6", InteractionBehaviourType: 3, FormulaList: []utils.FormulaStructure{
			utils.FormulaStructure{FormulaPriorityNum: 1, Formula: "([Musteri_Vadeli_Bakiyesi] + [Tutar]) >= 1000 && ([Musteri_Vadeli_Bakiyesi] + [Tutar]) < 5000", Result: "0"},
			utils.FormulaStructure{FormulaPriorityNum: 2, Formula: "([Musteri_Vadeli_Bakiyesi] + [Tutar]) >= 5000 && ([Musteri_Vadeli_Bakiyesi] + [Tutar]) < 50000", Result: "0.25"},
			utils.FormulaStructure{FormulaPriorityNum: 3, Formula: "([Musteri_Vadeli_Bakiyesi] + [Tutar]) >= 50000 && ([Musteri_Vadeli_Bakiyesi] + [Tutar]) < 100000", Result: "0.50"},
		}},

		{InteractionId: 7, PriorityOrderNum: 7, InteractionType: 4, EffectedCriterionId: "1", InteractionBehaviourType: 3, FormulaList: []utils.FormulaStructure{utils.FormulaStructure{FormulaPriorityNum: 1, Formula: "[P_1] < [P_2]", Result: "[P_2]"}}},
		{InteractionId: 8, PriorityOrderNum: 8, InteractionType: 4, EffectedCriterionId: "1", InteractionBehaviourType: 3, FormulaList: []utils.FormulaStructure{utils.FormulaStructure{FormulaPriorityNum: 1, Formula: "[P_3] < [P_1]", Result: "[P_3]"}}},

	}

	modelAdditionalInfo := utils.ModelAdditionalInfoStructure{
		ModelId:                        45,
		SegmentLimit:                   "U",
		ModelWeight:                    100,
		MaxCriterionScoreOfModel:       100,
		MinCriterionScoreOfModel:       -100,
		MaxModelSensitiveValue:         10,
		MinModelSensitiveValue:         -10,
		ModelWeightDistributionTypeId: utils.DistributeType_DontDistribute,
	}

	scoring := ScoreCalculationModules.ModelWithInteractionPointOfView{ScoringCalculationHashMap: scoringCalculationHashMap, ModelStructureInfo: modelStructureInfo, InteractionInfo: interactionInfo, ModelAdditionalInfo: modelAdditionalInfo }

	sonuc, err := ScoreCalculation.CalculateScore(scoring)
	if err == nil {
		model := sonuc["Model"].(*ScoreCalculationMainConcept.Model)
		type SkorHesaplamaOutput struct {
			DemographicalScore           float64
			SegmentScore              string
			ModelTreeDataStructure     *utils.Node
			InteractionTreeDataStructure *utils.Node
		}
		skorHesaplamaOutput := SkorHesaplamaOutput{DemographicalScore: sonuc["DemographicalScore"].(float64), SegmentScore: "YOK", ModelTreeDataStructure: model.CreateModelLogTree(), InteractionTreeDataStructure: model.CreateInteractionLogTree()}
		mapB, _ := json.Marshal(skorHesaplamaOutput)
		fmt.Println(string(mapB))
	} else {
		fmt.Println(err)
	}

}
*/
