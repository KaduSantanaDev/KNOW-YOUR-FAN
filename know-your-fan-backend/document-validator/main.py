import base64
import io
import json
from PIL import Image
import pytesseract
from confluent_kafka import Consumer, KafkaException
from sentence_transformers import SentenceTransformer, util

# Configurações do Kafka
conf = {
    'bootstrap.servers': 'localhost:9092',
    'group.id': 'document-checker',
    'auto.offset.reset': 'earliest'
}

consumer = Consumer(conf)
consumer.subscribe(['document-validation'])

model = SentenceTransformer('paraphrase-MiniLM-L6-v2')

print("🧠 Aguardando mensagens Kafka...")

try:
    while True:
        msg = consumer.poll(1.0)
        if msg is None:
            continue
        if msg.error():
            raise KafkaException(msg.error())

        # Carregar e interpretar mensagem
        data = json.loads(msg.value().decode('utf-8'))
        name = data.get('Name') or data.get('name')
        document_b64 = data.get('Document') or data.get('document')

        if not name or not document_b64:
            print("❌ Mensagem inválida (faltando nome ou documento)")
            continue

        # Decodifica imagem do base64
        try:
            image = Image.open(io.BytesIO(base64.b64decode(document_b64)))
        except Exception as e:
            print(f"❌ Erro ao decodificar imagem: {e}")
            continue

        # OCR com pytesseract
        extracted_text = pytesseract.image_to_string(image)
        print(f"📄 Texto extraído:\n{extracted_text.strip()}")

        # Similaridade com IA
        score = util.cos_sim(
            model.encode(name),
            model.encode(extracted_text)
        ).item()

        if score > 0.6:
            print(f"✅ Nome '{name}' encontrado no documento (score={score:.2f})")
        else:
            print(f"❌ Nome '{name}' NÃO encontrado (score={score:.2f})")

except KeyboardInterrupt:
    print("🛑 Encerrando consumer...")
finally:
    consumer.close()
