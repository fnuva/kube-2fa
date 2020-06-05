package tools

import (
	"encoding/json"
	"errors"
	duoapi "github.com/duosecurity/duo_api_golang"
	"github.com/spf13/viper"
	"kube-2fa/internal/executor"
	"net/url"
	"os"
)

type Duo struct {
	DuoDto DuoDto
	api    *duoapi.DuoApi
}

type DuoResult struct {
	Response DuoResponse
	Stat     string
}

type DuoResponse struct {
	Txid string
}

func (dr *DuoResult) Result() string {
	return dr.Response.Txid
}

func (d *Duo) Init() error {

	duoConfig := DuoDto{}
	if err := viper.UnmarshalKey("mfa_config.duo", &duoConfig); err != nil {
		return err
	}
	d.DuoDto = duoConfig
	d.api = duoapi.NewDuoApi(
		d.DuoDto.Ikey,
		d.DuoDto.Skey,
		d.DuoDto.Host,
		d.DuoDto.UserAgent)
	if !fileExists(FileName) {
		return errors.New("file " + FileName + " does not exist")
	}
	return nil

}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func (d *Duo) Run() (MfaResult, error) {
	if len(Code) > 0 {
		return d.push_with_mfa(Code)

	} else {
		return d.push()

	}
}

func (d *Duo) push() (MfaResult, error) {

	params := url.Values{}
	params["username"] = []string{d.DuoDto.UserName}
	params["factor"] = []string{"push"}
	params["device"] = []string{"auto"}
	params["async"] = []string{"1"}
	_, payload, err := d.api.SignedCall("POST", "/auth/v2/auth", params, duoapi.UseTimeout)
	if err != nil {
		return nil, err
	}
	result := &DuoResult{}
	err = json.Unmarshal(payload, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (d *Duo) push_with_mfa(mfa_code string) (MfaResult, error) {
	params := url.Values{}
	params["username"] = []string{d.DuoDto.UserName}
	params["factor"] = []string{"passcode"}
	params["async"] = []string{"1"}
	params["passcode"] = []string{mfa_code}

	_, payload, err := d.api.SignedCall("POST", "/auth/v2/auth", params, duoapi.UseTimeout)
	if err != nil {
		return nil, err
	}

	result, err := parse(payload)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (d *Duo) Apply(args []string, response *MfaResult) error {

	if len(args) != 1 {
		return errors.New("args must be equal to 1")
	}
	executor.Apply(args[0], (*response).Result())
	return nil
}

func parse(response []byte) (*DuoResult, error) {
	result := &DuoResult{}
	err := json.Unmarshal(response, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
