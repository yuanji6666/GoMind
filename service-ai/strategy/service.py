import asyncio
import os
import sys
from pathlib import Path

# Running `python strategy/service.py` puts `strategy/` on sys.path; ensure project root is first.
if __package__ is None:
    sys.path.insert(0, str(Path(__file__).resolve().parents[1]))

from pipeline import RagPipeline
from strategy.schemas import ChatMessage, SourceOut

from langchain_core.prompts import ChatPromptTemplate
from langchain_core.output_parsers import StrOutputParser
from langchain_openai import ChatOpenAI


base_url = os.environ.get("OPENAI_BASE_URL") or ""
model_name = os.environ.get("OPENAI_MODEL_NAME") or ""

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
    
    messages=[{
        'role':'system', 'content':"""你是一个智能助手。请判断用户的问题：
        1. 如果是简单的问候、常识或你可以直接确定的事实，请直接回答。
        2. 如果问题涉及具体的私有文档、需要查询特定资料或你不能确定的复杂情况，请**只输出**标记：[NEED_RAG]"""
    }]
    
    for h in history:
        messages.append({'role': h.role, 'content': h.content})
    
    messages.append({'role':'user', 'content': message})

    template = ChatPromptTemplate(messages)
    

    llm = ChatOpenAI(base_url=base_url,model=model_name )
    
    chain = template | llm | StrOutputParser() 

    answer = chain.invoke(input = {})
    

    if '[NEED_RAG]' not in answer:
        print("="*10+'llm选择直接输出'+"="*10)
        print(answer)
        return answer, []
    print('='*10+'llm选择查询知识库'+'='*10)
    print(answer)
    
    docs = await asyncio.to_thread(pipeline.search, message, top_k,user_kb_id)
    blocks = [d.page_content for d in docs]
    
    messages[0] = {
        'role':'system',
        'content':'你是一个智能助手,你需要根据从知识库中搜索得来的资料来回答用户的问题'
    }

    messages[-1]['content']+= ("[资料]:"+ "\n".join(blocks))
    
    template = ChatPromptTemplate(messages)
    
    chain = template | llm | StrOutputParser()
    
    answer = chain.invoke(input={})
    
    print(answer)

    sources = [
        SourceOut(
            index=i + 1,
            content=d.page_content[:2000] + ("…" if len(d.page_content) > 2000 else ""),
            metadata=dict(d.metadata) if d.metadata else {},
        )
        for i, d in enumerate(docs)
    ]
    return answer, sources

    


    
