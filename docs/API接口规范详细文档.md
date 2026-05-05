# GoMind API 接口规范文档

## 概述

GoMind 是一个完全模仿 X 平台功能的应用。本文档详细规定了所有前端需要的 API 接口，供后端实现参考。

## API 基础信息

- **Base URL**: `http://localhost:8080/api/v1`
- **认证方式**: Bearer JWT Token (HTTP Header: `Authorization: Bearer {token}`)
- **请求格式**: JSON (Content-Type: application/json)
- **响应格式**: JSON

## 通用响应格式

### 成功响应 (2xx)
```json
{
  "code": 0,
  "message": "success",
  "data": {}
}
```

### 错误响应
```json
{
  "code": 400,
  "message": "error message",
  "data": null
}
```

## 认证相关接口

### 1. 获取验证码

**请求**
```http
POST /user/captcha
Content-Type: application/json

{
  "email": "user@example.com"
}
```

**响应**
```json
{
  "code": 0,
  "message": "验证码已发送到邮箱",
  "data": null
}
```

### 2. 用户注册

**请求**
```http
POST /user/register
Content-Type: application/json

{
  "email": "user@example.com",
  "username": "newuser",
  "password": "password123",
  "captcha": "123456"
}
```

**响应**
```json
{
  "code": 0,
  "message": "注册成功",
  "data": {
    "id": 1,
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

### 3. 用户登录

**请求**
```http
POST /user/login
Content-Type: application/json

{
  "username": "user@example.com",
  "password": "password123"
}
```

**响应**
```json
{
  "code": 0,
  "message": "登录成功",
  "data": {
    "id": 1,
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": 1,
      "username": "newuser",
      "email": "user@example.com",
      "avatar": "https://...",
      "cover": "https://...",
      "bio": "Hello World",
      "location": "Beijing",
      "website": "https://example.com",
      "createdAt": "2024-01-01T00:00:00Z",
      "followersCount": 10,
      "followingCount": 5,
      "postsCount": 3
    }
  }
}
```

## 用户相关接口

### 4. 获取当前用户信息

**请求**
```http
GET /user/me
Authorization: Bearer {token}
```

**响应**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 1,
    "username": "newuser",
    "email": "user@example.com",
    "avatar": "https://...",
    "cover": "https://...",
    "bio": "Hello World",
    "location": "Beijing",
    "website": "https://example.com",
    "createdAt": "2024-01-01T00:00:00Z",
    "followersCount": 10,
    "followingCount": 5,
    "postsCount": 3
  }
}
```

### 5. 获取指定用户信息

**请求**
```http
GET /user/:userId
Authorization: Bearer {token}
```

**响应**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 1,
    "username": "newuser",
    "email": "user@example.com",
    "avatar": "https://...",
    "cover": "https://...",
    "bio": "Hello World",
    "location": "Beijing",
    "website": "https://example.com",
    "createdAt": "2024-01-01T00:00:00Z",
    "followersCount": 10,
    "followingCount": 5,
    "postsCount": 3
  }
}
```

### 6. 更新用户信息

**请求**
```http
PUT /user/:userId
Authorization: Bearer {token}
Content-Type: application/json

{
  "bio": "New bio",
  "location": "Shanghai",
  "website": "https://newsite.com"
}
```

**响应**
```json
{
  "code": 0,
  "message": "更新成功",
  "data": {
    // 更新后的用户对象
  }
}
```

### 7. 上传头像

**请求**
```http
POST /user/:userId/avatar
Authorization: Bearer {token}
Content-Type: multipart/form-data

FormData:
- avatar: File (image file)
```

**响应**
```json
{
  "code": 0,
  "message": "上传成功",
  "data": {
    "avatar": "https://cdn.example.com/avatar.jpg"
  }
}
```

### 8. 上传封面

**请求**
```http
POST /user/:userId/cover
Authorization: Bearer {token}
Content-Type: multipart/form-data

FormData:
- cover: File (image file)
```

**响应**
```json
{
  "code": 0,
  "message": "上传成功",
  "data": {
    "cover": "https://cdn.example.com/cover.jpg"
  }
}
```

### 9. 获取用户的帖子

**请求**
```http
GET /user/:userId/posts?page=1&pageSize=20
Authorization: Bearer {token}
```

**响应**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "items": [
      {
        "id": "post-uuid-1",
        "userId": 1,
        "author": { /* 用户对象 */ },
        "content": "Hello World",
        "images": ["https://..."],
        "createdAt": "2024-01-01T00:00:00Z",
        "likesCount": 5,
        "commentsCount": 2,
        "retweetsCount": 1,
        "isLiked": false,
        "isRetweeted": false
      }
    ],
    "total": 100,
    "page": 1,
    "pageSize": 20,
    "hasMore": true
  }
}
```

### 10. 获取用户赞过的帖子

**请求**
```http
GET /user/:userId/likes?page=1&pageSize=20
Authorization: Bearer {token}
```

