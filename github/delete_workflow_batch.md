## Delete GitHub workflows

GitHub allows you to delete each GitHub Action run separately. Removing all of them at once is not possible in UI. As a workaround, the GitHub API can be used to automate the workflow deletion.

<!--more-->

This can be necessary, if you want to make you private repo, publicly available but your Action executions have sensitive data stored in logs.

You can use the [GitHub CLI](https://cli.github.com/) to make API calls.

```bash
# 1. Export organization and repository name
export OWNER=mszostok
export REPO=aaas

# 2. List workflows
gh api -X GET /repos/$OWNER/$REPO/actions/workflows | jq '.workflows[] | .name,.id'

# 3. Copy the ID of the workflow you want to clear and set it
export WORKFLOW_ID=12478829

# 4. [SANITY CHECK] list runs that will be removed
gh api -X GET /repos/$OWNER/$REPO/actions/workflows/$WORKFLOW_ID/runs --paginate | jq '.workflow_runs[] | .id'

# 5. Delete all runs for a given workflow
gh api -X GET /repos/$OWNER/$REPO/actions/workflows/$WORKFLOW_ID/runs --paginate | jq '.workflow_runs[] | .id' | xargs -I{} gh api -X DELETE /repos/$OWNER/$REPO/actions/runs/{}  --silent
```
