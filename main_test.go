package main

import (
	"gopkg.in/mgo.v2/bson"
	"testing"
)

var ID = bson.NewObjectId()
var user_data = &User{ID,"user_testing","test@gmail.com"}


func TestInsertData(t *testing.T)  {
	t.Logf("data used: %v \n",user_data)
	err := user_data.insert()
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestDataId(t *testing.T)  {
	var lookup,err = user_data.getData()
	if err != nil {
		t.Errorf("tidak mendapatkan data dengan id %v",user_data.ID)
	}

	if user_data.ID != lookup.ID {
		t.Errorf("salah, data tidak match! antara %v dan %v ",user_data.ID,lookup.ID)
	}
}

func TestDataUsername(t *testing.T)  {
	var lookup,err = user_data.getData()
	if err != nil {
		t.Errorf("tidak mendapatkan data dengan id %s",user_data.Username)
	}

	if user_data.Username != lookup.Username {
		t.Errorf("salah, data tidak match! antara %s dan %s ", user_data.Username, lookup.Username)
	}
}

func TestDataEmail(t *testing.T)  {
		var lookup,err = user_data.getData()
		if err != nil {
			t.Errorf("tidak mendapatkan data dengan id %s",user_data.Email)
		}

		if user_data.Email != lookup.Email {
			t.Errorf("salah, data tidak match! antara %s dan %s ", user_data.Email, lookup.Email)
		}
}

func BenchmarkLooping(b *testing.B) {
	for i := 0; i < b.N ; i++  {
		user_data.getData()
	}
}