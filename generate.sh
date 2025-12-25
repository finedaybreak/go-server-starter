#!/bin/bash

# Usage:
#   ./generate.sh <module_name>          - Create module files
#   ./generate.sh -d <module_name>       - Delete module files
# Example:
#   ./generate.sh user
#   ./generate.sh -d user

# Parse arguments
DELETE_MODE=false
MODULE_NAME=""

if [ "$1" = "-d" ]; then
    DELETE_MODE=true
    MODULE_NAME=$2
else
    MODULE_NAME=$1
fi

if [ -z "$MODULE_NAME" ]; then
    echo "Usage:"
    echo "  ./generate.sh <module_name>          - Create module files"
    echo "  ./generate.sh -d <module_name>       - Delete module files"
    echo ""
    echo "Example:"
    echo "  ./generate.sh user"
    echo "  ./generate.sh -d user"
    exit 1
fi
# Convert module name to snake_case
# Handles: user-role -> user_role, userRole -> user_role, UserRole -> user_role
MODULE_NAME_SNAKE=$(echo "$MODULE_NAME" | sed 's/-/_/g' | sed 's/\([A-Z]\)/_\1/g' | sed 's/^_//' | tr '[:upper:]' '[:lower:]')
MODULE_NAME_LOWER=$MODULE_NAME_SNAKE
MODULE_NAME_UPPER=$(echo "$MODULE_NAME_SNAKE" | awk -F_ '{for(i=1;i<=NF;i++) $i=toupper(substr($i,1,1)) tolower(substr($i,2))}1' OFS="")

# Define file paths
MODEL_FILE="internal/model/${MODULE_NAME_LOWER}.go"
REPO_FILE="internal/repo/${MODULE_NAME_LOWER}_repo.go"
DTO_FILE="internal/dto/${MODULE_NAME_LOWER}_dto.go"
SERVICE_FILE="internal/service/${MODULE_NAME_LOWER}_service.go"
HANDLER_FILE="internal/handler/${MODULE_NAME_LOWER}_handler.go"
EXCEPTION_FILE="internal/exception/${MODULE_NAME_LOWER}_exception.go"

# ============================================
# DELETE MODE
# ============================================
if [ "$DELETE_MODE" = true ]; then
    echo "[delete] Deleting module: $MODULE_NAME"

    if [ -f "$MODEL_FILE" ]; then
        rm "$MODEL_FILE"
        echo "[ok] Deleted: $MODEL_FILE"
    fi

    if [ -f "$REPO_FILE" ]; then
        rm "$REPO_FILE"
        echo "[ok] Deleted: $REPO_FILE"
    fi

    if [ -f "$DTO_FILE" ]; then
        rm "$DTO_FILE"
        echo "[ok] Deleted: $DTO_FILE"
    fi

    if [ -f "$SERVICE_FILE" ]; then
        rm "$SERVICE_FILE"
        echo "[ok] Deleted: $SERVICE_FILE"
    fi

    if [ -f "$HANDLER_FILE" ]; then
        rm "$HANDLER_FILE"
        echo "[ok] Deleted: $HANDLER_FILE"
    fi

    if [ -f "$EXCEPTION_FILE" ]; then
        rm "$EXCEPTION_FILE"
        echo "[ok] Deleted: $EXCEPTION_FILE"
    fi

    echo ""
    echo "Module deletion completed!"
    exit 0
fi

# ============================================
# CREATE MODE
# ============================================
echo "🚀 Generating module: $MODULE_NAME"

# Create directories if they don't exist
mkdir -p internal/model
mkdir -p internal/repo
mkdir -p internal/service
mkdir -p internal/handler
mkdir -p internal/dto
mkdir -p internal/exception

# ============================================
# 1. Model
# ============================================
if [ ! -f "$MODEL_FILE" ]; then
    cat > "$MODEL_FILE" << EOF
package model
EOF
    echo "[ok] Created: $MODEL_FILE"
