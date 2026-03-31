from langchain_qdrant import FastEmbedSparse

sparse_embedding_model = FastEmbedSparse()

_sparse_embedding_model = None

def get_global_parse_model():
    global _sparse_embedding_model
    if _sparse_embedding_model is None:
        _sparse_embedding_model = FastEmbedSparse()
    return _sparse_embedding_model
    