Build me a full-stack web app called **TextVault** that allows users to securely paste a text, lock it with a passphrase, and generate a shareable link. The recipient should only be able to view the text if they provide the correct key. Use the exact tech stack below:

---

### 🔷 FRONTEND

**Framework**: Next.js 14+ with App Router  
**UI Library**: shadcn/ui  
**Styling**: Tailwind CSS (integrated via shadcn)  
**Animations**: Framer Motion  
**Form Handling**: react-hook-form + zod  
**Encryption**: Use AES-GCM via Web Crypto API on the client-side.  
**Pages**:
1. `/` – Paste text, enter key, click "Lock & Generate Link", receive a unique link (`/view/[id]`)
2. `/view/[id]` – Input field for the key, click "Unlock", decrypt and view the message

**Frontend Behavior**:
- Encrypt the text before sending to backend
- Show error on incorrect key
- Use Framer Motion for animations: fade-ins, slide transitions, error shake on wrong key

---

### 🟡 BACKEND

**Framework**: Flask (Python)  
**Database**: SQLite  
**Endpoints**:
- `POST /api/lock`: Accepts `{ encrypted_text: string }`, stores it with a UUID and timestamp, returns `{ id }`
- `GET /api/unlock/<id>`: Returns the stored `{ encrypted_text }` for given ID

**Backend Notes**:
- No decryption on server
- Only store encrypted content
- Optionally include `created_at` and `expires_at` fields

---

### 🔧 PROJECT STRUCTURE


---

### 🚀 DEPLOYMENT

- Deploy frontend to **Vercel**
- Deploy Flask backend to **Render**
- Connect frontend to backend via absolute API URL (CORS-enabled)

---

Start coding now. Include working example for encryption/decryption with Web Crypto API, and connect frontend forms to Flask endpoints using fetch. Create the minimum necessary UI using shadcn and Tailwind. Animate transitions with Framer Motion.
