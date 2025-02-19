import pytest

from emb.model import Embedder


@pytest.fixture(scope='session')
def model():
    return Embedder(device='cpu')


def test_string(model):
    text = "Hello, world!"
    embeddings = model(text)
    assert embeddings.shape == (1, 384)

def test_list_of_strings(model):
    texts = ["Hello, world!", "Goodbye, world!"]
    embeddings = model(texts)
    assert embeddings.shape == (2, 384)

def test_list_of_one_string(model):
    texts = ["Hello, world!"]
    embeddings = model(texts)
    assert embeddings.shape == (1, 384)
