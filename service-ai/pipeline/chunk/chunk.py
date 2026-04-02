from pipeline.utils.approx_token import approx_token
from langchain_text_splitters import RecursiveCharacterTextSplitter
from langchain_core.documents import Document

def _split_markdown_to_paragraphs(text: str) -> list[dict]:
    """
    A paragraph is like:
    {
        'content': '....',
        'heading_path': '...',
        'start': 123,
        'end': 246,
    }
    """

    print('分段落算法启动')

    buf: list[str] = [] # 段落缓冲区
    heading_stack: list[str] = ['#']*7 # 标题存储
    char_pos: int = 0
    paragraphs: list[dict] = []
    para_level: int = 0
    is_in_code: bool = False

    # 更新缓冲区
    def flush_buf(pos: int):
        content = '\n'.join(buf)
        heading_path = ''
        for head in heading_stack:
            if head != '#':
                heading_path += (head+'>')

        if content == '':
            return


        paragraphs.append(
            {
                'heading_path': heading_path,
                'content': content,
                'start': pos-len(content),
                'end': pos,
            }
        )

        buf.clear()

    for line in text.splitlines():
        if line == '':
            continue
        if line.startswith('#'):
            flush_buf(char_pos)
        # 逐个行分析

        # 在代码块内部/外部 状态切换
        if line.startswith('```'):
            is_in_code = not is_in_code

        # 在内部 直接入buf
        if is_in_code:
            buf.append(line)
            continue


        if line.startswith('#'):

            flush_buf(char_pos)

            # parse level and heading
            level = line.count('#')
            heading = line.strip('#').strip()
            if level > para_level:
                heading_stack[level] = heading
                para_level = level
            else:
                heading_stack[level] = heading
                heading_stack[level+1:] = ['#']*(6-level)
                para_level = level
        else:
            if line.strip() != '':
                buf.append(line)

        char_pos += len(line)+1

    if buf:
        flush_buf(char_pos)

    print('分段落算法完成，得到%d段落'%len(paragraphs))

    return paragraphs


def _chunk_paragraphs(paragraphs: list[dict], max_chunk_tokens: int, max_overlap_tokens: int) -> list[dict]:

    print('段落分块算法启动')

    cur_tokens = 0
    cur_para = []
    chunked_paragraphs: list[dict] = []

    def save_cur():
        heading = next(p['heading_path'] for p in reversed(cur_para))
        chunked_paragraphs.append({
            'heading_path': heading,
            'content': ''.join(p['content'] for p in cur_para),
            'start': cur_para[0]['start'],
            'end': cur_para[-1]['end'],
        })

    for para in paragraphs:
        if cur_tokens + approx_token(para['content']) < max_chunk_tokens:
            cur_para.append(para)
            cur_tokens += approx_token(para['content'])

        else:
            if len(cur_para) > 0 :
                save_cur()

            if max_overlap_tokens > 0:
                temp_tokens =0
                temp_para = []
                for p in reversed(cur_para):
                    if temp_tokens + approx_token(p['content']) < max_overlap_tokens:
                        temp_para.append(p)
                    else:
                        cur_para = [p for p in reversed(temp_para)]
                        cur_tokens = temp_tokens
            else:
                cur_para.clear()
                cur_tokens = 0

            cur_para.append(para)
            cur_tokens += approx_token(para['content'])

    if cur_para:
        save_cur()

    print('段落分块算法完成，得到%d块'%len(chunked_paragraphs))

    return chunked_paragraphs


def split_and_chunk_markdown(text: str, max_chunk_tokens: int, max_overlap_tokens: int) -> list[dict]:
    paragraphs = _split_markdown_to_paragraphs(text)
    chunked_paragraphs = _chunk_paragraphs(paragraphs, max_chunk_tokens, max_overlap_tokens)
    return chunked_paragraphs

def split_and_chunk_ordinary_text(text: str, max_chunk_tokens: int, max_overlap_tokens: int) -> list[dict]:
    paragraphs = []
    content = []
    curr_tokens = 0
    char_pos = 0
    for line in text.splitlines():
        line = line.strip()
        if line == '':
            continue
        if curr_tokens + approx_token(line) >= max_chunk_tokens and curr_tokens > 0:
            text = '\n'.join(content)
            paragraphs.append({
                'heading_path': content[0],
                'content': text,
                'start': char_pos - len(text),
                'end': char_pos
            })


            if max_overlap_tokens > 0:
                temp_tokens = 0
                temp_content = []
                for l in reversed(content):
                    if temp_tokens + approx_token(l) < max_overlap_tokens:
                        temp_content.append(l)
                    else:
                        content = [l for l in reversed(temp_content)]
                        break
            else:
                content.clear()

        content.append(line)
        char_pos += len(line)
        curr_tokens += approx_token(line)

    return paragraphs



def langchain_split_documents(documents, max_chunk_tokens: int):

    splitter = RecursiveCharacterTextSplitter(chunk_size=max_chunk_tokens)

    return splitter.split_documents(documents)










