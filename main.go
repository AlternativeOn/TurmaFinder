package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/alternativeon/pgo"
	"github.com/andlabs/ui"
	"github.com/mattn/go-colorable"
	"github.com/rs/zerolog"
	llgg "github.com/rs/zerolog/log"
)

var mainwin *ui.Window

type UserClass []struct {
	Value  string   `json:"value"`
	Label  string   `json:"label"`
	Turmas []Turmas `json:"turmas"`
}

type Turmas struct {
	NomeTurma   string `json:"nomeTurma"`
	TurmaValida bool   `json:"turmaValida"`
	NomeSerie   string `json:"nomeSerie"`
}

func init() {
	//Configuração do logger
	logOutput := colorable.NewColorableStdout()
	llgg.Logger = llgg.Output(zerolog.ConsoleWriter{Out: logOutput})
}

func loginPage() ui.Control {
	//Cria a UI de login
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	vbox.Append(ui.NewLabel("Para começarmos, faça seu login. Use o usuário e senha do Positivo On"), false)
	vbox.Append(ui.NewVerticalSeparator(), false)

	group := ui.NewGroup("Login")
	group.SetMargined(true)
	vbox.Append(group, true)

	aboutform := ui.NewForm()
	aboutform.SetPadded(true)
	group.SetChild(aboutform)

	//aboutform.Append("Usuário", ui.NewEntry(), false)
	user := ui.NewEntry()
	user.SetText("")
	aboutform.Append("Usuário", user, false)
	pass := ui.NewPasswordEntry()
	pass.SetText("")
	aboutform.Append("Senha", pass, false)
	loginButton := ui.NewButton("Entrar")
	loginButton.OnClicked(func(*ui.Button) {
		//disable button
		loginButton.Disable()
		//get user and password
		user := user.Text()
		pass := pass.Text()
		//check if user and password are empty
		if user == "" || pass == "" {
			llgg.Error().Msg("Usuário tentou fazer login sem preencher os dados necessários.")
			//error message
			ui.MsgBoxError(mainwin, "Preencha todos os campos", "O campo de usuário e senha não podem ser vazios.")
			loginButton.Enable()
			return
		}
		//do login
		llgg.Trace().Msg("Iniciando login...")

		usrtoken, err := pgo.Login(user, pass)
		if err != nil {
			llgg.Error().Str("Detalhes", err.Error()).Msg("Erro ao fazer login!")
			ui.MsgBoxError(mainwin, "Erro ao realizar a autenticação!", err.Error())
			loginButton.Enable()
			return
		}
		llgg.Info().Msg("Login realizado com sucesso!")
		llgg.Trace().Msg("Tentando pegar as informações do usuário...")
		username, err := pgo.GetUserName(usrtoken)
		if err != nil {
			llgg.Error().Str("Erro", err.Error()).Msg("Problema ao obter o nome do usuário!")
		}
		//Corta o nome do usuário, mostrando apenas o primeiro nome
		username = strings.Split(username, " ")[0]
		llgg.Trace().Str("Nome do usuário", username).Msg("Nome obtido com sucesso!")

		_, userid, _, err := pgo.GetUserInfo(usrtoken)
		if err != nil {
			llgg.Error().Str("Erro", err.Error()).Msg("Problema ao obter o ID do usuário!")
			ui.MsgBoxError(mainwin, "Não foi possível obter sua turma.", "Por favor verifique a sua conexão com a internet.")
			os.Exit(0)
		}

		ensino, nivel, sala, err := GetClass(usrtoken, userid)

		if err != nil {
			llgg.Error().Str("Erro", err.Error()).Msg("Problema ao obter o ID do usuário!")
			ui.MsgBoxError(mainwin, "Não foi possível obter sua turma.", "Por favor verifique a sua conexão com a internet.")
			os.Exit(0)
		}

		ui.MsgBox(mainwin, "Parabéns, "+username, "Você é do "+ensino+", está na "+nivel+" e é da sala "+trimLeftChars(sala, 2))

		//Vai para a página apos o login
		mainwin.Destroy()

		ui.Quit()

	})
	aboutform.Append("", loginButton, false)

	return vbox
}

func setupUI() {
	mainwin = ui.NewWindow("Turma Finder", 640, 480, true)
	mainwin.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	ui.OnShouldQuit(func() bool {
		mainwin.Destroy()
		return true
	})

	tab := ui.NewTab()
	mainwin.SetChild(tab)
	mainwin.SetMargined(true)

	tab.Append("Inicio", loginPage())
	tab.SetMargined(0, true)

	mainwin.Show()
}

func main() {
	ui.Main(setupUI)
}

func GetClass(token string, userId string) (string, string, string, error) {
	getclassresponse, err := tokenRequest("https://apihub.positivoon.com.br/api/NivelEnsino?usuarioId="+userId, "GET", token)
	if err != nil {
		return "", "", "", err
	}
	//unmarshal response
	var userclass UserClass
	json.Unmarshal([]byte(getclassresponse), &userclass)
	return userclass[0].Label, userclass[0].Turmas[0].NomeSerie, userclass[0].Turmas[0].NomeTurma, nil

}

func tokenRequest(url string, method string, token string) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return "", errors.New("Não foi possível criar a requesição:" + err.Error())
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	res, err := client.Do(req)

	if err != nil {
		return "", errors.New("Não foi possível enviar a requisão:" + err.Error())
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", errors.New("Não foi possível ler a resposta:" + err.Error())
	}

	return string(body), nil
}

func trimLeftChars(s string, n int) string {
	//See https://stackoverflow.com/a/48798875
	m := 0
	for i := range s {
		if m >= n {
			return s[i:]
		}
		m++
	}
	return s[:0]
}
