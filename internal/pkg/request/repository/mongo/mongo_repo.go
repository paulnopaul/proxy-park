package mongo

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"proxy-park/internal/pkg/domain"
	"time"
)

type mongoRepo struct {
	requestsCollection *mgo.Collection
}

func NewMongoRepo(requestsCollection *mgo.Collection) domain.RequestRepository {
	return &mongoRepo{
		requestsCollection: requestsCollection,
	}
}

type mongoRequest struct {
	http.Request
	Ts int64
}

func (m mongoRepo) SaveRequest(request http.Request) error {
	err := m.requestsCollection.Insert(mongoRequest{
		Request: request,
		Ts:      time.Now().Unix(),
	})
	if err != nil {
		return err
	}
	return nil
}

func (m mongoRepo) GetAllRequests() ([]http.Request, error) {
	var found []mongoRequest
	err := m.requestsCollection.Find(bson.M{}).All(&found)
	if err != nil {
		return nil, err
	}
	var res []http.Request
	for _, req := range found {
		res = append(res, req.Request)
	}
	return res, nil
}

func (m mongoRepo) GetRequest(id string) (http.Request, error) {
	found := &mongoRequest{}
	err := m.requestsCollection.Find(bson.M{
		"id": id,
	}).One(found)
	if err != nil {
		return http.Request{}, err
	}

	res := found.Request
	return res, nil
}
