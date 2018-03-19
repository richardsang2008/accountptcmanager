package test

import (
	"fmt"
	"github.com/gin-gonic/gin/json"
	"github.com/richardsang2008/accountptcmanager/model"
	"log"
	"net/http"
	"testing"
	//	"sync"
	"sync"
)

func callGetbyId(t *testing.T, wg *sync.WaitGroup) {

	url := fmt.Sprintf("http://localhost:4242/account/v1/2")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest", err)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}
	defer resp.Body.Close()
	var accounts []model.PogoAccount
	if err := json.NewDecoder(resp.Body).Decode(&accounts); err != nil {
		log.Println(err)
	}

	if accounts[0].Username != "30bot00022pokm" {
		t.Error("error ", "incorrect")
	}
	wg.Done()
}
func TestGetById(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go callGetbyId(t,&wg)
		//callGetbyId()
	}
	wg.Wait()
	//wg.Wait()
	fmt.Println("Done")

}
