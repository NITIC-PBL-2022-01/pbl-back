package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func GetUserEmail(token string) (string, error) {
  type RequestBody struct {
    IDToken string `json:"idToken"`
  }

  data := RequestBody {
    IDToken: token,
  }

  j, err := json.Marshal(data)
  if err != nil {
    return "", err
  }

  req, err := http.NewRequest("POST", "https://identitytoolkit.googleapis.com/v1/accounts:lookup", bytes.NewBuffer(j))
  if err != nil {
    return "", err
  }

  q := req.URL.Query()
  q.Add("key", API_KEY)
  req.URL.RawQuery = q.Encode()


  client := &http.Client{}
  res, err := client.Do(req)
  if err != nil {
    return "", err
  }
  defer res.Body.Close()

  b, err := io.ReadAll(res.Body)
  if err != nil {
    return "", err
  }

  fmt.Println(string(b))

  type ResponseBody struct {
    Users []struct {
      Email string `json:"email"`
    } `json:"users"`
  }

  body := ResponseBody{}

  err = json.Unmarshal(b, &body)
  if err != nil {
    return "", err
  }

  return body.Users[0].Email, nil
}

func Refresh(token string) (Token, error) {
  type RequestBody struct {
    GrantType string `json:"grant_type"`
    RefreshToken string `json:"refresh_token"`
  }

  data := RequestBody {
    GrantType: "refresh_token",
    RefreshToken: token,
  }

  j, err := json.Marshal(data)
  if err != nil {
    return Token{}, err
  }

  req, err := http.NewRequest("POST", "https://securetoken.googleapis.com/v1/token", bytes.NewBuffer(j))
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
    IDToken string `json:"id_token"`
    RefreshToken string `json:"refresh_token"`
    ExpireIn string `json:"expires_in"`
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
