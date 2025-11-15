FROM node:20-alpine

WORKDIR /app

# Copia package.json e package-lock.json
COPY package*.json ./
RUN npm ci

# Copia o código
COPY . .

# Build da aplicação Next.js
RUN npm run build

EXPOSE 3000

ENV NODE_ENV=production
ENV PORT=3000

# Usa npm start que roda o servidor Next.js em produção
CMD ["npm", "start"]