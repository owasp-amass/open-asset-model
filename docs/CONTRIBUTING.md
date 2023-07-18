# Contributing

Thank you for your interest in contributing to the amass ecosystem!
Open Asset Model is only as good as the community that's backing it.
Contributions could simply be a bug or suggestion you drop in an issue,
dropping a question in [Discord](https://discord.gg/HNePVyX3cp),
or tweaking a line (or character) or two in the project as a pull request.
[Open issues][issue-url] is another good place to look for ways to contribute.

If you need any help or guidance, please reach out
to the community on [Discord](https://discord.gg/HNePVyX3cp).

## Getting Started

Open Asset Model uses a fork & pull model for contributions.
This means that you will need to fork this repository using
the *Fork* button in the top right of the page.

1. Once you have forked the repository, you'll need to clone it.
   The following commands will clone the repository to your local
   machine and allow Go to resolve the import paths correctly.
   This will set owasp-amass as your origin remote and your fork
   as a secondary remote.

```bash
git clone https://github.com/owasp-amass/open-asset-model $GOPATH/src/github.com/owasp-amass/open-asset-model
git remote add [github-user] https://github.com/[github-user]/open-asset-model
```

1. Create a new branch for your contribution.

```bash
git checkout --track origin/develop
git pull origin develop
git checkout -b [fix-or-improvement-name]
```

2. Make the changes that you wish to contribute.

3. Commit your changes to your branch.

```bash
git add <updated files>
git commit -m "A short description of the changes"
```

4. Push your branch to your fork

```bash
git push [github-user] [fix-or-improvement-name]
```

5. Create a pull request into the `develop` branch of `owasp-amass/open-asset-model`

### Notes regarding contributions

- All commits should be pushed to your fork and be submitted as
  pull requests into the `develop` branch of `owasp-amass/open-asset-model`.
- Force pushing is not allowed. If you need to update your
  pull request, simply make an additional commit.
- Ensure that your pull request is up-to-date with the `develop` branch
  before submitting.


[issue-url]: https://github.com/owasp-amass/open-asset-model/issues