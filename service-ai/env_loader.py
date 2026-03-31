"""从 ENV_FILE、当前目录、service-ai 目录或仓库根目录加载 .env（与 Go 端约定一致）。"""

from __future__ import annotations

import os
from pathlib import Path

from dotenv import load_dotenv


def load_project_dotenv() -> None:
    if path := os.getenv("ENV_FILE"):
        load_dotenv(path, override=False)
        return

    here = Path(__file__).resolve().parent
    seen: list[Path] = []
    for p in (Path.cwd() / ".env", here / ".env", here.parent / ".env"):
        rp = p.resolve()
        if rp in seen:
            continue
        seen.append(rp)
        if rp.is_file():
            load_dotenv(rp, override=False)
            return

    load_dotenv(override=False)
