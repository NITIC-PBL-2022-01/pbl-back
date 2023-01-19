package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type Token struct {
  AccessToken string `json:"access_token"`
  RefreshToken string `json:"refresh_token"`
  ExpireIn int `json:"expire_in"`
}

var API_KEY string

func Signup(email string, password string) (Token, error) {
  type RequestBody struct {
    Email string `json:"email"`
    Password string `json:"password"`
    ReturnSecureToken bool `json:"returnSecureToken"`
  }

  data := RequestBody {
    Email: email,
    Password: password,
    ReturnSecureToken: true,
  }

  j, err := json.Marshal(data)
  if err != nil {
    return Token{}, err
  }

  req, err := http.NewRequest("POST", "https://identitytoolkit.googleapis.com/v1/accounts:signUp", bytes.NewBuffer(j))
  if err != nil {
    return Token{}, err
  }

  q := req.URL.Query()
  q.Add("key", API_KEY)
  req.URL.RawQuery = q.Encode()


  client := &http.Client{}
  res, err := client.Do(req)
  if err != nil {
    return Token{}, err
  }
  defer res.Body.Close()

  b, err := io.ReadAll(res.Body)
  if err != nil {
    return Token{}, err
  }
  fmt.Println(string(b))

  type ResponseBody struct {
    IDToken string `json:"idToken"`
    Email string `json:"email"`
    RefreshToken string `json:"refreshToken"`
    ExpireIn string `json:"expiresIn"`
    LocalID string `json:"localId"`
  }

  body := ResponseBody{}

  err = json.Unmarshal(b, &body)
  if err != nil {
    return Token{}, err
  }

  expireIn, err := strconv.Atoi(body.ExpireIn)
  if err != nil {
    return Token{}, err
  }

  return Token {
    AccessToken: body.IDToken,
    RefreshToken: body.RefreshToken,
    ExpireIn: expireIn,
  }, nil
}

func SignIn(email string, password string) (Token, error) {
  type RequestBody struct {
    Email string `json:"email"`
    Password string `json:"password"`
    ReturnSecureToken bool `json:"returnSecureToken"`
  }

  data := RequestBody {
    Email: email,
    Password: password,
    ReturnSecureToken: true,
  }

  j, err := json.Marshal(data)
  if err != nil {
    return Token{}, err
  }

  req, err := http.NewRequest("POST", "https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword", bytes.NewBuffer(j))
  if err != nil {
    return Token{}, err
  }

  q := req.URL.Query()
  q.Add("key", API_KEY)
  req.URL.RawQuery = q.Encode()


  client := &http.Client{}
  res, err := client.Do(req)
  if err != nil {
    return Token{}, err
  }
  defer res.Body.Close()

  b, err := io.ReadAll(res.Body)
  if err != nil {
    return Token{}, err
  }
  fmt.Println(string(b))

  type ResponseBody struct {
    IDToken string `json:"idToken"`
    Email string `json:"email"`
    RefreshToken string `json:"refreshToken"`
    ExpireIn string `json:"expiresIn"`
    LocalID string `json:"localId"`
  }

  body := ResponseBody{}

  err = json.Unmarshal(b, &body)
  if err != nil {
    return Token{}, err
  }

  expireIn, err := strconv.Atoi(body.ExpireIn)
  if err != nil {
    return Token{}, err
  }

  return Token {
    AccessToken: body.IDToken,
    RefreshToken: body.RefreshToken,
    ExpireIn: expireIn,
  }, nil
}
