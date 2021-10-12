package controllers

import (
	"context"
	"errors"
	"fmt"

	"github.com/cmwylie19/knative-poc/api"
	"github.com/cmwylie19/knative-poc/helper"
	"github.com/cmwylie19/knative-poc/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Server struct {
	api.UnimplementedTodoServer
}

func GetTodoByName(name string, author string) (*api.Task, error) {
	result := &api.Task{}
	filter := bson.M{"name": name, "author": author}
	client, err := helper.GetMongoClient()
	if err != nil {
		return result, err
	}
	collection := client.Database(helper.DB).Collection(helper.TODOS)

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return result, err
	}
	//Return result without any error.
	return result, nil
}

func (s *Server) DeleteTodo(ctx context.Context, in *api.DeleteTodoRequest) (*api.DeleteTodoResponse, error) {
	client, err := helper.GetMongoClient()
	if err != nil {
		return nil, err
	}
	collection := client.Database(helper.DB).Collection(helper.TODOS)
	objectId, _ := primitive.ObjectIDFromHex(in.GetId())
	result := models.Task{}
	filter := bson.M{"_id": objectId}

	err = collection.FindOneAndDelete(context.TODO(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	success := fmt.Sprintf("Deleted %v.", in.GetId())
	return &api.DeleteTodoResponse{
		Message: success,
	}, nil
}

func (s *Server) GetTodo(ctx context.Context, in *api.GetTodoRequest) (*api.GetTodoResponse, error) {
	client, err := helper.GetMongoClient()
	if err != nil {
		return nil, err
	}
	collection := client.Database(helper.DB).Collection(helper.TODOS)
	objectId, _ := primitive.ObjectIDFromHex(in.GetId())
	result := models.Task{}

	filter := bson.M{"_id": objectId}

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	grpc_results := &api.Task{
		Id:       result.ID.Hex(),
		Name:     result.Name,
		Complete: result.Complete,
		Author:   result.Author,
	}
	return &api.GetTodoResponse{
		Task: grpc_results,
	}, nil
}

func (s *Server) UpdateTodo(ctx context.Context, in *api.UpdateTodoRequest) (*api.UpdateTodoResponse, error) {
	client, err := helper.GetMongoClient()
	if err != nil {
		return nil, err
	}
	collection := client.Database(helper.DB).Collection(helper.TODOS)
	objectId, _ := primitive.ObjectIDFromHex(in.GetId())

	filter := bson.M{"_id": objectId}
	_, err = collection.UpdateOne(context.TODO(), filter, bson.M{"$set": bson.M{"complete": in.GetComplete()}})
	if err != nil {
		return nil, err
	}
	return &api.UpdateTodoResponse{
		Message: "Updated todo",
	}, nil
}

func (s *Server) DeleteTodosByUser(ctx context.Context, in *api.DeleteTodosByUserRequest) (*api.DeleteTodosByUserResponse, error) {
	client, err := helper.GetMongoClient()
	if err != nil {
		return nil, err
	}
	filter := bson.M{"author": in.GetAuthor()}
	collection := client.Database(helper.DB).Collection(helper.TODOS)

	_, err = collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	success := fmt.Sprintf("Deleted %v.", in.GetAuthor())
	return &api.DeleteTodosByUserResponse{
		Message: success,
	}, nil
}

func (s *Server) GetTodosByUser(ctx context.Context, in *api.GetTodosByUserRequest) (*api.GetTodosByUserResponse, error) {
	client, err := helper.GetMongoClient()
	if err != nil {
		return nil, err
	}
	filter := bson.M{"author": in.GetAuthor()}
	tasks := []*api.Task{}
	//tasks_models := []models.Task{}

	collection := client.Database(helper.DB).Collection(helper.TODOS)
	cur, findError := collection.Find(context.TODO(), filter)
	if findError != nil {
		return nil, findError
	}

	for cur.Next(context.TODO()) {

		tm := models.Task{}
		err := cur.Decode(&tm)
		fmt.Println("T is ", tm)
		if err != nil {
			return nil, err
		}
		t := &api.Task{
			Id:       tm.ID.Hex(),
			Name:     tm.Name,
			Author:   tm.Author,
			Complete: tm.Complete,
		}
		tasks = append(tasks, t)
	}

	cur.Close(context.TODO())

	return &api.GetTodosByUserResponse{
		Task: tasks,
	}, nil
}

func (s *Server) GetTodos(ctx context.Context, in *api.GetTodosRequest) (*api.GetTodosResponse, error) {
	client, err := helper.GetMongoClient()
	if err != nil {
		return nil, err
	}
	filter := bson.M{}
	tasks := []*api.Task{}
	//tasks_models := []models.Task{}

	collection := client.Database(helper.DB).Collection(helper.TODOS)
	cur, findError := collection.Find(context.TODO(), filter)
	if findError != nil {
		return nil, findError
	}

	for cur.Next(context.TODO()) {

		tm := models.Task{}
		err := cur.Decode(&tm)
		fmt.Println("T is ", tm)
		if err != nil {
			return nil, err
		}
		t := &api.Task{
			Id:       tm.ID.Hex(),
			Name:     tm.Name,
			Author:   tm.Author,
			Complete: tm.Complete,
		}
		tasks = append(tasks, t)
	}

	cur.Close(context.TODO())

	return &api.GetTodosResponse{
		Task: tasks,
	}, nil
}
func (s *Server) CreateTodo(ctx context.Context, in *api.CreateTodoRequest) (*api.CreateTodoResponse, error) {
	client, err := helper.GetMongoClient()
	if err != nil {
		err1 := errors.New("could not connect to mongo client")
		return nil, err1
	}
	local_new_todo := models.Task{
		Name:     in.Task.GetName(),
		Author:   in.Task.GetAuthor(),
		Complete: "false",
	}
	task, err := GetTodoByName(in.Task.GetName(), in.Task.GetAuthor())
	if err != nil {
		collection := client.Database(helper.DB).Collection(helper.TODOS)
		_, err = collection.InsertOne(context.TODO(), local_new_todo)
		if err != nil {
			return nil, err
		}
		success := fmt.Sprintf("Created %v by user %v", in.Task.GetName(), in.Task.GetAuthor())
		return &api.CreateTodoResponse{
			Message: success,
		}, nil
	} else {
		return nil, fmt.Errorf("task with name %v already exists", task.Name)
	}
}
