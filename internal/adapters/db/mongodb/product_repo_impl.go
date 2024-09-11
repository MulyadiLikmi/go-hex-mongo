package mongodb

import (
	"context"
	"go-hex-mongo/internal/domains/entity"
	"go-hex-mongo/internal/ports"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepoImpl struct {
	collection *mongo.Collection
}

func NewProductRepoImpl(db *mongo.Database, collectionName string) ports.IProductRepo {
	return &ProductRepoImpl{
		collection: db.Collection(collectionName),
	}
}

func (r *ProductRepoImpl) CreateProduct(product *entity.Product) error {
	_, err := r.collection.InsertOne(context.Background(), product)
	return err
}

func (r *ProductRepoImpl) UpdateProduct(id string, product *entity.Product) error {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"name":  product.Name,
			"stock": product.Stock,
		},
	} // Only update name and stock
	_, err := r.collection.UpdateOne(context.Background(), filter, update)
	return err
}

func (r *ProductRepoImpl) DeleteProduct(id string) error {
	filter := bson.M{"_id": id}
	_, err := r.collection.DeleteOne(context.Background(), filter)
	return err
}

func (r *ProductRepoImpl) GetProductByID(id string) (*entity.Product, error) {
	var product entity.Product
	filter := bson.M{"_id": id}
	err := r.collection.FindOne(context.Background(), filter).Decode(&product)
	return &product, err
}

func (r *ProductRepoImpl) GetAllProducts() ([]*entity.Product, error) {
	var products []*entity.Product
	cursor, err := r.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.Background(), &products)
	return products, err
}
