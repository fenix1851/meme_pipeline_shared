#!/bin/bash

# Получаем последнюю версию тега
LATEST_VERSION=$(gh release list --limit 1 --json tagName --jq '.[0].tagName')

# Убираем префикс "v"
LATEST_VERSION=${LATEST_VERSION#v}

# Разбиваем версию на части (major, minor, patch)
IFS='.' read -r MAJOR MINOR PATCH <<< "$LATEST_VERSION"

# Инкрементируем patch-версию
PATCH=$((PATCH + 1))

# Собираем новую версию
NEW_VERSION="v$MAJOR.$MINOR.$PATCH"

# Получаем последний commit message
LAST_COMMIT_MESSAGE=$(git log -1 --pretty=%B | tr '\n' ' ')  # Убираем переносы строк

# Формируем заголовок и описание релиза
TITLE="Release $NEW_VERSION $LAST_COMMIT_MESSAGE"
NOTES="$LAST_COMMIT_MESSAGE"

# Создаем новый релиз
gh release create "$NEW_VERSION" --title "$TITLE" --notes "$NOTES"

echo "New release created: $NEW_VERSION"