else
    echo "[warn] Skipped: $MODEL_FILE (already exists)"
fi

# ============================================
# 2. repo Interface & Implementation
# ============================================
# Function to generate repo template content
generate_repo_content() {
    cat << EOF
package repo

import (
	"go-server-starter/internal/model"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ${MODULE_NAME_UPPER}Repo interface {
	BaseRepo[model.${MODULE_NAME_UPPER}]
	WithTx(tx *gorm.DB) ${MODULE_NAME_UPPER}Repo
}

type ${MODULE_NAME_UPPER}RepoImpl struct {
	BaseRepo[model.${MODULE_NAME_UPPER}]
	db     *gorm.DB
	logger *zap.Logger
}

func New${MODULE_NAME_UPPER}Repo(db *gorm.DB, logger *zap.Logger) ${MODULE_NAME_UPPER}Repo {
	return &${MODULE_NAME_UPPER}RepoImpl{
		BaseRepo: NewBaseRepo[model.${MODULE_NAME_UPPER}](db, logger),
		db:       db,
		logger:   logger,
	}
}

func (r *${MODULE_NAME_UPPER}RepoImpl) WithTx(tx *gorm.DB) ${MODULE_NAME_UPPER}Repo {
	return &${MODULE_NAME_UPPER}RepoImpl{
		BaseRepo: NewBaseRepo[model.${MODULE_NAME_UPPER}](tx, r.logger),
		db:       tx,
		logger:   r.logger,
	}
}
EOF
}

if [ ! -f "$REPO_FILE" ]; then
    generate_repo_content > "$REPO_FILE"
    echo "[ok] Created: $REPO_FILE"
else
    # Check if file has more than 2 lines
    LINE_COUNT=$(wc -l < "$REPO_FILE" | tr -d ' ')
    if [ "$LINE_COUNT" -le 2 ]; then
        generate_repo_content > "$REPO_FILE"
        echo "[ok] Updated: $REPO_FILE (template code inserted)"
    else
        echo "[warn] Skipped: $REPO_FILE (already has content)"
    fi
fi

# ============================================
# 3. DTO (Data Transfer Objects)
# ============================================
if [ ! -f "$DTO_FILE" ]; then
    cat > "$DTO_FILE" << EOF
package dto
EOF
    echo "[ok] Created: $DTO_FILE"
else
    echo "[warn] Skipped: $DTO_FILE (already exists)"
fi

# ============================================
# 4. Service Interface & Implementation
# ============================================
if [ ! -f "$SERVICE_FILE" ]; then
    cat > "$SERVICE_FILE" << EOF
package service
EOF
    echo "[ok] Created: $SERVICE_FILE"
else
    echo "[warn] Skipped: $SERVICE_FILE (already exists)"
fi

# ============================================
# 5. Handler (HTTP Layer)
# ============================================
if [ ! -f "$HANDLER_FILE" ]; then
    cat > "$HANDLER_FILE" << EOF
package handler
EOF
    echo "[ok] Created: $HANDLER_FILE"
else
    echo "[warn] Skipped: $HANDLER_FILE (already exists)"
fi

# ============================================
# 6. Exception
# ============================================
if [ ! -f "$EXCEPTION_FILE" ]; then
    cat > "$EXCEPTION_FILE" << EOF
package exception
EOF
    echo "[ok] Created: $EXCEPTION_FILE"
else
    echo "[warn] Skipped: $EXCEPTION_FILE (already exists)"
fi

echo ""
echo "Module generation completed!"
echo ""
echo " Generated files:"
echo "   - $MODEL_FILE"
echo "   - $REPO_FILE"
echo "   - $DTO_FILE"
echo "   - $SERVICE_FILE"
echo "   - $HANDLER_FILE"
echo "   - $EXCEPTION_FILE"
echo ""
echo " Next steps:"
echo "   1. Implement TODO items in each file"
echo "   2. Register handler in your router"
echo "   3. Add migration for the model"
echo "   4. Run tests"