**响应**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "items": [ /* 帖子数组 */ ],
    "total": 50,
    "page": 1,
    "pageSize": 20,
    "hasMore": true
  }
}
```

## 帖子相关接口

### 11. 获取主页时间线（推荐内容）

**请求**
```http
GET /posts/feed?page=1&pageSize=20
Authorization: Bearer {token}
```

**响应**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "items": [
      {
        "id": "post-uuid",
        "userId": 1,
        "author": {
          "id": 1,
          "username": "user1",
          "avatar": "https://...",
          "email": "user1@example.com"
        },
        "content": "Great post!",
        "images": ["https://image1.jpg", "https://image2.jpg"],
        "createdAt": "2024-01-01T12:00:00Z",
        "likesCount": 100,
        "commentsCount": 20,
        "retweetsCount": 5,
        "isLiked": true,
        "isRetweeted": false
      }
    ],
    "total": 1000,
    "page": 1,
    "pageSize": 20,
    "hasMore": true
  }
}
```

### 12. 获取探索页面（发现内容）

**请求**
```http
GET /posts/explore?page=1&pageSize=20
Authorization: Bearer {token}
```

**响应**
同主页时间线格式

### 13. 获取单个帖子详情

**请求**
```http
GET /posts/:postId
Authorization: Bearer {token}
```

**响应**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": "post-uuid",
    "userId": 1,
    "author": { /* 用户对象 */ },
    "content": "Hello World",
    "images": ["https://..."],
    "createdAt": "2024-01-01T00:00:00Z",
    "likesCount": 100,
    "commentsCount": 20,
    "retweetsCount": 5,
    "isLiked": false,
    "isRetweeted": false
  }
}
```

### 14. 创建帖子

**请求**
```http
POST /posts
Authorization: Bearer {token}
Content-Type: multipart/form-data

FormData:
- content: "Hello World" (必需)
- images[0]: File (可选，可多个)
- images[1]: File
- replyToId: "post-uuid" (可选，回复帖子)
```

**响应**
```json
{
  "code": 0,
  "message": "发布成功",
  "data": {
    "id": "post-uuid",
    "userId": 1,
    "author": { /* 当前用户对象 */ },
    "content": "Hello World",
    "images": ["https://..."],
    "createdAt": "2024-01-01T00:00:00Z",
    "likesCount": 0,
    "commentsCount": 0,
    "retweetsCount": 0,
    "isLiked": false,
    "isRetweeted": false
  }
}
```

### 15. 删除帖子

**请求**
```http
DELETE /posts/:postId
Authorization: Bearer {token}
```

**响应**
```json
{
  "code": 0,
  "message": "删除成功",
  "data": null
}
```

### 16. 搜索帖子

**请求**
```http
GET /posts/search?query=golang&page=1&pageSize=20
Authorization: Bearer {token}
```

**响应**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "items": [ /* 帖子数组 */ ],
    "total": 50,
    "page": 1,
    "pageSize": 20,
    "hasMore": true
  }
}
```

## 点赞相关接口

### 17. 点赞帖子

**请求**
```http
POST /posts/:postId/like
Authorization: Bearer {token}
```

**响应**
```json
{
  "code": 0,
  "message": "点赞成功",
  "data": null
}
```

### 18. 取消点赞帖子

**请求**
```http
DELETE /posts/:postId/like
Authorization: Bearer {token}
```

**响应**
```json
{
  "code": 0,
  "message": "取消点赞成功",
  "data": null
}
```

### 19. 点赞评论

**请求**
```http
POST /comments/:commentId/like
Authorization: Bearer {token}
```

**响应**
```json
{
  "code": 0,
  "message": "点赞成功",
  "data": null
}
```

### 20. 取消点赞评论

**请求**
```http
DELETE /comments/:commentId/like
Authorization: Bearer {token}
```

**响应**
```json
{
  "code": 0,
  "message": "取消点赞成功",
  "data": null
}
```

### 21. 获取赞过某个帖子的用户列表

**请求**
```http
GET /posts/:postId/likes?page=1&pageSize=20
Authorization: Bearer {token}
```

**响应**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "items": [
      {
        "id": 1,
        "username": "user1",
        "avatar": "https://...",
        "email": "user1@example.com"
      }
    ],
    "total": 100,
    "page": 1,
    "pageSize": 20,
    "hasMore": true
  }
}
```

## 评论相关接口

### 22. 获取帖子的评论列表

**请求**
```http
GET /posts/:postId/comments?page=1&pageSize=20
Authorization: Bearer {token}
```

**响应**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "items": [
      {
        "id": "comment-uuid",
        "postId": "post-uuid",
        "userId": 1,
        "author": { /* 用户对象 */ },
        "content": "Great comment!",
        "createdAt": "2024-01-01T00:00:00Z",
        "likesCount": 5,
        "repliesCount": 2,
        "isLiked": false
      }
    ],
    "total": 20,
    "page": 1,
    "pageSize": 20,
    "hasMore": false
  }
}
```

### 23. 创建评论

**请求**
```http
POST /posts/:postId/comments
Authorization: Bearer {token}
Content-Type: multipart/form-data

FormData:
- content: "Great post!" (必需)
- images[0]: File (可选)
```

