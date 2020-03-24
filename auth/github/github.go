package github

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type User struct {
	Login                   string      `json:"login"`
	ID                      int         `json:"id"`
	NodeID                  string      `json:"node_id"`
	AvatarURL               string      `json:"avatar_url"`
	GravatarID              string      `json:"gravatar_id"`
	URL                     string      `json:"url"`
	HTMLURL                 string      `json:"html_url"`
	FollowersURL            string      `json:"followers_url"`
	FollowingURL            string      `json:"following_url"`
	GistsURL                string      `json:"gists_url"`
	StarredURL              string      `json:"starred_url"`
	SubscriptionsURL        string      `json:"subscriptions_url"`
	OrganizationsURL        string      `json:"organizations_url"`
	ReposURL                string      `json:"repos_url"`
	EventsURL               string      `json:"events_url"`
	ReceivedEventsURL       string      `json:"received_events_url"`
	Type                    string      `json:"type"`
	SiteAdmin               bool        `json:"site_admin"`
	Name                    string      `json:"name"`
	Company                 interface{} `json:"company"`
	Blog                    string      `json:"blog"`
	Location                interface{} `json:"location"`
	Email                   string      `json:"email"`
	Hireable                interface{} `json:"hireable"`
	Bio                     interface{} `json:"bio"`
	PublicRepos             int         `json:"public_repos"`
	PublicGists             int         `json:"public_gists"`
	Followers               int         `json:"followers"`
	Following               int         `json:"following"`
	CreatedAt               time.Time   `json:"created_at"`
	UpdatedAt               time.Time   `json:"updated_at"`
	PrivateGists            int         `json:"private_gists"`
	TotalPrivateRepos       int         `json:"total_private_repos"`
	OwnedPrivateRepos       int         `json:"owned_private_repos"`
	DiskUsage               int         `json:"disk_usage"`
	Collaborators           int         `json:"collaborators"`
	TwoFactorAuthentication bool        `json:"two_factor_authentication"`
	Plan                    struct {
		Name          string `json:"name"`
		Space         int    `json:"space"`
		Collaborators int    `json:"collaborators"`
		PrivateRepos  int    `json:"private_repos"`
	} `json:"plan"`
	AccessToken string
}

type access struct {
	Access_token string `json:'access_token'`
}

const (
	PATH_ACCESS = "https://github.com/login/oauth/access_token"
	PATH_USER   = "https://api.github.com/user"
)

// 调用 Github auth api
type Github struct {
	ClientId     string
	ClientSecret string
}

func (g *Github) GetUser(code string) (*User, error) {

	params := []string{
		"client_id=" + g.ClientId,
		"client_secret=" + g.ClientSecret,
		"code=" + code,
	}

	_url := PATH_ACCESS + "?" + strings.Join(params, "&")

	req, err := http.NewRequest("POST", _url, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("accept", "application/json")

	client := &http.Client{}

	res, err := client.Do(req)

	defer res.Body.Close()

	if err != nil {
		return nil, errors.New("请稍后再试")
	}

	b, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
	}

	data := &access{}

	json.Unmarshal(b, &data)

	if data.Access_token == "" {
		return nil, errors.New("已过期")
	}

	req, _ = http.NewRequest("GET", PATH_USER, nil)
	req.Header.Add("Authorization", "token "+data.Access_token)
	req.Header.Add("accept", "application/json")

	client = &http.Client{}
	res, err = client.Do(req)

	defer res.Body.Close()

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("请稍后再试")
	}

	b, _ = ioutil.ReadAll(res.Body)

	user := &User{}

	json.Unmarshal(b, &user)

	user.AccessToken = data.Access_token

	return user, nil
}
