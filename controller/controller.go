package controller

import (
	"github.com/richardsang2008/accountptcmanager/data"
	"github.com/richardsang2008/accountptcmanager/model"
	"github.com/richardsang2008/accountptcmanager/utility"
)

var (
	Data data.DataAccessLay
)

func AddAccount(account model.PogoAccount) (*string, error) {
	utility.MLog.Debug("Controller AddAccount starting")
	newid, err := Data.AddAccount(account)
	if err != nil {
		utility.MLog.Error("Controller AddAccount error " + err.Error())
		return nil, err
	} else {
		utility.MLog.Debug("Controller AddAccount end")
		return newid, nil
	}
}
func GetAccount(id uint) (*[]model.PogoAccount, error) {
	utility.MLog.Debug("Controller GetAccount starting")
	accounts, err := Data.GetAccount(id)
	if err != nil {
		utility.MLog.Error("Controller GetAccount error " + err.Error())
		return nil, err
	} else {
		utility.MLog.Debug("Controller GetAccount end")
		return accounts, nil
	}
}
func GetAccountByUserName(username string) (*[]model.PogoAccount, error) {
	utility.MLog.Debug("Controller GetAccountByUserName starting")
	accounts, err := Data.GetAccountByUserName(username)
	if err != nil {
		utility.MLog.Error("Controller GetAccountByUserName error " + err.Error())
		return nil, err
	} else {
		utility.MLog.Debug("Controller GetAccountByUserName end")
		return accounts, nil
	}
}
func GetNextUseableAccountByLevel(minlevel, maxlevel int) (*[]model.PogoAccount, error) {
	utility.MLog.Debug("Controller GetNextUseableAccountByLevel starting")
	accounts, err := Data.GetAccountByLevel(minlevel, maxlevel)
	if err != nil {
		utility.MLog.Error("Controller GetAccountByUserName error " + err.Error())
		return nil, err
	} else {
		//filter the accounts not usable
		ret := []model.PogoAccount{}
		for _, account := range *accounts {
			if account.Banned == false && account.SystemId == "" {
				ret = append(ret, account)
			}
		}
		utility.MLog.Debug("Controller GetAccountByUserName end")
		if len(ret)>0{
			return &ret,nil
		} else {
			return nil,nil
		}

	}
}
func UpdateAccountBySpecialFields(account model.PogoAccount) (*string,error) {
	utility.MLog.Debug("Controller UpdateAccountBySpecialFields starting")
	idptr,err:=Data.UpdateAccount(account)
	if err!=nil {
		utility.MLog.Error("Controller UpdateAccountBySpecialFields error "+err.Error())
		return nil,err
	} else {
		utility.MLog.Debug("Controller UpdateAccountBySpecialFields end")
		return idptr,nil
	}
}
