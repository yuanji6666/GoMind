"""
运行: 在 service-ai 目录执行
  uvicorn api.main:app --reload --host 0.0.0.0 --port 8000
.env 优先从当前目录、service-ai 目录或上一级仓库根目录加载；也可设置环境变量 ENV_FILE 指定路径。
需配置 QDRANT_URL, EMBED_*, OPENAI_API_KEY, OPENAI_MODEL_NAME 等。
"""

from __future__ import annotations

import asyncio
import os
import tempfile
from contextlib import asynccontextmanager
from pathlib import Path


from fastapi import APIRouter, FastAPI, File, HTTPException, UploadFile
from fastapi.middleware.cors import CORSMiddleware

from strategy.schemas import ChatRequest, ChatResponse
from strategy.service import (
    chat_with_kb,
)

from globals import global_pipeline

router = APIRouter()


@asynccontextmanager
async def lifespan(_: FastAPI):
    yield


app = FastAPI(title="FeatherRAG API", version="0.1.0", lifespan=lifespan)
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

@router.post("/knowledge-bases/{user_kb_id}/documents")
async def upload_documents(user_kb_id: str, files: list[UploadFile] = File(...)):
    if not files:
        raise HTTPException(status_code=400, detail="至少上传一个文件")
    pipeline = global_pipeline
    if pipeline is None:
        raise HTTPException(status_code=404, detail="知识库不存在")
    paths: list[str] = []
    try:
        for f in files:
            suffix = Path(f.filename or "bin").suffix or ".bin"
            fd, path = tempfile.mkstemp(suffix=suffix, prefix="feather_")
            os.close(fd)
            content = await f.read()
            with open(path, "wb") as fp:
                fp.write(content)
            paths.append(path)
        await asyncio.to_thread(pipeline.add_document, paths,user_kb_id)
    except RuntimeError as e:
        raise HTTPException(status_code=503, detail=str(e)) from e
    finally:
        for p in paths:
            try:
                os.unlink(p)
            except OSError:
                pass
    return {"status": "ok", "files_ingested": len(files)}


@router.post("/knowledge-bases/{user_kb_id}/chat", response_model=ChatResponse)
async def chat(user_kb_id: str, body: ChatRequest):
    pipeline = global_pipeline
    if pipeline is None:
        raise HTTPException(status_code=404, detail="知识库不存在")
    try:
        answer, sources = await chat_with_kb(
            pipeline, user_kb_id,body.message, body.history, body.top_k
        )
    except RuntimeError as e:
        raise HTTPException(status_code=503, detail=str(e)) from e
    return ChatResponse(answer=answer, sources=sources)


app.include_router(router)
