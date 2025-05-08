import json
from confluent_kafka import Producer

producer = Producer({
    'bootstrap.servers': 'kafka:9092'
})

def send_validation_result(id, name, is_valid):
    result = {
        'id': id,
        'name': name,
        'valid': is_valid
    }
    try:
        producer.produce(
            topic='document-validation-result',
            value=json.dumps(result).encode('utf-8'),
            key=str(id).encode('utf-8')
        )
        producer.flush()
        print("Resultado enviado com sucesso.")
    except Exception as e:
        print(f"Erro ao enviar resultado: {e}")