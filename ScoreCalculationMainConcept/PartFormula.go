package ScoreCalculationMainConcept

import (
	"engine/utils"
	"errors"
	"fmt"
	"github.com/Knetic/govaluate"
	"github.com/jinzhu/now"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type PartFormulaLogStructure struct {
	FormulaOrderNum    int
	Formula            string
	FormulaResult      string
	IsFormulaExecuted  bool
	FormulaLastState   string
	VariablesOfFormula map[string]string
}

type PartFormula struct {
	FormulaOrderNum           int
	Formula                   string
	FormulaResult             string
	IsFormulaResultCalculable bool
	BelongedModel             *Model
}

func CreatePartFormula(formulaOrderNum int, formula string, isFormulaResultCalculable bool, formulaResult string, belongedMode *Model) *PartFormula {
	return &PartFormula{
		FormulaOrderNum:           formulaOrderNum,
		Formula:                   formula,
		IsFormulaResultCalculable: isFormulaResultCalculable,
		FormulaResult:             formulaResult,
		BelongedModel:             belongedMode,
	}
}

func CreatePartFormulaLogInfo(formulaOrderNum int, formula string, formulaResult string) *PartFormulaLogStructure {
	return &PartFormulaLogStructure{
		FormulaOrderNum:    formulaOrderNum,
		Formula:            formula,
		FormulaResult:      formulaResult,
		IsFormulaExecuted:  false,
		FormulaLastState:   "",
		VariablesOfFormula: map[string]string{},
	}
}

func (partFormula *PartFormula) ExecutePartFormula() (result interface{}, formulaLastState string, variablesOfFormula map[string]string, err error) {
	resultLastState := partFormula.FormulaResult
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("Formula last state : '%s' ; Formula result last state : '%s' --> %s", formulaLastState, resultLastState, err.Error()))
		}
	}()
	formulaLastState = partFormula.Formula
	variablesOfResult := map[string]string{}
	if len(partFormula.FormulaResult) > 0 && partFormula.IsFormulaResultCalculable {
		resultLastState, variablesOfResult = replaceVariableValue(partFormula.FormulaResult, partFormula.BelongedModel)
		expressionOfInteractionResult, _ := govaluate.NewEvaluableExpression(resultLastState)
		evaluatedInteractionResult, _ := expressionOfInteractionResult.Evaluate(nil)
		partFormula.FormulaResult = strconv.FormatFloat(evaluatedInteractionResult.(float64), 'f', -1, 64)
	}
	formulaLastState, variablesOfFormula = replaceVariableValue(formulaLastState, partFormula.BelongedModel)
	formulaLastState = changeFunctionValue(formulaLastState)
	formulaLastState = utils.ConvertCommaToDot(formulaLastState)
	expression, errEvalExpression := govaluate.NewEvaluableExpression(formulaLastState)

	for key, value := range variablesOfResult {
		variablesOfFormula[key] = value
	}
	if errEvalExpression == nil {
		resultEval, errEval := expression.Evaluate(nil)
		if errEval == nil {
			result = resultEval
			return result, formulaLastState, variablesOfFormula, nil
		} else {
			err = errors.New(fmt.Sprintf("Formula '%s' could not be executed --> %s", formulaLastState, err.Error()))
			return result, formulaLastState, variablesOfFormula, err
		}
	} else {
		err = errors.New(fmt.Sprintf("Formula is not valid! --> '%s' , %s", formulaLastState, errEvalExpression.Error()))
		return result, formulaLastState, variablesOfFormula, err
	}
}

func changeFunctionValue(formulaLastState string) string {
	if strings.Contains(formulaLastState, "ORT") || strings.Contains(formulaLastState, "NVL") || strings.Contains(formulaLastState, "BOS") || strings.Contains(formulaLastState, "ADD") || strings.Contains(formulaLastState, "EOQ") {
		reFunctions := regexp.MustCompile(`\w\w\w\((.*?)\)`)
		var functions []string = reFunctions.FindAllString(formulaLastState, -1)
		for _, function := range functions {
			resultOfFunction := calculateFunction(function)
			resultOfFunction = utils.ConvertCommaToDot(resultOfFunction)
			formulaLastState = strings.Replace(formulaLastState, function, resultOfFunction, -1)
			formulaLastState = strings.Replace(formulaLastState, "--", "+", -1)
		}
	}
	return formulaLastState
}

