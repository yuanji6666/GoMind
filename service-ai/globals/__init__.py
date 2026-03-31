from env_loader import load_project_dotenv

load_project_dotenv()

from .global_pipeline import global_pipeline  # noqa: E402

__all__ = ["global_pipeline"]