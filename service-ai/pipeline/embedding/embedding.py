import threading
from abc import ABC, abstractmethod
import os
from openai import OpenAI

from langchain_openai import OpenAIEmbeddings


class EmbeddingModel(ABC):

    @abstractmethod
    def encode(self, text: str | list[str] ):
        pass

    @property
    def dimension(self) -> int:
        raise NotImplementedError



class OpenAIEmbeddingModel(EmbeddingModel):
    def __init__(
        self,
        model_name: str | None = None,
        api_key: str | None = None,
        base_url: str | None = None,
    ):
        self.model_name = model_name or os.environ.get("EMBED_MODEL_NAME") or ""
        api_key = api_key or os.environ.get("EMBED_API_KEY")
        base_url = base_url or os.environ.get("EMBED_BASE_URL")
        self.client = OpenAI(
            api_key=api_key,
            base_url=base_url,
        )

        # 通过调用一次接口来获取embedding的维度
        dimension_check_response = self.encode("check_embedding_model_health")

        self._dimension = len(dimension_check_response[0])

    @property
    def dimension(self) -> int:
        return self._dimension

    def encode(self, text: str | list[str] ):
        """
        对文本进行编码器
        返回一个二维列表，每个元素是一个向量list[float]
        """
        embedding_response = self.client.embeddings.create(
            input=text,
            model=self.model_name,
            encoding_format = 'float',
            timeout = 64,
        )
        print(f'OpenAIClient:模型{self.model_name}已对文本{text}编码')

        vectors : list[list[float]] = []

        for embedding in embedding_response.data:
            vectors.append(embedding.embedding)

        return vectors


# 单例全局编码器，并发保护
global_embedding_model = None
_lock = threading.RLock()

def get_global_embedding_model(model_name: str | None = None, api_key: str | None = None, base_url: str | None = None):
    global global_embedding_model
    if global_embedding_model is not None:
        return global_embedding_model
    with _lock:
        if global_embedding_model is not None:
            return global_embedding_model
        global_embedding_model = OpenAIEmbeddingModel(model_name=model_name, api_key=api_key, base_url=base_url)
        return global_embedding_model
    



_langchain_embedder = None

def get_global_langchain_embedding():
    global _langchain_embedder
    if _langchain_embedder is None:
        _langchain_embedder = OpenAIEmbeddings(
            base_url=os.environ.get("EMBED_BASE_URL"),
            api_key=os.environ.get("EMBED_API_KEY"),
            model=os.environ.get("EMBED_MODEL_NAME"),
            check_embedding_ctx_length=False,
            chunk_size=10,
        )
    return _langchain_embedder
    
    
