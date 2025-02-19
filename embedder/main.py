from emb import app
import dotenv
import os

if __name__ == '__main__':
    dotenv.load_dotenv('.env.embedder')
    PORT = os.getenv('PORT')
    app.run(port=PORT, debug=True)
