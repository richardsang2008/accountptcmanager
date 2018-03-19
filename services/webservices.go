package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/richardsang2008/accountptcmanager/controller"
	"github.com/richardsang2008/accountptcmanager/model"
	"github.com/richardsang2008/accountptcmanager/utility"
	"net/http"
	"strconv"
)

func AddAccount(c *gin.Context) {
	var account model.PogoAccount
	c.BindJSON(&account)
	utility.MLog.Debug("Services AddAccount starting")
	newid, err := controller.AddAccount(account)
	if err != nil {
		utility.MLog.Error("Services AddAccount error " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	} else {
		utility.MLog.Debug("Services AddAccount end")
		c.JSON(http.StatusOK, gin.H{"id": newid})
	}
}
func GetAccountById(c *gin.Context) {
	id := c.Params.ByName("id")
	utility.MLog.Debug("Services GetAccount starting ")
	var accounts *[]model.PogoAccount
	var err error
	u, _ := strconv.ParseUint(id, 10, 32)
	accounts, err = controller.GetAccount(uint(u))
	if err != nil {
		utility.MLog.Error("Services GetAccount error " + err.Error())
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
	} else {
		utility.MLog.Debug("Services GetAccount end ")
		c.JSON(http.StatusOK, *accounts)
	}
}

func UpdateAccountBySpecificFields(c *gin.Context) {
	utility.MLog.Debug("Services UpdateAccountBySpecificFields starting ")
	var account model.PogoAccount
	c.BindJSON(&account)
	var updateAccountIDs []uint
	if account.ID == 0 {
		//get accounts by username
		accounts, err := controller.GetAccountByUserName(account.Username)
		if err != nil {
			utility.MLog.Error("Services UpdateAccountBySpecificFields error " + err.Error())
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		} else {
			if len(*accounts) > 0 {
				for _, ac := range *accounts {
					updateAccountIDs = append(updateAccountIDs, ac.ID)
				}
			}
		}
	} else {
		updateAccountIDs = append(updateAccountIDs, account.ID)
	}
	allok := true
	for _, updateAccountID := range updateAccountIDs {
		account.ID =updateAccountID
		_, err := controller.UpdateAccountBySpecialFields(account)
		if err != nil {
			utility.MLog.Error("Services UpdateAccountBySpecificFields error " + err.Error())
			allok = false
			break
		} else {
			utility.MLog.Debug("Services UpdateAccountBySpecificFields end ")
		}
	}
	if len(updateAccountIDs) ==0 {
		utility.MLog.Debug("Services UpdateAccountBySpecificFields return no records ")
		utility.MLog.Debug("Services UpdateAccountBySpecificFields end ")
		c.JSON(http.StatusOK,gin.H{"username": account.Username})
	}
	if allok == false {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "some errors"})
	} else {
		c.JSON(http.StatusOK, gin.H{"username": account.Username})
	}
}
func GetAccountBySystemIdAndLevel(c *gin.Context) {
	utility.MLog.Debug("Services GetAccountBySystemIdAndLevel starting ")
	query := c.Request.URL.Query()
	fmt.Print(query)
	systemId := query["system_id"][0]
	countstr := query["count"][0]
	minLevelstr := query["min_level"][0]
	maxLevelstr := query["max_level"][0]
	bannedOrNewstr := query["banned_or_new"][0]
	if systemId == "" || countstr == "" || minLevelstr == "" || maxLevelstr == "" || bannedOrNewstr == "" {
		c.JSON(http.StatusBadRequest, gin.H{})
	}
	count, err := strconv.Atoi(countstr)
	if err != nil {
		utility.MLog.Error("input count wrong format", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{})
	} else {
		min, err := strconv.Atoi(minLevelstr)
		if err != nil {
			utility.MLog.Error("input min_level wrong format", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{})
		} else {
			max, err := strconv.Atoi(maxLevelstr)
			if err != nil {
				utility.MLog.Error("input max_level wrong format", err.Error())
				c.JSON(http.StatusBadRequest, gin.H{})
			} else {
				accounts, err := controller.GetNextUseableAccountByLevel(min, max)
				if err != nil {
					utility.MLog.Error("controller.GetNextUseableAccountByLevel error ", err.Error())
					c.JSON(http.StatusBadRequest, gin.H{})
				} else {
					msg := fmt.Sprintf("System ID [%s] requested %v accounts level %d-%d from server %s", systemId, count, min, max, c.Request.Host)
					utility.MLog.Info(msg)
					if accounts == nil {
						utility.MLog.Warning("Could only deliver 0 accounts")
						c.JSON(http.StatusOK, nil)
					} else {
						if count == 1 {
							account := (*accounts)[0]
							msg := fmt.Sprintf("Event for account %v: Got assigned to [%s]", account.Username, systemId)
							utility.MLog.Info(msg)
							utility.MLog.Debug("Services GetAccountBySystemIdAndLevel end ")
							c.JSON(http.StatusOK, account)
						} else {
							for _, account := range *accounts {
								msg := fmt.Sprintf("Event for account %v: Got assigned to [%s]", account.Username, systemId)
								utility.MLog.Info(msg)
							}
							utility.MLog.Debug("Services GetAccountBySystemIdAndLevel end ")
							c.JSON(http.StatusOK, *accounts)
						}

					}
				}
			}
		}
	}
}
