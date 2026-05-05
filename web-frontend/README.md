# GoMind Frontend - X Platform Clone

完全模仿 X 平台的 Vue3 前端应用，包含用户系统、发布内容、评论、关注、点赞等核心功能。

## 项目特点

- ✨ Vue3 + TypeScript + Vite
- 🎨 完全模仿 X 平台的 UI 设计
- 📱 响应式设计，支持手机和桌面
- 🔐 JWT 认证
- 📝 完整的功能接口定义
- 🚀 模块化组件架构

## 项目结构

```
src/
├── components/          # 可复用组件
│   ├── Sidebar.vue     # 左侧导航栏
│   ├── Trends.vue      # 右侧趋势栏
│   ├── Post.vue        # 帖子卡片组件
│   ├── ComposeModal.vue # 发布帖子模态框
│   └── icons/          # 图标组件集合
├── views/              # 页面组件
│   ├── Login.vue       # 登录页
│   ├── Register.vue    # 注册页
│   ├── Home.vue        # 主页时间线
│   ├── Explore.vue     # 探索页面
│   ├── Profile.vue     # 用户资料页
│   ├── PostDetail.vue  # 帖子详情页
│   └── Placeholder.vue # 占位页面
├── stores/             # Pinia 状态管理
│   ├── user.ts        # 用户信息和认证
│   └── post.ts        # 帖子数据管理
├── api/                # API 调用
│   └── index.ts       # 所有 API 接口定义
├── router/             # 路由配置
│   └── index.ts       # 路由表和路由守卫
├── types/              # TypeScript 类型定义
│   └── index.ts       # 所有类型定义
├── styles/             # 全局样式
│   └── main.css       # 全局 CSS
├── App.vue            # 根组件
└── main.ts            # 应用入口

```

## 已实现的功能

### 认证系统
- ✅ 邮箱注册（带验证码）
- ✅ 用户登录
- ✅ JWT Token 管理
- ✅ 自动登录重定向

### 用户系统
- ✅ 用户信息展示
- ✅ 用户资料编辑
- ✅ 头像和封面上传
- ✅ 用户粉丝/关注列表

### 帖子功能
- ✅ 发布文字和图片帖子
- ✅ 主页时间线（推荐内容）
- ✅ 探索页面（发现内容）
- ✅ 帖子详情查看
- ✅ 删除自己的帖子

### 交互功能
- ✅ 点赞帖子/评论
- ✅ 发布评论
- ✅ 评论回复
- ✅ 删除评论
- ✅ 关注/取消关注用户

### 个性化功能
- ✅ 用户资料页面
- ✅ 我的帖子列表
- ✅ 我赞过的帖子
- ✅ 粉丝/关注管理

## API 接口文档

所有接口已在 `docs/GoMind完整接口文档v2.postman_collection.json` 中详细定义，包括：

### 认证 (Authentication)
- `POST /api/v1/user/captcha` - 获取验证码
- `POST /api/v1/user/register` - 用户注册
- `POST /api/v1/user/login` - 用户登录

### 用户 (User)
- `GET /api/v1/user/me` - 获取当前用户
- `GET /api/v1/user/:userId` - 获取用户信息
- `PUT /api/v1/user/:userId` - 更新用户信息
- `POST /api/v1/user/:userId/avatar` - 上传头像
- `POST /api/v1/user/:userId/cover` - 上传封面
- `GET /api/v1/user/:userId/posts` - 获取用户帖子
- `GET /api/v1/user/:userId/likes` - 获取用户赞过的帖子

### 帖子 (Post)
- `GET /api/v1/posts/feed` - 主页时间线
- `GET /api/v1/posts/explore` - 探索页面
- `GET /api/v1/posts/:postId` - 帖子详情
- `POST /api/v1/posts` - 创建帖子
- `DELETE /api/v1/posts/:postId` - 删除帖子
- `GET /api/v1/posts/search` - 搜索帖子

### 点赞 (Like)
- `POST /api/v1/posts/:postId/like` - 点赞帖子
- `DELETE /api/v1/posts/:postId/like` - 取消点赞帖子
- `GET /api/v1/posts/:postId/likes` - 获取赞过该帖子的用户

### 评论 (Comment)
- `GET /api/v1/posts/:postId/comments` - 获取帖子评论
- `POST /api/v1/posts/:postId/comments` - 创建评论
- `DELETE /api/v1/comments/:commentId` - 删除评论
- `POST /api/v1/comments/:commentId/replies` - 回复评论

