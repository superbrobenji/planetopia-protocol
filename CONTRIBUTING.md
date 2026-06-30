# Contributing

## Repository setup

After the first merge to `main`, configure branch protection in GitHub settings (Settings → Branches → Add rule for `main`):

- Require status checks to pass before merging
  - Required checks: `Test` (go-test job), `Header sync` (header-sync job)
- Require approvals: 1
- Dismiss stale pull request approvals when new commits are pushed
- Do not allow bypassing the above settings
