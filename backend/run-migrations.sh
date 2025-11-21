
set -e  


GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m' 


DB_HOST="${DB_HOST:-localhost}"
DB_PORT="${DB_PORT:-15432}"
DB_USER="${DB_USER:-chathadmin}"
DB_PASSWORD="${DB_PASSWORD:-1239012}"
DB_NAME="${DB_NAME:-chatdb}"


DB_URL="postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"

MIGRATIONS_PATH="./migrations"

echo -e "${YELLOW}🔄 Iniciando migrations...${NC}"


if command -v pg_isready &> /dev/null; then

    if ! PGPASSWORD=$DB_PASSWORD pg_isready -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER" &> /dev/null; then
        echo -e "${RED} Banco de dados não está acessível!${NC}"
        exit 1
    fi
else

    echo -e "${YELLOW} pg_isready não encontrado, verificando via Docker...${NC}"
    if ! docker exec db-chatwebsocket pg_isready -U "$DB_USER" -d "$DB_NAME" &> /dev/null; then
        echo -e "${RED} Banco de dados não está acessível!${NC}"
        echo -e "${YELLOW}Certifique-se que o container 'db-chatwebsocket' está rodando${NC}"
        exit 1
    fi
fi

echo -e "${GREEN}✅ Banco acessível!${NC}"


echo -e "${YELLOW}⬆️  Aplicando migrations...${NC}"
if migrate -path "$MIGRATIONS_PATH" -database "$DB_URL" up; then
    echo -e "${GREEN} Migrations aplicadas com sucesso!${NC}"
else
    echo -e "${RED} Erro ao aplicar migrations${NC}"
    exit 1
fi

# Mostra versão atual
VERSION=$(migrate -path "$MIGRATIONS_PATH" -database "$DB_URL" version 2>&1 | tail -n 1)
echo -e "${GREEN}📊 Versão atual: ${VERSION}${NC}"