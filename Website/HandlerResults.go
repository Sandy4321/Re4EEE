package Website

import (
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/Templates"
	"github.com/SommerEngineering/Ocean/Tools"
	"github.com/SommerEngineering/Re4EEE/Algorithm"
	"github.com/SommerEngineering/Re4EEE/DB"
	"github.com/SommerEngineering/Re4EEE/DB/Scheme"
	"github.com/SommerEngineering/Re4EEE/XML"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func HandlerResults(response http.ResponseWriter, request *http.Request) {
	session := request.FormValue(`session`)
	amountText := request.FormValue(`amount`)
	lang := Tools.GetRequestLanguage(request)[0]
	answers := DB.LoadAnswers(session)
	groups := XML.GetData()
	amountValue := -1
	resultSet := Scheme.Recommendation{}

	if len(session) != 36 {
		Log.LogFull(senderName, LM.CategoryAPP, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameSTATE, `Session's length was not valid!`, session)
		response.WriteHeader(http.StatusNotFound)
		return
	}

	if amountText != `` && len(amountText) > 2 {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	if !DB.CheckRecommendation(session) {
		resultSet.ProductGroups, resultSet.Influence = Algorithm.ExecuteAnswers(answers)
		resultSet.CreateTimeUTC = time.Now().UTC()
		resultSet.Session = session
		resultSet.SchemeVersion = Scheme.CURRENT_VERSION
		DB.StoreRecommendation(resultSet)

	} else {
		resultSet = DB.LoadRecommendation(session)
	}

	if value, errConv := strconv.Atoi(amountText); errConv != nil {
		Log.LogFull(senderName, LM.CategoryAPP, LM.LevelERROR, LM.SeverityMiddle, LM.ImpactNone, LM.MessageNameREQUEST, `Cannot read the amount value!`, amountText, errConv.Error())
	} else {
		amountValue = value
	}

	if amountValue >= 1 {
		resultSet.ProductGroups = resultSet.ProductGroups[0:amountValue]
	}

	data := PageResults{}
	data.Basis.Version = VERSION
	data.Basis.Lang = lang.Language
	data.Basis.Session = session
	data.Groups = groups.ProductsCollection.Products
	data.Questions = groups.QuestionsCollection.Questions
	data.Recommendation = resultSet
	data.AmountCurrent = amountValue
	data.Strings = groups.ResultStrings

	if strings.Contains(lang.Language, `de`) {

		if amountValue > 0 {
			data.TextAllGroups = `Alle Gruppen anzeigen`
			data.AmountToggle = -1
		} else {
			data.TextAllGroups = `Top 6 anzeigen`
			data.AmountToggle = 6
		}

		data.Basis.Name = NAME_DE
		data.Basis.Logo = LOGO_DE
		data.LangPos = 0
	} else {

		if amountValue > 0 {
			data.TextAllGroups = `Show all groups`
			data.AmountToggle = -1
		} else {
			data.TextAllGroups = `Show top 6`
			data.AmountToggle = 6
		}

		data.Basis.Name = NAME_EN
		data.Basis.Logo = LOGO_UK
		data.LangPos = 1
	}

	Tools.SendChosenLanguage(response, lang)
	Templates.ProcessHTML(`results`, response, data)
}

func (data PageResults) GetProgressState(influence int) string {
	if influence > 0 {
		return ` progressitemdone`
	} else if influence < 0 {
		return ` progressitemundone`
	} else {
		return ``
	}
}

func (data PageResults) GetGroupName(xmlIndex int) string {
	return data.Groups[xmlIndex].GroupName.Names[data.LangPos].Text
}

func (data PageResults) Lang(strings []XML.String) string {
	return strings[data.LangPos].Text
}
