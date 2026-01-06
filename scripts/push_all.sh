#!/usr/bin/env bash
set -euo pipefail

# ===== Config =====
REMOTE_1="github"
REMOTE_2="gitee"

# ===== Helpers =====
die() { echo "❌ $*" >&2; exit 1; }

# Ensure we are inside a git repo
git rev-parse --is-inside-work-tree >/dev/null 2>&1 || die "当前目录不是 Git 仓库。请在仓库根目录运行。"

REPO_ROOT="$(git rev-parse --show-toplevel)"
cd "$REPO_ROOT"

# Fix 'dubious ownership' in dev containers / bind mounts
git config --global --add safe.directory "$REPO_ROOT" >/dev/null 2>&1 || true

# Determine branch
BRANCH="$(git rev-parse --abbrev-ref HEAD)"

# Commit message
MSG="${1:-}"
if [[ -z "$MSG" ]]; then
  MSG="chore: update $(date '+%Y-%m-%d %H:%M:%S')"
fi

echo "📦 Repo:   $REPO_ROOT"
echo "🌿 Branch: $BRANCH"
echo "📝 Msg:    $MSG"
echo

# Stage changes
git add -A

# If nothing to commit, still push (in case upstream behind)
if git diff --cached --quiet; then
  echo "ℹ️  没有检测到可提交的变更（working tree clean）。"
else
  git commit -m "$MSG"
  echo "✅ Commit done."
fi

# Push to both remotes
# Ensure remotes exist
git remote get-url "$REMOTE_1" >/dev/null 2>&1 || die "找不到远端 '$REMOTE_1'，请先 git remote -v 检查。"
git remote get-url "$REMOTE_2" >/dev/null 2>&1 || die "找不到远端 '$REMOTE_2'，请先 git remote -v 检查。"

echo
echo "🚀 Pushing to $REMOTE_1 ($BRANCH)..."
git push "$REMOTE_1" "$BRANCH"

echo
echo "🚀 Pushing to $REMOTE_2 ($BRANCH)..."
git push "$REMOTE_2" "$BRANCH"

echo
echo "🎉 All done."
