import torch
import torch.nn as nn
from transformers import AutoModel, AutoTokenizer


class Embedder(nn.Module):
    def __init__(self, model_name="sentence-transformers/all-MiniLM-L6-v2", device='cpu'):
        super(Embedder, self).__init__()
        self.device = device
        self.to(device)
        self.embedder = AutoModel.from_pretrained(model_name)
        self.tokenizer = AutoTokenizer.from_pretrained(model_name)

    def forward(self, text):
        input_ids = self.tokenizer(text, return_tensors='pt', padding=True, truncation=True)['input_ids']
        input_ids = input_ids.to(self.device)
        output = self.embedder(input_ids)
        return output.last_hidden_state[:, 0, :]
