package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoClient struct {
	cl *mongo.Client
}
type Database interface {
	// Burada veritabanı ile ilgili işlemleri tanımlayabilirsiniz.
}
type Client interface {
	Database(string) Database
	Connect(context.Context) error
	Disconnect(context.Context) error
	StartSession() (mongo.Session, error)
	UseSession(ctx context.Context, fn func(mongo.SessionContext) error) error
	Ping(context.Context) error
}

func NewClient(connection string) (Client, error) {
	time.Local = time.UTC
	c, err := mongo.NewClient(options.Client().ApplyURI(connection))

	return &mongoClient{cl: c}, err
}

// Connect methodunu implement ediyoruz.
func (m *mongoClient) Connect(ctx context.Context) error {
	return m.cl.Connect(ctx)
}

// Disconnect methodunu implement ediyoruz.
func (m *mongoClient) Disconnect(ctx context.Context) error {
	return m.cl.Disconnect(ctx)
}

// Database methodunu implement ediyoruz.
func (m *mongoClient) Database(name string) Database {
	// Burada gerçek veritabanı işlemleri olmalı.
	// Şimdilik sadece interface'e uygun boş bir dönüş ekliyoruz.
	return nil
}

// StartSession methodunu implement ediyoruz.
func (m *mongoClient) StartSession() (mongo.Session, error) {
	return m.cl.StartSession()
}

// UseSession methodunu implement ediyoruz.
func (m *mongoClient) UseSession(ctx context.Context, fn func(mongo.SessionContext) error) error {
	return m.cl.UseSession(ctx, fn)
}

// Ping methodunu implement ediyoruz.
func (m *mongoClient) Ping(ctx context.Context) error {
	return m.cl.Ping(ctx, nil)
}
