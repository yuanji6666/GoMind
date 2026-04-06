from .load import opendataloader_load_document, unstructured_load_document
from .embedding import get_global_langchain_embedding,get_global_parse_model
from .storage import QdrantStoreManager
from .chunk import langchain_split_documents
from langchain_core.documents import Document
from qdrant_client.models import Filter, FieldCondition, MatchValue

def _load_documents_from_paths(paths: list[str]) -> list[Document]:
    """PDF 走 OpenDataLoader，其余格式走 Unstructured。"""
    pdfs = [p for p in paths if p.lower().endswith(".pdf")]
    rest = [p for p in paths if not p.lower().endswith(".pdf")]
    documents = []
    if pdfs:
        documents.extend(opendataloader_load_document(pdfs))
    if rest:
        documents.extend(unstructured_load_document(rest))
    return documents


class RagPipeline:
    """支持多租户的RAG Pipeline，底层使用Qdrant作为向量数据库。"""
    """每个用户的知识库通过 user_kb_id 区分，添加文档和搜索时都需要提供 user_kb_id 来确保数据隔离。"""
    """user_kb_id 应为“用户唯一id_知识库名”格式，例如 "user123_kb1"，以确保在同一Qdrant集合中不同用户的数据不会混淆。"""
    def __init__(
        self,
        url:str,
        api_key:str,
        collection_name:str,
    ):
        self.url = url
        self.api_key = api_key
        self.collection_name = collection_name
        
        dense_embedder = get_global_langchain_embedding()
        sparse_embedder = get_global_parse_model()
        self.qdrant_store = QdrantStoreManager.get_qdrant_instance(
            vector_size=1024,
            distance="Cosine",
            url=self.url,
            api_key=self.api_key,
            collection_name=self.collection_name,
            dense_embedding=dense_embedder,
            sparse_embedding=sparse_embedder
        )
    
    def add_document(
        self,
        paths:list[str],
        user_kb_id:str
    ):
        if self.qdrant_store is None:
            raise RuntimeError("Qdrant 未初始化，请检查 QDRANT_URL")
        document = _load_documents_from_paths(paths)
        
        for doc in document:
            doc.metadata["user_kb_id"] = user_kb_id
        
        document_chunks = langchain_split_documents(document,800)
        self.qdrant_store.add_documents(document_chunks)
    
    def search(
        self,
        query: str,
        k: int,
        user_kb_id: str 
    ):
        if self.qdrant_store is None:
            raise RuntimeError("Qdrant 未初始化，请检查 QDRANT_URL")
        result = self.qdrant_store.similarity_search(
            query=query,
            k=k,
            filter=Filter(
                must=FieldCondition(
                    key="metadata.user_kb_id",
                    match=MatchValue(value=user_kb_id)
                )
            )
        )

        
        return result
        
        