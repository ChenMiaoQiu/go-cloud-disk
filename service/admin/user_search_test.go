package admin

import (
	"fmt"
	"os"
	"testing"

	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/joho/godotenv"
)

func TestUserSearchStatus(t *testing.T) {
	search := UserSearchService{
		Status: "super_admin",
	}

	res := search.UserSearch()
	if res.Code != 200 {
		t.Fatal("search status err", res)
	}
	fmt.Println(res.Data)
}

func TestUserSearchFix(t *testing.T) {
	search := UserSearchService{
		Status:   "super_admin",
		NickName: "ad",
	}

	res := search.UserSearch()
	if res.Code != 200 {
		t.Fatal("search status err", res)
	}
	fmt.Println(res.Data)
}

func TestMain(m *testing.M) {
	// before test need put a .env file in this filefolder
	godotenv.Load()
	model.Database(os.Getenv("MYSQL_DSN"))
	m.Run()
}
