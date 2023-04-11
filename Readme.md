# dcard backend - linked page

- [repo link](https://github.com/peterxcli/linked-page)
- [swagger](https://peterxcli.github.io/linked-page/)


## setup
### use `docker compose`
```
docker compose up -d
``` 
> or

### install
```
make install
```

### test
```
make test
```

### run 
```
make
```
> link to http://localhost:9000/swagger/index.html to view or interact with the openAPI docs

## service endpoints
### RESTful api: port 9000
> https://peterxcli.github.io/linked-page/

### gRPC 
#### list proto : 

```
syntax = "proto3";
option go_package = "/protos/list";
service List {
    rpc GetList(GetListRequest) returns (GetListResponse);
    rpc PatchList(PatchListRequest) returns (SetHeadResponse);
    rpc InsertList(InsertListRequest) returns (SetHeadResponse);
}

message GetListRequest {
    uint32 ListId = 1;
    uint32 UserId = 2;
}

message GetListResponse {
    uint32 ListId = 1;
    uint32 UserId = 2;
    uint32 HeadId = 3;
}

message PatchListRequest {
	uint32 ListId = 1;
	uint32 UserId = 2;
	uint32 HeadId = 3;
}

message InsertListRequest {
	uint32 ListId = 1;
	uint32 UserId = 2;
	uint32 HeadId = 3;
}

message SetHeadResponse {
    bool IsSuccess = 1;
}
```

#### page proto : 
```
syntax = "proto3";
option go_package = "/proto/page";

service Page {
    rpc GetPage(GetPageRequest) returns (GetPageResponse);
    rpc SetPage(SetPageRequest) returns (SetPageResponse);
    rpc InsertPage(InsertPageRequest) returns (InsertPageResponse);
    rpc DeletePageCertainHourBefore(DeletePageRequest) returns (DeletePageResponse);
    rpc DeletePage(DeletePageRequest) returns (DeletePageResponse);
}

message GetPageRequest {
    uint32 PageId = 1;
}

message GetPageResponse {
    uint32 PageId = 1;
	uint32 NextPageId = 2;
	uint32 PrevPageId = 3;
    repeated uint32 ArticleIds = 4;
}

message SetPageRequest {
	uint32 PageId = 1;
    uint32 NextPageId = 2;
    uint32 PrevPageId = 3;
    repeated uint32 ArticleIds = 4;
}

message SetPageResponse {
    bool IsSuccess = 1;
}

message InsertPageRequest {
    uint32 PageId = 1;
    uint32 NextPageId = 2;
    uint32 PrevPageId = 3;
    repeated uint32 ArticleIds = 4;
}

message InsertPageResponse {
    bool IsSuccess = 1;
    uint32 NewPageId = 2;
}

message DeletePageRequest {
    uint32 PageId = 1;
    int32 Hour = 2;
}

message DeletePageResponse {
    uint32 RowAffected = 1;
}
```

## Why use PostgresSQL?
PostgreSQL is a powerful and versatile relational database management system that is open source and free to use. Here are some of the reasons why you might choose to use PostgreSQL:

1. Reliability and Stability: PostgreSQL is known for its reliability and stability, even when dealing with large amounts of data and complex queries. It has a reputation for being a robust database that is used by many large organizations and businesses.

2. Scalability: PostgreSQL is designed to handle high traffic and a large number of users. It can easily scale up to handle growing amounts of data and increased demand.

3. Advanced Features: PostgreSQL has many advanced features that make it a popular choice for developers. These include support for JSON, XML, and spatial data types, as well as advanced indexing and query optimization.

4. Extensibility: PostgreSQL is highly extensible, allowing developers to add custom functionality through its extensive library of extensions and plugins.

5. Open Source: PostgreSQL is an open source database, meaning that it is freely available and can be modified and customized to fit your specific needs. This can save you money on licensing fees and provide greater flexibility in your development process.

------

- 嗚嗚嗚寫不完了QQ
