package post

import (
	"log"

	"github.com/BottleneckStudio/km-api/services/dynamo"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Client ...
type Client struct {
	*dynamo.Client
}

// New creates new Client instance
func New(
	tableName,
	endpoint string,
	provider dynamo.Provider,
) *Client {
	return &Client{
		dynamo.New(tableName, endpoint, provider),
	}
}

// GetCount gets the total number of item in the table
func (c *Client) GetCount() int64 {
	// input
	in := &dynamodb.DescribeTableInput{}
	in.SetTableName(c.TableName)

	resp, err := c.Provider.DescribeTable(in)
	if err != nil {
		// we log non-actionable errors
		log.Println("DescribeTable Error:", c.TableName, err.Error())
		return 0
	}

	return *resp.Table.ItemCount
}

// GetPost ...
func (c *Client) GetPost(username, id string) *Post {
	ks := "id = :id and created > :created"
	fs := "username = :username"
	vals := map[string]interface{}{
		":id":       id,
		":created":  0,
		":username": username,
	}

	// we set index name to blank since we're not querying
	// global secondary index
	out, err := c.Query("", ks, fs, vals, false, 0)
	if err != nil {
		log.Println("Query error:", err.Error())
		return nil
	}

	if len(out.Items) == 0 {
		log.Println("Query error: Not Found")
		return nil
	}

	var posts []*Post
	_ = dynamodbattribute.UnmarshalListOfMaps(out.Items, &posts)

	return posts[0]
}

// GetUserPosts gets all the posts from user
func (c *Client) GetUserPosts(username string) []*Post {
	ks := "publish = :publish and created > :created"
	fs := "username = :username"
	vals := map[string]interface{}{
		":publish":  1,
		":created":  0,
		":username": username,
	}

	out, err := c.Query("publish_index", ks, fs, vals, false, 0)
	if err != nil || len(out.Items) == 0 {
		return nil
	}

	var posts []*Post
	_ = dynamodbattribute.UnmarshalListOfMaps(out.Items, &posts)
	return posts
}

// GetPosts ...
func (c *Client) GetPosts() []*Post {
	ks := "publish = :publish and created > :created"
	vals := map[string]interface{}{
		":publish": 1,
		":created": 0,
	}

	out, err := c.Query("publish_index", ks, "", vals, false, 0)
	if err != nil || len(out.Items) == 0 {
		return nil
	}

	var posts []*Post
	_ = dynamodbattribute.UnmarshalListOfMaps(out.Items, &posts)
	return posts
}
