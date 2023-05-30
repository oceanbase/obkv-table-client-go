# How to contribute

OBKV Table Client is a community-driven open source project and we welcome all the contributors. Contributions to the ODP project are expected to adhere to our [Code of Conduct](CODE_OF_CONDUCT.md).

## Before you contribute

Before you contribute, please click the **Sign in with Gitub to agree button** to sign the CLA.

What is [CLA](https://en.wikipedia.org/wiki/Contributor_License_Agreement)?

## Contribution guide

Please follow these steps to create your Pull Request to this repository.

> **Note:**
>
> This section takes creating a PR to the `master` branch as an example. The steps of creating PRs for other branches are similar.

### Step 1: Fork the repository

1. Visit the project (link TODO)
2. Click the **Fork** button to establish an online fork.

### Step 2: Clone your fork to local

```bash
# Define your working directory
working_dir=$HOME/Workspace

# Configure GitHub
user={your Github profile name}

# Create your clone
mkdir -p $working_dir
cd $working_dir
git clone link TODO

# Add upstream
git remote add upstream link TODO

# Set no push for the upstream master
git remote set-url --push upstream no_push

# Confirm your remote setting
git remote -v
```

### Step 3: Create a new branch

1. Get your local master up-to-date with the upstream/master.

    ```bash
    cd $working_dir/docs
    git fetch upstream
    git checkout main
    git rebase upstream/main
    ```

2. Create a new branch based on the master branch.

    ```bash
    git checkout -b new-branch-name
    ```

### Step 4: Develop

Edit some file(s) on the `new-branch-name` branch and save your changes.

### Step 5: Commit your changes

```bash
git status # Checks the local status
git add <file> ... # Adds the file(s) you want to commit. If you want to commit all changes, you can directly use `git add.`
git commit -m "commit-message: update the xx"
```

### Step 6: Keep your branch in sync with upstream/master

```bash
# While on your new branch
git fetch upstream
git rebase upstream/main
```

### Step 7: Push your changes to the remote

```bash
git push -u origin new-branch-name # "-u" is used to track the remote branch from the origin
```

### Step 8: Create a pull request

1. Visit your fork at <https://github.com/$user/docs> (replace `$user` with your GitHub account).
2. Click the `Compare & pull request` button next to your `new-branch-name` branch to create your PR.