package mongo_demo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go_demo/mongo_demo/client"
	"go_demo/mongo_demo/entry"
	"testing"
)

func TestProfile(t *testing.T) {

	collProfile := client.GetCollProfile()

	ctx := context.TODO()
	filter := map[string]string {
		"ipv4": "160.172.121.147",
	}

	// 查询一个profile
	profile := &entry.Profile{}
	_ = collProfile.FindOne(ctx, filter).Decode(profile)
	fmt.Println(profile)

	// 查看共有多少种兴趣爱好
	res, _ := collProfile.Distinct(ctx, "hobby", map[string]string{})
	fmt.Println(res)

	// 查询有多少个人不重复名字
	resName, _ := collProfile.Distinct(ctx, "name", map[string]string{})
	fmt.Println(len(resName))

}

func TestProfileDistinctName(t *testing.T) {
	collProfile := client.GetCollProfile()

	ctx := context.TODO()
	// 查询有多少个人不重复名字
	resName, err := collProfile.Distinct(ctx, "name", map[string]string{})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(len(resName))
}

func TestInsertProfile(t *testing.T) {
	
	collProfile := client.GetCollProfile()
	
	profile := &entry.ProfileInsert{
		Name:      "ljr",
		Phone:     "13630461363",
		Email:     "13630461363@qq.com",
		Ipv4:      "192.168.1.101",
		Hobby:     []string{"swimming"},
	}

	ctx := context.TODO()
	res, err := collProfile.InsertOne(ctx, profile)
	if err != nil {
		t.Log(err)
		return
	}

	t.Log(res.InsertedID)  // 5fd4b87ba076d7677c51593a

}

func TestUpdateProfile(t *testing.T) {

	collProfile := client.GetCollProfile()

	oid, err := primitive.ObjectIDFromHex("5fd4b87ba076d7677c51593a")
	if err != nil {
		t.Error(err)
		return
	}
	filter := &bson.M{
		"_id": oid,
	}

	update := &bson.M{
		"$set": &bson.M{
			"name": "msh",
		},
	}

	ctx := context.TODO()
	res, err := collProfile.UpdateOne(ctx, filter, update)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v", res)

	profile := &entry.Profile{}
	err = collProfile.FindOne(ctx, filter).Decode(profile)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v", profile)

}

func TestDeleteProfile(t *testing.T) {
	collProfile := client.GetCollProfile()
	oid, err := primitive.ObjectIDFromHex("5fd4b87ba076d7677c51593a")
	if err != nil {
		t.Error(err)
		return
	}
	filter := &bson.M{
		"_id": oid,
	}

	ctx := context.TODO()
	res, err := collProfile.DeleteOne(ctx, filter)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v", res)

}
