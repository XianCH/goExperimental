package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/update"
	"github.com/elastic/go-elasticsearch/v8/typedapi/some"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type Config struct {
	Addresses []string
}

// Review 评价数据
type Review struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"userID"`
	Score       uint8     `json:"score"`
	Content     string    `json:"content"`
	Tags        []Tag     `json:"tags"`
	Status      int       `json:"status"`
	PublishTime time.Time `json:"publishDate"`
}

// Tag 评价标签
type Tag struct {
	Code  int    `json:"code"`
	Title string `json:"title"`
}

// createIndex 创建索引
func CreateIndex(client *elasticsearch.TypedClient) {
	resp, err := client.Indices.
		Create("my-review-1").
		Do(context.Background())
	if err != nil {
		log.Printf("create index failed, err:%v\n", err)
		return
	}
	log.Printf("index:%#v\n", resp.Index)
}

// indexDocument 索引文档
func indexDocument(client *elasticsearch.TypedClient) {
	// 定义 document 结构体对象
	d1 := Review{
		ID:      1,
		UserID:  147982601,
		Score:   5,
		Content: "这是一个好评！",
		Tags: []Tag{
			{1000, "好评"},
			{1100, "物超所值"},
			{9000, "有图"},
		},
		Status:      2,
		PublishTime: time.Now(),
	}

	// 添加文档
	resp, err := client.Index("my-review-1").
		Id(strconv.FormatInt(d1.ID, 10)).
		Document(d1).
		Do(context.Background())
	if err != nil {
		log.Printf("indexing document failed, err:%v\n", err)
		return
	}
	log.Printf("result:%#v\n", resp.Result)
}

func GetDocument(client *elasticsearch.TypedClient, id string) {
	resp, err := client.Get("my-review-1", id).
		Do(context.Background())
	if err != nil {
		log.Printf("get document by id failed, err:%v\n", err)
		return
	}
	log.Printf("fields:%s\n", resp.Source_)
}

func SearchDocument(client *elasticsearch.TypedClient) {
	// 搜索文档
	resp, err := client.Search().
		Index("my-review-1").
		Query(&types.Query{
			MatchAll: &types.MatchAllQuery{},
		}).
		Do(context.Background())
	if err != nil {
		log.Printf("search document failed, err:%v\n", err)
		return
	}
	log.Printf("total: %d\n", resp.Hits.Total.Value)
	// 遍历所有结果
	for _, hit := range resp.Hits.Hits {
		log.Printf("%s\n", hit.Source_)
	}
}

func SearchDocument2(client *elasticsearch.TypedClient) {
	// 搜索content中包含好评的文档
	resp, err := client.Search().
		Index("my-review-1").
		Query(&types.Query{
			MatchPhrase: map[string]types.MatchPhraseQuery{
				"content": {Query: "好评"},
			},
		}).
		Do(context.Background())
	if err != nil {
		log.Printf("search document failed, err:%v\n", err)
		return
	}
	log.Printf("total: %d\n", resp.Hits.Total.Value)
	// 遍历所有结果
	for _, hit := range resp.Hits.Hits {
		log.Printf("%s\n", hit.Source_)
	}
}

//聚合

func AggregationDemo(client *elasticsearch.TypedClient) {
	avgScoreAgg, err := client.Search().
		Index("my-review-1").
		Request(
			&search.Request{
				Size: some.Int(0),
				Aggregations: map[string]types.Aggregations{
					"avg_score": { // 将所有文档的 score 的平均值聚合为 avg_score
						Avg: &types.AverageAggregation{
							Field: some.String("score"),
						},
					},
				},
			},
		).Do(context.Background())
	if err != nil {
		log.Printf("aggregation failed, err:%v\n", err)
		return
	}
	log.Printf("avgScore:%#v\n", avgScoreAgg.Aggregations["avg_score"])
}

// updateDocument 更新文档
func UpdateDocument(client *elasticsearch.TypedClient) {
	// 修改后的结构体变量
	d1 := Review{
		ID:      1,
		UserID:  147982601,
		Score:   5,
		Content: "这是一个修改后的好评！", // 有修改
		Tags: []Tag{ // 有修改
			{1000, "好评"},
			{9000, "有图"},
		},
		Status:      2,
		PublishTime: time.Now(),
	}

	resp, err := client.Update("my-review-1", "1").
		Doc(d1). // 使用结构体变量更新
		Do(context.Background())
	if err != nil {
		log.Printf("update document failed, err:%v\n", err)
		return
	}
	log.Printf("result:%v\n", resp.Result)
}

// updateDocument2 更新文档
func UpdateDocument2(client *elasticsearch.TypedClient) {
	// 修改后的JSON字符串
	str := `{
					"id":1,
					"userID":147982601,
					"score":5,
					"content":"这是一个二次修改后的好评！",
					"tags":[
						{
							"code":1000,
							"title":"好评"
						},
						{
							"code":9000,
							"title":"有图"
						}
					],
					"status":2,
					"publishDate":"2023-12-10T15:27:18.219385+08:00"
				}`
	// 直接使用JSON字符串更新
	resp, err := client.Update("my-review-1", "1").
		Request(&update.Request{
			Doc: json.RawMessage(str),
		}).
		Do(context.Background())
	if err != nil {
		log.Printf("update document failed, err:%v\n", err)
		return
	}
	log.Printf("result:%v\n", resp.Result)
}

// deleteDocument 删除 document
func DeleteDocument(client *elasticsearch.TypedClient) {
	resp, err := client.Delete("my-review-1", "1").
		Do(context.Background())
	if err != nil {
		log.Printf("delete document failed, err:%v\n", err)
		return
	}
	log.Printf("result:%v\n", resp.Result)
}

// deleteIndex 删除 index
func DeleteIndex(client *elasticsearch.TypedClient) {
	resp, err := client.Indices.
		Delete("my-review-1").
		Do(context.Background())
	if err != nil {
		log.Printf("delete document failed, err:%v\n", err)
		return
	}
	log.Printf("Acknowledged:%v\n", resp.Acknowledged)
}

func main() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	}
	es, err := elasticsearch.NewTypedClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s\n", err)
	} else {
		fmt.Println("success")
	}

	// 1. Create an index
	CreateIndex(es)

	// 2. Index a document
	indexDocument(es)

	// 3. Get the document by ID
	GetDocument(es, "1")

	// 4. Search for documents
	SearchDocument(es)

	// 5. Perform an aggregation
	AggregationDemo(es)

	// 6. Update the document
	UpdateDocument(es)

	// 7. Delete the document
	DeleteDocument(es)

	// 8. Delete the index
	DeleteIndex(es)
}
