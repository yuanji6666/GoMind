import asyncio
import os

from openai import OpenAI
from openai.types.chat import ChatCompletionMessageParam

from pipeline import RagPipeline
from api.schemas import ChatMessage, SourceOut

def _chat_sync(
    question: str,
    context_blocks: list[str],
    history: list[ChatMessage],
) -> str:
    api_key = os.getenv("OPENAI_API_KEY")
    if not api_key:
        raise RuntimeError("未配置 OPENAI_API_KEY")
    base_url = os.getenv("OPENAI_BASE_URL")
    client = OpenAI(api_key=api_key, base_url=base_url or None)
    model = os.getenv("OPENAI_MODEL_NAME", "gpt-4o-mini")
    context = "\n\n".join(f"[{i + 1}] {block}" for i, block in enumerate(context_blocks))
    system = (
        "你是知识库问答助手。用户消息中会给出【当前问题】与【检索到的参考片段】（编号 [1]、[2]…）。"
        "若片段与问题相关，请优先依据片段作答，并引用对应编号；可归纳、转述，不编造片段中不存在的事实。"
        "若片段与问题无关或完全无法支撑任何要点，再说明无法从知识库中找到依据。"
    )
    # 必须把「当前问题」放在最前：若上下文很长，部分兼容接口会截断单条 user，
    # 放在末尾时「用户问题」容易被截掉，模型表现为未收到问题。
    user = (
        f"【当前问题】\n{question}\n\n"
        f"【检索到的参考片段】\n{context}\n\n"
        "请直接回答【当前问题】，并显式说明上述哪些编号支撑了你的答案。"
    )
    messages: list[ChatCompletionMessageParam] = [
        {"role": "system", "content": system},
    ]
    for m in history:
        if m.role == "user":
            messages.append({"role": "user", "content": m.content})
        else:
            messages.append({"role": "assistant", "content": m.content})
    messages.append({"role": "user", "content": user})
    
    print(messages)

    resp = client.chat.completions.create(model=model, messages=messages)
    return (resp.choices[0].message.content or "").strip()


async def chat_with_kb(
    pipeline: RagPipeline,
    user_kb_id: str,
    message: str,
    history: list[ChatMessage],
    top_k: int,
) -> tuple[str, list[SourceOut]]:
    message = message.strip()
    if not message:
        raise RuntimeError("问题内容为空")
    docs = await asyncio.to_thread(pipeline.search, message, top_k,user_kb_id)
    blocks = [d.page_content for d in docs]
    answer = await asyncio.to_thread(_chat_sync, message, blocks, history)
    sources = [
        SourceOut(
            index=i + 1,
            content=d.page_content[:2000] + ("…" if len(d.page_content) > 2000 else ""),
            metadata=dict(d.metadata) if d.metadata else {},
        )
        for i, d in enumerate(docs)
    ]
    return answer, sources
