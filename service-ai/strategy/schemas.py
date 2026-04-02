from typing import Literal

from pydantic import BaseModel, Field




class ChatMessage(BaseModel):
    role: Literal["user", "assistant"]
    content: str


class ChatRequest(BaseModel):
    message: str = Field(..., min_length=1, max_length=16000)
    history: list[ChatMessage] = Field(default_factory=list)
    top_k: int = Field(4, ge=1, le=32)


class SourceOut(BaseModel):
    index: int
    content: str
    metadata: dict


class ChatResponse(BaseModel):
    answer: str
    sources: list[SourceOut]
