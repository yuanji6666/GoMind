import os

from pipeline import RagPipeline

global_pipeline = RagPipeline(
    url=os.environ.get("QDRANT_URL") or "http://localhost:6333",
    api_key=os.environ.get("QDRANT_API_KEY") or "",
    collection_name="global"
)