**响应**
```json
{
  "code": 0,
  "message": "评论成功",
  "data": {
    "id": "comment-uuid",
    "postId": "post-uuid",
    "userId": 1,
    "author": { /* 当前用户对象 */ },
    "content": "Great comment!",
    "createdAt": "2024-01-01T00:00:00Z",
    "likesCount": 0,
    "repliesCount": 0,
    "isLiked": false
  }
}
```

### 24. 删除评论

**请求**
```http
DELETE /comments/:commentId
Authorization: Bearer {token}
```

**响应**
```json
{
  "code": 0,
  "message": "删除成功",
  "data": null
}
```

### 25. 获取评论的回复

**请求**
```http
GET /comments/:commentId/replies?page=1&pageSize=20
Authorization: Bearer {token}
```

**响应**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "items": [ /* 评论数组 */ ],
    "total": 5,
    "page": 1,
    "pageSize": 20,
    "hasMore": false
  }
}
```

### 26. 回复评论

**请求**
```http
POST /comments/:commentId/replies
Authorization: Bearer {token}
Content-Type: multipart/form-data

FormData:
- content: "Nice reply!" (必需)
- images[0]: File (可选)
```

**响应**
```json
{
  "code": 0,
  "message": "回复成功",
  "data": {
    "id": "comment-uuid",
    "postId": "post-uuid",
    "userId": 1,
    "author": { /* 当前用户对象 */ },
    "content": "Nice reply!",
    "createdAt": "2024-01-01T00:00:00Z",
    "likesCount": 0,
    "repliesCount": 0,
    "isLiked": false
  }
}
```

## 关注相关接口

### 27. 关注用户

**请求**
```http
POST /user/:userId/follow
Authorization: Bearer {token}
```

**响应**
```json
{
  "code": 0,
  "message": "关注成功",
  "data": null
}
```

### 28. 取消关注用户

**请求**
```http
DELETE /user/:userId/follow
Authorization: Bearer {token}
```

**响应**
```json
{
  "code": 0,
  "message": "取消关注成功",
  "data": null
}
```

### 29. 获取用户的粉丝列表

**请求**
```http
GET /user/:userId/followers?page=1&pageSize=20
Authorization: Bearer {token}
```

**响应**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "items": [
      {
        "id": 1,
        "username": "user1",
        "avatar": "https://...",
        "email": "user1@example.com"
      }
    ],
    "total": 100,
    "page": 1,
    "pageSize": 20,
    "hasMore": true
  }
}
```

### 30. 获取用户的关注列表

**请求**
```http
GET /user/:userId/following?page=1&pageSize=20
Authorization: Bearer {token}
```

**响应**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "items": [ /* 用户数组 */ ],
    "total": 50,
    "page": 1,
    "pageSize": 20,
    "hasMore": true
  }
}
```

### 31. 检查是否关注某个用户

**请求**
```http
GET /user/:userId/following-status
Authorization: Bearer {token}
```

**响应**
```json
{
  "code": 0,
  "message": "success",
  "data": true
}
```

## 通知相关接口

### 32. 获取通知列表

**请求**
```http
GET /notifications?page=1&pageSize=20
Authorization: Bearer {token}
```

**响应**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "items": [
      {
        "id": "notif-uuid",
        "userId": 1,
        "actorId": 2,
        "actor": { /* 用户对象 */ },
        "type": "like",
        "targetId": "post-uuid",
        "isRead": false,
        "createdAt": "2024-01-01T00:00:00Z"
      }
    ],
    "total": 20,
    "page": 1,
    "pageSize": 20,
    "hasMore": false
  }
}
```

### 33. 标记通知为已读

**请求**
```http
PUT /notifications/:notificationId/read
Authorization: Bearer {token}
```

**响应**
```json
{
  "code": 0,
  "message": "标记成功",
  "data": null
}
```

### 34. 标记所有通知为已读

**请求**
```http
PUT /notifications/read-all
Authorization: Bearer {token}
```

**响应**
```json
{
  "code": 0,
  "message": "标记成功",
  "data": null
}
```

## 错误代码

| 代码 | 含义 |
|-----|------|
| 400 | 请求参数错误 |
| 401 | 未授权（需要登录） |
| 403 | 禁止访问（无权限） |
| 404 | 资源不存在 |
| 409 | 冲突（如邮箱已存在） |
| 500 | 服务器内部错误 |

## 实现注意事项

1. **分页**：所有列表接口都支持分页，default: page=1, pageSize=20
2. **图片**：所有图片上传应返回可访问的 URL
3. **时间戳**：统一使用 ISO 8601 格式（例如：2024-01-01T00:00:00Z）
4. **UUID**：帖子和评论使用 UUID，用户使用自增 ID
5. **计数器**：需要实时更新相关计数（赞数、评论数等）
6. **通知**：创建帖子评论、被赞、被关注、被评论时应创建相应通知
7. **时间线算法**：主页和探索页面的时间线应使用不同的推荐算法
