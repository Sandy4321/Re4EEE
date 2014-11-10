package Website

type Basis struct {
	Name    string
	Version string
	Lang    string
	Session string
}

type PagePassword struct {
	Basis        Basis
	TextPassword string
}

type PageStart struct {
	Basis           Basis
	TextWelcome     string
	TextStartButton string
}

type PageQuestion struct {
	Basis             Basis
	NoQuestion        string
	NoQuestions       string
	Progress          string
	TextQuestion      string
	TextQuestionTopic string
	TextQuestionBody  string
	TextButton1       string
	TextButton2       string
	TextButton3       string
	TextButton4       string
	TextButton5       string
	Button1Status     string
	Button2Status     string
	Button3Status     string
	Button4Status     string
	Button5Status     string
	Button1Data       string
	Button2Data       string
	Button3Data       string
	Button4Data       string
	Button5Data       string
}