func calculateFunction(function string) string {
	value := function
	if function[0:3] == "ORT" {
		numbers := strings.Split(function[4:len(function)-1], ",")
		var total float64 = 0
		for _, number := range numbers {
			newScore, err := strconv.ParseFloat(utils.ConvertDotToComma(number), 64)
			if err == nil {
				total += newScore
			}
		}
		value = strconv.FormatFloat(total/float64(len(numbers)), 'f', -1, 64)
	} else if function[0:3] == "NVL" {
		params := strings.Split(function[4:len(function)-1], ",")
		reVariables := regexp.MustCompile(`\\[(.*?)\\]`)
		variables := reVariables.FindAllString(function, -1)
		if len(variables) > 0 {
			value = params[1]
		} else {
			value = params[0]
		}
	} else if function[0:3] == "BOS" {
		x := function[4 : len(function)-1]
		if x == "-99999.99" || x == "DEFAULT" || x == "0001-01-01" {
			value = "true"
		} else {
			value = "false"
		}
	} else if function[0:3] == "ADD" {
		params := strings.Split(function[4:len(function)-1], ",")
		ret, _ := time.Parse("2006-01-02", strings.Replace(params[0], "'", "", -1))
		i, _ := strconv.Atoi(strings.TrimSpace(params[1]))
		result := ret.AddDate(0, 0, i)
		value = "'" + result.Format("2006-01-02") + "'"

	} else if function[0:3] == "EOQ" {
		t, _ := time.Parse("2006-01-02", strings.Replace(function[4:len(function)-1], "'", "", -1))
		result := now.New(t).EndOfQuarter()
		value = "'" + result.Format("2006-01-02") + "'"
	}
	return value
}

func replaceVariableValue(formulaLastState string, belongedModel *Model) (string, map[string]string) {
	reVariables := regexp.MustCompile("\\[(.*?)\\]")
	variables := reVariables.FindAllString(formulaLastState, -1)
	variablesOfFormula := map[string]string{}
	for _, variable := range variables {
		variableValue := findVariableValue(variable, belongedModel)
		_, err := strconv.ParseFloat(variableValue, 64)
		if err != nil {
			variableValue = "'" + variableValue + "'"
		} else {
			variableValue = utils.ConvertCommaToDot(variableValue)
		}
		if variableValue != variable {
			variablesOfFormula[variable] = variableValue
		}
		formulaLastState = strings.Replace(formulaLastState, variable, variableValue, -1)
		formulaLastState = strings.Replace(formulaLastState, "--", "+", -1)
	}
	return formulaLastState, variablesOfFormula
}

func findVariableValue(variable string, belongedModel *Model) string {
	switch {
	case variable[0:3] == "[P_":
		{
			criterionId := variable[3 : len(variable)-1]
			result := belongedModel.FindModelCriterions()[criterionId].CriterionScore
			return strconv.FormatFloat(result, 'f', -1, 64)
		}
	case variable[0:3] == "[D_":
		{
			if belongedModel.ScoringCalculationHashMap[variable[1:len(variable)-1]] != nil {
				return belongedModel.ScoScoringCalculationHashMap[variable[1:len(variable)-1]].(string)
			} else {
				return variable
			}
		}
	default:
		return belongedModel.GetInputValue(variable)
	}
}

type PartFormulaList []*PartFormula

func (s PartFormulaList) Len() int {
	return len(s)
}
func (s PartFormulaList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s PartFormulaList) Less(i, j int) bool {
	return s[i].FormulaOrderNum < s[j].FormulaOrderNum
}
