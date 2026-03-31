from langchain_unstructured import UnstructuredLoader
from langchain_opendataloader_pdf import OpenDataLoaderPDFLoader

def unstructured_load_document(
    paths: list[str],
):
    loader = UnstructuredLoader(
        file_path=paths
    )

    return loader.load()


def opendataloader_load_document(
    paths: list[str],
):
    loader = OpenDataLoaderPDFLoader(
        file_path=paths
    )
    return loader.load()


if __name__ == "__main__":
    for document in opendataloader_load_document(["/Users/yuanji/Desktop/example/uv.pdf"]):
        print('-'*80)
        print(document)

