import torch
from flask import Flask, request, jsonify
from flask_cors import CORS

from .model import Embedder
from . import utils

app = Flask(__name__)
CORS(app)


device = utils.get_device()
print('Loading model onto {}'.format(device))
model = Embedder(device=device)


@app.route('/embed', methods=['POST'])
def embed():
    data = request.json
    text = data['text']
    with torch.no_grad():
        embeddings = model(text).cpu().numpy().tolist()
    data = {'embedding': embeddings[0]}
    return jsonify(data)

@app.route('/embed_batch', methods=['POST'])
def embed_batch():
    data = request.json
    texts = data['texts']
    with torch.no_grad():
        embeddings = model(texts).cpu().numpy().tolist()
    data = {'embeddings': embeddings}
    return jsonify(data)
