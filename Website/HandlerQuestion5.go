package Website

import (
	"fmt"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/Templates"
	"github.com/SommerEngineering/Ocean/Tools"
	"net/http"
	"strings"
)

func HandlerQuestion5(response http.ResponseWriter, request *http.Request) {
	noQuestion := 5
	readSession := request.FormValue(`session`)
	if readSession == `` {
		defer http.Redirect(response, request, "/", 307)
		Log.LogFull(senderName, LM.CategoryAPP, LM.LevelSECURITY, LM.SeverityMiddle, LM.ImpactNone, LM.MessageNameUSER, `A request without session.`)
		return
	}

	lang := Tools.GetRequestLanguage(request)[0]
	data := PageQuestion{}
	data.Basis.Version = VERSION
	data.Progress = 5
	data.Basis.Lang = lang.Language
	data.Basis.Session = readSession

	data.Button1Status = BUTTON_SHOW
	data.Button2Status = BUTTON_SHOW
	data.Button3Status = BUTTON_SHOW
	data.Button4Status = BUTTON_HIDDEN
	data.Button5Status = BUTTON_HIDDEN
	data.ButtonBackStatus = BUTTON_SHOW
	data.ButtonInfoStatus = BUTTON_SHOW

	data.Button1Data = `1`
	data.Button2Data = `0`
	data.Button3Data = `*`
	data.Button4Data = ``
	data.Button5Data = ``

	data.NoQuestion = fmt.Sprintf(`%d`, noQuestion)
	data.PreNoQuestion = fmt.Sprintf(`%d`, noQuestion-1)
	data.NoQuestions = totalQuestions

	if strings.Contains(lang.Language, `de`) {
		data.Basis.Name = NAME_DE
		data.Basis.Logo = LOGO_DE
		data.TextButton1 = `Ja`
		data.TextButton2 = `Nein`
		data.TextButton3 = `Enthaltung`
		data.TextButton4 = ``
		data.TextButton5 = ``
		data.TextBackButton = `Vorherige Frage`
		data.TextImportant = `Diese Aussage ist mir besonderst wichtig`
		data.TextQuestion = `Frage`
		data.TextQuestionTopic = `Asynchrone Interaktion`
		data.TextQuestionBody = `Soll das E-Learning-Format eine zeitlich asynchrone Interaktion ermöglichen?`
		data.QuestionInfoHeader = `Zusätzliche Hinweise`
		data.QuestionInfoClose = `Schließen`
		data.QuestionInfoText = `Dies können z.B. Diskussionsforen und Wikis sein, in denen die Studierenden zeitlich unabhängig voneinander interagieren.`
	} else {
		data.Basis.Name = NAME_EN
		data.Basis.Logo = LOGO_UK
		data.TextButton1 = `Yes`
		data.TextButton2 = `No`
		data.TextButton3 = `Skip question`
		data.TextButton4 = ``
		data.TextButton5 = ``
		data.TextBackButton = `Previous question`
		data.TextImportant = `This statement is important for me`
		data.TextQuestion = `Question`
		data.TextQuestionTopic = `Asynchronous Interaction`
		data.TextQuestionBody = `Should the e-learning format provide a functionality for asynchronous interaction?`
		data.QuestionInfoHeader = `Additional Information`
		data.QuestionInfoClose = `Close`
		data.QuestionInfoText = `Examples are e.g. discussion boards and Wikis, in which the work occurs time-shifted.`
	}

	Tools.SendChosenLanguage(response, lang)
	Templates.ProcessHTML(`question`, response, data)
}
