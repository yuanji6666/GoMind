import os
import threading
from typing import Literal, Any

#Qdrant Client
from qdrant_client import QdrantClient
from qdrant_client.http.models import VectorParams, SparseVectorParams
from qdrant_client.models import Distance, KeywordIndexParams, KeywordIndexType, HnswConfigDiff
#langchain intergration
from langchain_qdrant import QdrantVectorStore, RetrievalMode


# 四种距离指标
distance_map = {
    "Cosine": Distance.COSINE,
    "Euclid": Distance.EUCLID,
    "Dot": Distance.DOT,
    "Manhattan": Distance.MANHATTAN,
}

class QdrantStoreManager:
    """
    Qdrant管理器 用来安全获取Qdrant实例 防止重复初始化
     """
    _instances: dict[tuple[str, str], QdrantVectorStore] = {}
    _lock = threading.Lock()

    @classmethod
    def get_qdrant_instance(
        cls,
        vector_size: int ,
        distance: Literal["Cosine", "Euclid", "Dot", "Manhattan"],
        dense_embedding,
        sparse_embedding,
        collection_name: str,
        url: str | None = None,
        api_key: str | None = None,
        
    ):
        api_key = api_key or os.getenv("QDRANT_API_KEY")
        url = url or os.getenv("QDRANT_URL")

        print(url, api_key)

        if not url:
            print("❌ 没有配置Qdrant Url")
            return None

        # key ("https://......", "default_collection")
        key = (url, collection_name)

        if key not in cls._instances:
            with cls._lock:
                if key not in cls._instances:
                    client = QdrantClient(
                        url=url,
                        api_key=api_key,
                    )
                    
                    if not client.collection_exists(collection_name):
                        client.create_collection(
                            collection_name=collection_name,
                            vectors_config = VectorParams(size=vector_size, distance=distance_map[distance]),
                            sparse_vectors_config= {'text':SparseVectorParams()},
                            hnsw_config=HnswConfigDiff(
                                payload_m = 16,
                                m=0
                            )
                        )
                        client.create_payload_index(
                            collection_name=collection_name,
                            field_name='user_kb_id',
                            field_schema=KeywordIndexParams(
                                type=KeywordIndexType.KEYWORD,
                                is_tenant=True,
                            )
                        )
                        print(f'集合{collection_name}已创建')
                    else:
                        print(f'复用集合{collection_name}')
                    
                    cls._instances[key] = QdrantVectorStore(
                        client=client,
                        collection_name=collection_name,
                        embedding=dense_embedding,
                        sparse_embedding=sparse_embedding,
                        retrieval_mode=RetrievalMode.HYBRID,
                        sparse_vector_name='text',
                    )
        return cls._instances[key]
    
    
    



