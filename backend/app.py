from flask import Flask, request, jsonify
from flask_cors import CORS
from models import VaultEntry, SessionLocal, init_db
from datetime import datetime, timedelta
import uuid

app = Flask(__name__)
CORS(app)

init_db()

@app.route('/api/lock', methods=['POST'])
def lock_text():
    data = request.get_json()
    encrypted_text = data.get('encrypted_text')
    if not encrypted_text:
        return jsonify({'error': 'Missing encrypted_text'}), 400

    db = SessionLocal()
    entry = VaultEntry(
        encrypted_text=encrypted_text,
        # Optionally set expires_at here, e.g. expires_at=datetime.utcnow() + timedelta(days=7)
    )
    db.add(entry)
    db.commit()
    db.refresh(entry)
    db.close()
    return jsonify({'id': entry.id})

@app.route('/api/unlock/<id>', methods=['GET'])
def unlock_text(id):
    db = SessionLocal()
    entry = db.query(VaultEntry).filter(VaultEntry.id == id).first()
    db.close()
    if not entry:
        return jsonify({'error': 'Not found'}), 404
    return jsonify({'encrypted_text': entry.encrypted_text})

if __name__ == '__main__':
    app.run(debug=True)
