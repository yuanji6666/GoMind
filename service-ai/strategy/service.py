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



async def hyde_search(pipeline, llm, message, history, top_k, user_kb_id):
    # 假设文档嵌入
    hypothesis_messages = [{
        'role': 'system',
        'content': (
            '根据用户问题生成一段用于知识库检索的假设答案。'
            '只输出答案内容本身，不要解释，不要提及'
            '尽量使用与知识库文档相近的表达。'
        )
    }]

    for h in history:
        hypothesis_messages.append({'role': h.role, 'content': h.content})

    hypothesis_messages.append({'role': 'user', 'content': message})

    hypothesis_template = ChatPromptTemplate.from_messages(hypothesis_messages)
    hypothesis_chain = hypothesis_template | llm | StrOutputParser()
    hypothetical_answer = await hypothesis_chain.ainvoke(input={})
    if not hypothetical_answer:
        hypothetical_answer = message

    print('='*10+'HyDE假设答案'+'='*10)
    print(hypothetical_answer)

    hyde_docs = await asyncio.to_thread(pipeline.search, hypothetical_answer, top_k, user_kb_id)
    
    return hyde_docs
async def mqe_stepback_search(pipeline, llm, message, history, top_k, user_kb_id):
    # 重写/后退
    stepback_messages = [{
        'role': 'system',
        'content': (
            '你将根据用户问题生成一个更适合知识库检索的新问题。'
            '如果原问题过于宽泛，就改写成更具体、可检索的问题；'
            '如果原问题过于具体，就改写成更高层次、概括性更强的问题。'
            '只输出新问题本身，不要解释，不要列举多个问题。'
        )
    }]

    for h in history:
        stepback_messages.append({'role': h.role, 'content': h.content})

    stepback_messages.append({'role': 'user', 'content': message})

    stepback_template = ChatPromptTemplate.from_messages(stepback_messages)
    stepback_chain = stepback_template | llm | StrOutputParser()
    rewritten_question = await stepback_chain.ainvoke(input={})
    rewritten_question = rewritten_question.strip()
    if not rewritten_question:
        rewritten_question = message

    print('='*10+'MQE/后退提示重写问题'+'='*10)
    print(rewritten_question)

    stepback_docs = await asyncio.to_thread(pipeline.search, rewritten_question, top_k, user_kb_id)
    
    return stepback_docs


    


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
        'role':'system', 'content':"""
        除非是非常简单的问题和问候直接输出答案，否则输出“[NEED_RAG]”查询知识库；
        注意！输出[NEED_RAG]时不要带有其他描述，只输出就行 [NEED_RAG]
        只要涉及到知识性的问题，就输出[NEED_RAG]，不需要过于纠结问题是否真的需要查询知识库，倾向于多查询一些相关资料来帮助回答。
        """
    }]
    
    for h in history:
        messages.append({'role': h.role, 'content': h.content})
    
    messages.append({'role':'user', 'content': message})

    template = ChatPromptTemplate(messages)
    

    llm = ChatOpenAI(base_url=base_url,model=model_name )
    
    chain = template | llm | StrOutputParser() 

    answer = await chain.ainvoke(input = {})
    

    if '[NEED_RAG]' not in answer:
        print("="*10+'llm选择直接输出'+"="*10)
        print(answer)
        return answer, []
    print('='*10+'llm选择查询知识库'+'='*10)
    print(answer)
    

    # ----------知识库查询----------
    # 并行执行
    hyde_docs, stepback_docs, original_docs = await asyncio.gather(
        hyde_search(pipeline, llm, message, history, top_k, user_kb_id),
        mqe_stepback_search(pipeline, llm ,message, history, top_k, user_kb_id),
        asyncio.to_thread(pipeline.search, message, top_k, user_kb_id)
        )

    # 去重
    merged_docs = []
    seen_docs = set()
    for docs in (hyde_docs, stepback_docs, original_docs):
        for doc in docs:
            metadata_items = tuple(sorted((doc.metadata or {}).items()))
            doc_key = (doc.page_content, metadata_items)
            if doc_key in seen_docs:
                continue
            seen_docs.add(doc_key)
            merged_docs.append(doc)


    merged_blocks = [d.page_content for d in merged_docs]
    print("一共找到%d条相关文档"%len(merged_blocks))
    
    
    messages[0] = {
        'role':'system',
        'content':'你是一个智能助手,你需要根据从知识库中搜索得来的资料来回答用户的问题'
    }

    messages[-1]['content']+= ("[资料]:"+ "\n".join(merged_blocks))
    
    template = ChatPromptTemplate(messages)
    
    chain = template | llm | StrOutputParser()
    
    answer = await chain.ainvoke(input={})
    
    print(answer)

    sources = [
        SourceOut(
            index=i + 1,
            content=d.page_content[:2000] + ("…" if len(d.page_content) > 2000 else ""),
            metadata=dict(d.metadata) if d.metadata else {},
        )
        for i, d in enumerate(merged_docs)
    ]
    return answer, sources

    


    
