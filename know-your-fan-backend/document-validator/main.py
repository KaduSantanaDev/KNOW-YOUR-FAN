import base64
import io
import json

import pytesseract
from confluent_kafka import Consumer, Producer
from PIL import Image
import producer

c = Consumer({
    'bootstrap.servers': 'kafka:9092',
    'group.id': 'test',
})

c.subscribe(['document-validation'])



while True:
    msg = c.poll(1.0)   
    if msg is None:
        continue    
    if msg.error():
        print(f"❌ Erro do consumidor: {msg.error()}")
        continue    
    raw_value = msg.value()
    try:
        data = json.loads(raw_value.decode('utf-8'))
    except json.JSONDecodeError as e:
        print(f"❌ Erro ao decodificar JSON: {e}")
        continue
    id = data.get('id')    
    name = data.get('name')
    document_b64 = data.get('document') 
    if not name or not document_b64:
        print("❌ Mensagem inválida (faltando 'name' ou 'document')")
        continue    
    print(f"👤 Nome: {name}")   
    
    try:
        image_data = base64.b64decode(document_b64)
        image = Image.open(io.BytesIO(image_data))
        
    except Exception as e:
        print(f"❌ Erro ao decodificar imagem: {e}")
        continue   
     
    try:
        extracted_text = pytesseract.image_to_string(image)
        
        if name.lower() in extracted_text.lower():
            producer.send_validation_result(id, name, True)
            continue
        
    except Exception as e:
        print(f"❌ Erro no OCR: {e}")
        continue    
    c.commit(msg)
    print("✅ Offset comitado com sucesso.")
c.close()