### 关注 (Follow)
- `POST /api/v1/user/:userId/follow` - 关注用户
- `DELETE /api/v1/user/:userId/follow` - 取消关注
- `GET /api/v1/user/:userId/followers` - 获取粉丝列表
- `GET /api/v1/user/:userId/following` - 获取关注列表
- `GET /api/v1/user/:userId/following-status` - 检查关注状态

### 通知 (Notification)
- `GET /api/v1/notifications` - 获取通知列表
- `PUT /api/v1/notifications/:notificationId/read` - 标记通知已读
- `PUT /api/v1/notifications/read-all` - 标记所有通知已读

## 安装和运行

### 前置要求
- Node.js >= 16
- npm 或 yarn

### 安装依赖
```bash
cd web-frontend
npm install
```

### 开发模式
```bash
npm run dev
```
应用将在 `http://localhost:5173` 运行

### 生产构建
```bash
npm run build
```

### 预览生产构建
```bash
npm run preview
```

## 环境配置

在 `src/api/index.ts` 中配置 API 地址：
```typescript
const API_BASE_URL = 'http://localhost:8080/api/v1'
```

## 后端实现指南

### 需要实现的功能

1. **用户系统**
   - 用户注册、登录认证
   - JWT token 生成和验证
   - 用户信息管理（更新、上传头像/封面）

2. **帖子系统**
   - 创建、删除、查询帖子
   - 帖子图片存储和管理
   - 主页和探索时间线算法
   - 帖子搜索功能

3. **互动系统**
   - 点赞功能（帖子和评论）
   - 评论和回复功能
   - 关注功能（粉丝管理）

4. **通知系统**
   - 通知创建（当被赞、被评论、被关注时）
   - 通知查询和管理

### 数据库设计建议

#### 用户表 (users)
```sql
- id (PK)
- username (UNIQUE)
- email (UNIQUE)
- password_hash
- avatar_url
- cover_url
- bio
- location
- website
- followers_count
- following_count
- posts_count
- created_at
- updated_at
```

#### 帖子表 (posts)
```sql
- id (PK, UUID)
- user_id (FK)
- content (TEXT)
- images (JSON 数组)
- likes_count
- comments_count
- retweets_count
- created_at
- updated_at
```

#### 评论表 (comments)
```sql
- id (PK, UUID)
- post_id (FK)
- user_id (FK)
- content (TEXT)
- images (JSON 数组)
- likes_count
- created_at
- updated_at
```

#### 点赞表 (likes)
```sql
- id (PK)
- user_id (FK)
- post_id (FK, 允许 NULL)
- comment_id (FK, 允许 NULL)
- created_at
- UNIQUE(user_id, post_id)
- UNIQUE(user_id, comment_id)
```

#### 关注表 (follows)
```sql
- id (PK)
- follower_id (FK)
- following_id (FK)
- created_at
- UNIQUE(follower_id, following_id)
```

#### 通知表 (notifications)
```sql
- id (PK, UUID)
- user_id (FK)
- actor_id (FK)
- type (like, comment, follow, reply)
- target_id (post_id or user_id)
- is_read
- created_at
```

## 关键代码说明

### API 客户端 (`src/api/index.ts`)
- 使用 Axios 进行 HTTP 请求
- 自动注入 JWT token
- 统一的错误处理和响应拦截

### 状态管理 (`src/stores/`)
- User Store: 管理用户认证和个人信息
- Post Store: 管理帖子列表和分页

### 路由守卫 (`src/router/index.ts`)
- 自动重定向未认证用户到登录页
- 保护需要认证的页面

## 样式系统

- 全局变量和工具类在 `src/styles/main.css`
- 采用 X 平台的配色方案
- 支持暗黑模式的基础（可扩展）
- 响应式设计断点：1200px, 768px

## 后续功能扩展

- [ ] 消息系统（私信）
- [ ] 实时通知（WebSocket）
- [ ] 转推功能（Retweet）
- [ ] 话题标签（Hashtag）
- [ ] 趋势话题
- [ ] 用户搜索
- [ ] 高级搜索过滤
- [ ] 深色模式
- [ ] 多语言支持
- [ ] 分析统计

## 常见问题

### Q: 如何调试 API 请求？
A: 使用浏览器开发者工具的 Network 标签查看请求和响应，或在 `src/api/index.ts` 中添加日志。

### Q: 图片上传如何处理？
A: 前端通过 FormData 上传，后端需要实现文件存储（本地或云存储）和返回 URL。

### Q: 如何处理 CORS 跨域问题？
A: 后端需要设置正确的 CORS 响应头。在开发环境中，可使用 Vite 的代理配置。

## 许可证

MIT

## 贡献

欢迎提交 Issue 和 Pull Request！
