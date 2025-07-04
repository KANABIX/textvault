from sqlalchemy import Column, String, DateTime
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker
import uuid
from datetime import datetime, timedelta

Base = declarative_base()

class VaultEntry(Base):
    __tablename__ = 'vault_entries'
    id = Column(String, primary_key=True, default=lambda: str(uuid.uuid4()))
    encrypted_text = Column(String, nullable=False)
    created_at = Column(DateTime, default=datetime.utcnow)
    expires_at = Column(DateTime, nullable=True)

# Database setup
engine = create_engine('sqlite:///vault.db')
SessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=engine)

def init_db():
    Base.metadata.create_all(bind=engine)
