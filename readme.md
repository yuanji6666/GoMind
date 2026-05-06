# GoMind

GoMind 是一套**基于 RAG（检索增强生成）的知识库问答**应用：用户可注册登录、创建知识库、上传文档，并在对话中结合所选知识库进行问答。

## 功能特性

- 用户注册、登录、JWT 认证
- 知识库创建与管理
- 文档上传与向量化入库
- 基于 RAG 的智能问答
- 查询优化（MQE、Stepback、HyDE）
- LLM 语义路由（简单问题直接回答，复杂问题走 RAG）
- 历史消息分页加载（游标分页）
- 消息持久化（Redis 缓存 + MySQL 异步入库）
- Docker 部署支持

## 架构概览

项目由三部分组成：

| 组件 | 说明 |
|------|------|
| **web-frontend** | Vue 3 + Vite + TypeScript 单页应用，提供登录注册与聊天界面 |
| **server-go**（模块名：`gopherAI`） | Gin HTTP 服务：用户认证（JWT）、验证码、知识库与会话等业务 API，采用 DDD 架构 |
| **service-ai**（FeatherRAG） | FastAPI 服务：文档入库、向量检索与基于知识库的问答（LangChain + Qdrant） |

前端开发时通过 Vite 将 `/api` 代理到 Go 服务；上传文档与 RAG 对话由前端直连 FastAPI（需配置 `VITE_FASTAPI_URL`）。

## 技术栈

- **前端**：Vue 3、Vite 4、TypeScript、Vue Router、Pinia、Axios、dayjs
- **后端（Go）**：Gin、GORM（MySQL）、Redis、JWT、RabbitMQ（按业务配置使用），采用 DDD 分层架构
- **AI 服务（Python）**：FastAPI、LangChain、Qdrant、uvicorn  

## 目录结构

```text
GoMind/
├── web-frontend/       # Vue 3 前端
│   ├── src/
│   │   ├── api/        # API 客户端
│   │   ├── components/ # Vue 组件
│   │   ├── views/      # 页面视图
│   │   ├── stores/     # Pinia 状态管理
│   │   └── router/      # Vue Router 配置
│   └── ...
├── server-go/          # Go API 服务（DDD 架构）
│   ├── domain/         # 领域层：实体、仓储接口
│   ├── common/         # 公共模块
│   ├── middleware/     # 中间件（JWT、CORS等）
│   ├── router/         # 路由配置
│   └── main.go
├── service-ai/         # Python RAG / FastAPI
│   ├── api/            # FastAPI 路由
│   ├── pipeline/       # RAG 处理管道
│   ├── strategy/       # 查询策略
│   └── globals/        # 全局配置
└── deploy/             # Docker 部署配置
```

## 环境要求

- **Node.js**（建议 18+），用于构建与运行前端
- **Go 1.25+**，用于编译运行 `server-go`
- **Python 3.12+**，用于 `service-ai`（建议使用 `uv` 或 `pip` + 虚拟环境）
- **MySQL**：存储用户、会话、知识库元数据等
- **Redis**：验证码缓存、消息缓存
- **Qdrant**：向量存储（RAG 检索）
- 可选：**RabbitMQ**、邮件服务（与 `.env` 中邮件相关配置配合使用）  

## 配置说明

### 1. Go 服务（`server-go`）

在项目根目录复制环境变量模板并修改：

```bash
cp .env.example .env
```

主要变量包括：`APP_HOST`、`APP_PORT`（默认示例为 `127.0.0.1` 与 `3000`）、MySQL、Redis、JWT、邮件、OpenAI 兼容接口，以及 `RAG_HOST` / `RAG_PORT`（若后端需调用 RAG 服务可在此填写）。具体以 `.env.example` 为准。

### 2. AI 服务（`service-ai`）

```bash
cd service-ai
cp .env.example .env
```

需配置 OpenAI 兼容接口、嵌入模型相关变量，以及 `QDRANT_URL`、`QDRANT_API_KEY` 等。详见 `service-ai/.env.example` 与 `service-ai/api/main.py` 顶部说明。

启动示例（在项目 `service-ai` 目录下）：

```bash
uvicorn api.main:app --reload --host 0.0.0.0 --port 8000
```

### 3. 前端（`web-frontend`）

- 开发时默认将请求代理到 `http://127.0.0.1:3000`（与 Go 服务端口一致），见 `vite.config.ts`。
- 通过环境变量 **`VITE_FASTAPI_URL`** 指定 FastAPI 地址（未设置时默认为 `http://localhost:8000`），见 `src/api/client.ts`。

可在 `web-frontend` 下创建 `.env` 或 `.env.local`，例如：

```env
VITE_FASTAPI_URL=http://127.0.0.1:8000
```

## 本地启动顺序（建议）

1. 启动 **MySQL、Redis、Qdrant** 等依赖服务。
2. 启动 **service-ai**（FastAPI，默认端口 8000，可按需修改）。
3. 在项目根目录配置 `.env` 后运行 Go 服务（默认监听 `.env` 中的 `APP_PORT`，示例为 3000）。
4. 在 **`web-frontend`** 中执行 `npm install` 与 `npm run dev`，在浏览器中访问 Vite 提示的本地地址。

## 主要 API 说明（Go，`/api/v1`）

- **用户（无需登录）**
  - `POST /user/register` — 注册
  - `POST /user/login` — 登录
  - `POST /user/captcha` — 验证码

- **AI 相关（需 JWT，前缀 `/AI`）**
  - `GET /AI/kb/list` — 知识库列表
  - `POST /AI/kb/create` — 创建知识库
  - `GET /AI/session/list` — 会话列表
  - `POST /AI/session/create` — 创建会话并发送首条消息
  - `POST /AI/session/send` — 在已有会话中发送消息

FastAPI 侧提供文档上传与基于知识库的聊天等接口（如 `POST /knowledge-bases/{user_kb_id}/documents`），与前端 `src/api/fastapi.ts` 中的调用相对应。

## Docker 部署

项目提供 `deploy/` 目录下的 Docker 配置，支持一键部署：

```bash
cd deploy
docker-compose up -d
```

## 构建前端

在 `web-frontend` 目录：

```bash
npm run build
```

产物默认输出到 `web-frontend/dist`，可由任意静态资源服务器或反向代理托管。
