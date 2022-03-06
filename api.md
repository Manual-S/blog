# 后端接口api

## 获取全部文章列表
- 接口 `api/articles/list`  
- 方法 GET
- 参数 无
- 返回值
```go
{
    "code": 0,
    "data": [
        {
            "ArticleId": "ddd",
            "Title": "你好",
            "Summary": "",
            "Content": "",
            "Status": 0,
            "Votes": 0
        }
    ],
    "message": "success"
}
```
## 根据投票获取文章列表
- 接口 `articles/socrelist`
- 方法 GET
- 参数
start 排序的开始
stop 排序的结束
- 返回值
  
## 文章投票接口
- 接口  
`articles/votes`
- 方法  
POST
- 参数  
`title_id` 文章id
- 返回值 无
  
## 更新文章接口
