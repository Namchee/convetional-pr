# Conventional PR

[![Go Version Badge](https://img.shields.io/github/go-mod/go-version/namchee/conventional-pr)](https://github.com/Namchee/conventional-pr) [![Go Report Card](https://goreportcard.com/badge/github.com/Namchee/conventional-pr)](https://goreportcard.com/report/github.com/Namchee/conventional-pr) [![codecov](https://codecov.io/gh/Namchee/conventional-pr/branch/master/graph/badge.svg)](https://codecov.io/gh/Namchee/conventional-pr) <!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->
[![All Contributors](https://img.shields.io/badge/all_contributors-8-orange.svg?style=flat-square)](#contributors-)
<!-- ALL-CONTRIBUTORS-BADGE:END --> 

Conventional PR is a GitHub Action that validates all pull requests sent to a GitHub-hosted repository.

Conventional PR aims to ease your burden in managing your GitHub-hosted repository by validating, marking, even moderating low-quality attempts of pull request just by integrating Conventional PR to your existing CI/CD workflow.

## Features

- ✨ Configurable, tune Conventional PR easily to suit your needs.
- 💡 Sensible defaults, validates pull request with out-of-the-box sensibility while doesn't generate too much noise.
- ♿ Whitelisting, validates pull request that actually matters.
- 📈 Transparent reporting, see what Conventional PR is actually doing.

## Usage

> To use Conventional PR, you'll need to prepare a GitHub access token. Please refer to this [article](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token) on how to generate an access token.

You can integrate Conventional PR to your existing CI/CD workflow by using `Namchee/conventional-pr@<version>` in one of your jobs using the YAML syntax.

Below is the example of creating a new Conventional PR workflow in your CI/CD workflow.

```yml
on:
  pull_request:
    # You can add specific pull request types too.
    # If this is omitted, the workflow will only executed on opened, synchronize, and reopened
    # which is enough for generic use cases.
    # Below is the list of supported pull request sub-events
    # types: [opened, reopened, ready_for_review, unlocked, synchronize]

jobs:
  cpr:
    runs-on: ubuntu-latest
    steps:
      - name: Validate the pull request
        uses: Namchee/conventional-pr@v(version)
        with:
          access_token: YOUR_GITHUB_ACCESS_TOKEN_HERE
```

Please refer to [GitHub workflow syntax](https://docs.github.com/en/free-pro-team@latest/actions/reference/workflow-syntax-for-github-actions#about-yaml-syntax-for-workflows) for more advanced usage.

> Access token is **required**. Please generate one or use `${{ secrets.GITHUB_TOKEN }}` as your access token and the `github-actions` bot will run the job for you. Do note that the `github-actions` bot has [more limited functionalities](#caveats)

## Whitelist

Whitelist is one of the features of `conventional-pr` that allows any pull requests that satisfies the whitelist criterion to be marked as stable regardless whether the pull request itself determined as valid by the actual validation logic or not.

A pull request will be whitelisted and marked as valid if the pull request satisfies **one or more** enabled whitelisting criteria.

The following are all available whitelists criteria in `conventional-pr`.

### Pull request is still on the draft phase

<table>
  <tr>
    <th>Default</th>
    <td>Enabled</td>
  </tr>
  <tr>
    <th>Input</th>
    <td><code>draft</code></td>
  </tr>
</table>

This whitelist checks if the pull request status is `draft`. If the pull request status is `draft`, it will be marked as stable for the time being.

Do note that once the pull request has been marked as ready for review, the workflow will re-trigger the validation flow unless by default.

### Pull request is created by a bot

<table>
  <tr>
    <th>Default</th>
    <td>Enabled</td>
  </tr>
  <tr>
    <th>Input</th>
    <td><code>bot</code></td>
  </tr>
</table>

This whitelist checks if the pull request author is a bot account. If the author is a bot account, the pull request will be marked as valid.

Do note that this option only recognized accounts that are officialy recognized as a GitHub bot account and not an user-automated accounts. For this usecase, please whitelist the username instead.

### Pull request is created by an administrators

<table>
  <tr>
    <th>Default</th>
    <td>Disabled</td>
  </tr>
  <tr>
    <th>Input</th>
    <td><code>strict</code></td>
  </tr>
</table>

This whitelist checks if the pull request author has administrator privileges in the current repository. If the author has administrator privileges in the repository, the pull request will be marked as valid.

If set to `false`, any pull requests made by repository administrators will automatically be marked as valid.

### Pull request is created by a whitelisted user

<table>
  <tr>
    <th>Default</th>
    <td>Disabled</td>
  </tr>
  <tr>
    <th>Input</th>
    <td><code>ignored_users</code></td>
  </tr>
</table>

This whitelist checks if the pull request author username is in the list of names provided in the `ignored_users` input. If the author is in the list, the pull request will be marked as valid.

The list of names must be provided in a comma-separated GitHub username string.

## Validator
 
Validator is the core feature of Conventional PR. Validator will validate any pull requests that triggers the workflow according to a specified criteria.

A pull request is considered to be valid if it satisfies **all** enabled validation flow.

The following are all available whitelists criteria in `conventional-pr`.

### Pull request has a valid title

<table>
  <tr>
    <th>Default</th>
    <td>Enabled</td>
  </tr>
  <tr>
    <th>Input</th>
    <td><code>title_pattern</code></td>
  </tr>
</table>

This validator checks if the pull request title satisfies the regular expression provided in the `title_pattern` input.

The regular expression must be provided in Perl syntax. Defaults to [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/) message structure.

Do note that this validator only validates the pull request title and does not validate any titles in the pull request body.

Filling the input with an empty string will disable this validator.

### Pull request has a body

<table>
  <tr>
    <th>Default</th>
    <td>Enabled</td>
  </tr>
  <tr>
    <th>Input</th>
    <td><code>body</code></td>
  </tr>
</table>

This validator checks if the pull request has a non-empty body text.

### All commits have valid message

<table>
  <tr>
    <th>Default</th>
    <td>Disabled</td>
  </tr>
  <tr>
    <th>Input</th>
    <td><code>commit_pattern</code></td>
  </tr>
</table>

This validator checks if all commits on the pull request satisfies the regular expression provided in the `commit_pattern` input. The regular expression must be provided in Perl syntax.

Filling the input with an empty string will disabled this validator.

### The pull request have a valid branch name

<table>
  <tr>
    <th>Default</th>
    <td>Disabled</td>
  </tr>
  <tr>
    <th>Input</th>
    <td><code>branch_pattern</code></td>
  </tr>
</table>

This validator checks if the branch name of the pull request satisfies the regular expression provided in the `branch_pattern` input. The regular expression must be provided in Perl syntax.

Filling the input with an empty string will disabled this validator.

### Pull request links to an issue

<table>
  <tr>
    <th>Default</th>
    <td>Enabled</td>
  </tr>
  <tr>
    <th>Input</th>
    <td><code>issue</code></td>
  </tr>
</table>

This validator checks if the pull request is linked one or more issues in the same or different repository either by [directly linking issue through the sidebar](https://docs.github.com/en/issues/tracking-your-work-with-issues/using-issues/linking-a-pull-request-to-an-issue#manually-linking-a-pull-request-or-branch-to-an-issue-using-the-issue-sidebar) or by using [magic closing keywords](https://docs.github.com/en/issues/tracking-your-work-with-issues/using-issues/linking-a-pull-request-to-an-issue#linking-a-pull-request-to-an-issue-using-a-keyword).

> [!NOTE]
> Ensure that the access token you provide can access issues on the target repository.

### Pull request does not introduce too many changes

<table>
  <tr>
    <th>Default</th>
    <td>Enabled</td>
  </tr>
  <tr>
    <th>Input</th>
    <td><code>maximum_changes</code></td>
  </tr>
</table>

This validator checks if the pull request introduces too many changes to the base branch. A pull request that changes more files than the predetermined value will be considered as invalid.

Filling the input with zero will disable this validator.

## Inputs

You can customize this actions with these following options (fill it on `with` section):

| **Name**          | **Required?** | **Default Value**                       | **Description**                                                                                                                                                                                                                                                                                                                                                                                                                   |
| ----------------- | ------------- | --------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `access_token`    | `true`        |                                         | [GitHub access token](https://docs.github.com/en/free-pro-team@latest/github/authenticating-to-github/creating-a-personal-access-token) to interact with the GitHub API. It is recommended to store this token with [GitHub Secrets](https://docs.github.com/en/free-pro-team@latest/actions/reference/encrypted-secrets). **To support automatic close, labeling, and comment report, please grant a write access to the token** |
| `close`           | `false`       | `false`                                 | Immediately close invalid pull request.                                                                                                                                                                                                                                                                                                                                                                                           |
| `message`         | `false`       | `''`                                    | Extra message to be posted when the pull request is invalid.                                                                                                                                                                                                                                                                                                                                                                      |
| `label`           | `false`       | `''`                                    | Invalid pull requests label. Fill with an empty string to disable labeling.                                                                                                                                                                                                                                                                                                                                                       |
| `draft`           | `false`       | `true`                                  | Skip pull request validation if the pull request is a draft.                                                                                                                                                                                                                                                                                                                                                                      |
| `strict`          | `false`       | `true`                                  | Enforce validation rules to repository administrators.                                                                                                                                                                                                                                                                                                                                                                            |
| `bot`             | `false`       | `true`                                  | Skip pull request validation if the author is a bot.                                                                                                                                                                                                                                                                                                                                                                              |
| `title_pattern`   | `false`       | `([\w\-]+)(\([\w\-]+\))?!?: [\w\s:\-]+` | Valid pull request title regex pattern in Perl syntax. Defaults to the [conventional commit style](https://www.conventionalcommits.org/en/v1.0.0/) commit messages. Fill with an empty string to disabled pull request title validation.                                                                                                                                                                                          |
| `commit_pattern`  | `false`       | `''`                                    | Valid pull request commit messages regex pattern in Perl syntax. Fill with an empty string to disabled commit message validation.                                                                                                                                                                                                                                                                                                 |
| `branch_pattern`  | `false`       | `''`                                    | Valid pull request branch name regex pattern in Perl syntax. Fill with an empty string to disabled branch name validation.                                                                                                                                                                                                                                                                                                        |
| `body`            | `false`       | `true`                                  | Require all pull request to have a non-empty body.                                                                                                                                                                                                                                                                                                                                                                                |
| `issue`           | `false`       | `true`                                  | Require all pull request to reference an existing issue.                                                                                                                                                                                                                                                                                                                                                                          |
| `maximum_changes` | `false`       | `0`                                     | Limits file changes per one pull request. Fill with zero to disable this feature.                                                                                                                                                                                                                                                                                                                                                 |
| `ignored_users`   | `false`       | `''`                                    | GitHub usernames to be whitelisted from pull request validation. Must be a comma-separated string. Example: `Namchee, foo, bar` will bypass pull request validation for users `Namchee`, `foo`, `bar`. Case-sensitive.                                                                                                                                                                                                            |
| `verbose`         | `false`       | `false`                                 | Post validation report on every pull request validation flow.                                                                                                                                                                                                                                                                                                                                                                     |
| `edit`            | `false`       | `false`                                 | Edit existing validation report instead of submitting a new comment. Does not do anything if `verbose` is `false`.                                                                                                                                                                                                                                                                                                                |

## Supported Events

Ideally, Conventional PR workflow should only triggered when an event related to pull requests is fired. However, Conventional PR is only able to function properly with some pull request sub-events and will ignore the rest. Below is the list of supported [GitHub Action pull request sub-events](https://docs.github.com/en/actions/reference/events-that-trigger-workflows):

| **Name**           | **Triggered On**                                                                           |
| ------------------ | ------------------------------------------------------------------------------------------ |
| `opened`           | When a new pull request is submitted.                                                      |
| `reopened`         | When a closed pull request is re-opened, either by the original author or by someone else. |
| `ready_for_review` | When a draft pull request is finished and transformed into normal pull request.            |
| `sychronize`       | When a pull request has changes on its history, such as pushing new commits to the branch. |
| `unlocked`         | When a pull request is unlocked.                                                           |

> Omitting `types` is enough for generic use-cases since omitting it will cause the workflow to be triggered on
`opened`, `reopened`, and `synchronize` events.

## Caveats

- `edit` feature cannot be used with `github-actions` credentials as it [doesn't have the `user` scope](https://docs.github.com/en/actions/security-guides/automatic-token-authentication#permissions-for-the-github_token). If you want to use the `edit` feature, please generate a [personal access token](https://docs.github.com/en/free-pro-team@latest/github/authenticating-to-github/creating-a-personal-access-token) instead with `read:user` scope. `edit` **cannot** be combined with `report`.
- It is recommended to use a dedicated dummy account if you want to use the `edit` feature as it may lead to unintended edits.

## Forked Repository

Conventional PR is designed to be executed on internal environment, where only authorized users are allowed create pull request. With that design philosophy, Conventional PR is not designed to be executed on a forked repository in mind. Moreover, granting a token with write access to unauthorized user may lead to [GitHub repository access exploit via GitHub Action](https://securitylab.github.com/research/github-actions-preventing-pwn-requests/).  However, this limitation proves to be discouraging for open source project management. As an open-source project itself, it becomes a hinderance if Conventional PR cannot be used to manage itself.

To circumvent this issue, you must change the [event target](https://docs.github.com/en/actions/using-workflows/events-that-trigger-workflows) from `pull_request` to `pull_request_target` which changes the execution context from the fork to the base repository. Below is the example of action configuration using `pull_request_target`

```yml
on:
  pull_request_target:

jobs:
  cpr:
    runs-on: ubuntu-latest
    steps:
      - name: Validates the pull request
        uses: Namchee/conventional-pr@v(version)
        with:
          access_token: YOUR_GITHUB_ACCESS_TOKEN_HERE
```

Do note that `pull_request_target` allows unsafe code to be executed from the head of the pull request that could alter your repository or steal any secrets you use in your repository. Avoid using `pull_request_event` if you need to build or run code from the pull request.

## Contributors ✨

Thanks goes to these wonderful people ([emoji key](https://allcontributors.org/docs/en/emoji-key)):
<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tbody>
    <tr>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/Namchee"><img src="https://avatars.githubusercontent.com/u/32661241?v=4?s=100" width="100px;" alt="Cristopher"/><br /><sub><b>Cristopher</b></sub></a><br /><a href="https://github.com/Namchee/conventional-pr/commits?author=Namchee" title="Code">💻</a> <a href="https://github.com/Namchee/conventional-pr/issues?q=author%3ANamchee" title="Bug reports">🐛</a> <a href="https://github.com/Namchee/conventional-pr/commits?author=Namchee" title="Documentation">📖</a> <a href="#example-Namchee" title="Examples">💡</a> <a href="#ideas-Namchee" title="Ideas, Planning, & Feedback">🤔</a></td>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/smutel"><img src="https://avatars.githubusercontent.com/u/12967891?v=4?s=100" width="100px;" alt="Samuel Mutel"/><br /><sub><b>Samuel Mutel</b></sub></a><br /><a href="https://github.com/Namchee/conventional-pr/issues?q=author%3Asmutel" title="Bug reports">🐛</a></td>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/BinToss"><img src="https://avatars.githubusercontent.com/u/7243190?v=4?s=100" width="100px;" alt="Noah Sherwin"/><br /><sub><b>Noah Sherwin</b></sub></a><br /><a href="https://github.com/Namchee/conventional-pr/issues?q=author%3ABinToss" title="Bug reports">🐛</a> <a href="#ideas-BinToss" title="Ideas, Planning, & Feedback">🤔</a></td>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/avorima"><img src="https://avatars.githubusercontent.com/u/15158349?v=4?s=100" width="100px;" alt="Mario Valderrama"/><br /><sub><b>Mario Valderrama</b></sub></a><br /><a href="#ideas-avorima" title="Ideas, Planning, & Feedback">🤔</a></td>
      <td align="center" valign="top" width="14.28%"><a href="https://th3s4mur41.me"><img src="https://avatars.githubusercontent.com/u/5543595?v=4?s=100" width="100px;" alt="Michaël Vanderheyden"/><br /><sub><b>Michaël Vanderheyden</b></sub></a><br /><a href="https://github.com/Namchee/conventional-pr/issues?q=author%3ATh3S4mur41" title="Bug reports">🐛</a></td>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/berkon"><img src="https://avatars.githubusercontent.com/u/28842661?v=4?s=100" width="100px;" alt="Bernd Konnerth"/><br /><sub><b>Bernd Konnerth</b></sub></a><br /><a href="https://github.com/Namchee/conventional-pr/issues?q=author%3Aberkon" title="Bug reports">🐛</a></td>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/TheGoderGuy"><img src="https://avatars.githubusercontent.com/u/26481335?v=4?s=100" width="100px;" alt="TheGoderGuy"/><br /><sub><b>TheGoderGuy</b></sub></a><br /><a href="https://github.com/Namchee/conventional-pr/commits?author=TheGoderGuy" title="Code">💻</a></td>
    </tr>
    <tr>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/Lapotor"><img src="https://avatars.githubusercontent.com/u/17144397?v=4?s=100" width="100px;" alt="Lapotor"/><br /><sub><b>Lapotor</b></sub></a><br /><a href="#ideas-Lapotor" title="Ideas, Planning, & Feedback">🤔</a></td>
    </tr>
  </tbody>
</table>

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://allcontributors.org) specification.
Contributions of any kind are welcome!

## License

This project is licensed under the [MIT License](./LICENSE)
