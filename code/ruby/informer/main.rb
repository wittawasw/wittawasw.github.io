require 'informers'

ip = Informers::PreTrainedTokenizer

require_relative 'patch'

# Informers::PreTrainedTokenizer.test

sentences = ["This is a ภาษาไทย", "ทดสอบประโยค"]

model = Informers.pipeline("embedding", "sentence-transformers/paraphrase-MiniLM-L6-v2")
embeddings = model.(sentences, normalize: false)

# print(embeddings)

# Console.test
