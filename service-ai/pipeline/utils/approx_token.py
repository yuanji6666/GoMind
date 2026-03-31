import tiktoken as tk

def approx_token(text: str) -> int :
    encoding = tk.get_encoding('cl100k_base')
    tokens = encoding.encode(text)
    return len(tokens)




