package main

import (
	"gopkg.in/mgo.v2/bson"
	"testing"
	"github.com/stretchr/testify/assert"
)

var ID = bson.NewObjectId()
var user_data = &User{ID,"user_testing","test@gmail.com"}


func TestInsertData(t *testing.T)  {
	t.Logf("data used: %v \n",user_data)
	err := user_data.insert()
	if err != nil{
		t.Error(err.Error())
	}
}

func TestDataId(t *testing.T)  {
	var lookup,err = user_data.getData()
	if err != nil{
		t.Error(err.Error())
	}
	assert.Equal(t,user_data.ID,lookup.ID,"id tidak sama")
}

func TestDataUsername(t *testing.T)  {
	var lookup,err = user_data.getData()
	if err != nil{
		t.Error(err.Error())
	}
	assert.Equal(t,user_data.Username,lookup.Username,"username tidak sama")
}

func TestDataEmail(t *testing.T)  {
	var lookup,err = user_data.getData()
	if err != nil{
		t.Error(err.Error())
	}
	assert.Equal(t,user_data.Email,lookup.Email,"email tidak sama")
}

func BenchmarkLooping(b *testing.B) {
	for i := 0; i < b.N ; i++  {
		user_data.getData()
	}
}