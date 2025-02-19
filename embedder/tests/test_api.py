import pytest

from emb import app

@pytest.fixture(scope='session')
def server():
    with app.test_client() as client:
        yield client

def test_embed(server):
    data = {'text': 'Hello, world!'}
    response = server.post('/embed', json=data)

    assert response.status_code == 200
    assert 'embeddings' in response.json
    embeddings = response.json['embeddings']
    assert len(embeddings) == 1
    assert len(embeddings[0]) == 384

def test_embed_list(server):
    data = {'text': ['Hello, world!', 'Goodbye, world!']}
    response = server.post('/embed', json=data)

    assert response.status_code == 200
    assert 'embeddings' in response.json
    embeddings = response.json['embeddings']
    assert len(embeddings) == 2