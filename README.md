# Dependabot Autoapproval

This repository contains a GitHub Action that automatically approves pull requests raised by Dependabot, depending on whether the library is contained in a whitelist. It also allows you to provide a reason for whitelisting a library and add comments to the pull requests.

## Purpose

When using Dependabot to manage your project's dependencies, it can be helpful to automate the process of approving and merging pull requests that update these dependencies. The Dependabot Autoapproval GitHub Action provided in this repository allows you to streamline this process by automatically approving pull requests for libraries that are included in a predefined whitelist.

## Usage

To use the Dependabot Autoapproval GitHub Action in your repository, follow these steps:

1. Create a `.github/workflows/autoapproval.yml` file in your repository with the [Content](https://github.com/ying-jeanne/dependabot_autoapproval/blob/main/.github/workflows/dependabot_reviewer.yml)
2. Configure the whitelist in the .github/workflows/whitelist.json file. Specify the libraries that should be auto-approved, along with optional reasons for whitelisting. For example:
```json
  [
    {
      "name": "github.com/gin-gonic/gin",
      "reason": "The library is only used in test"
    },
    {
      "name": "github.com/google/uuid",
      "reason": "The library is stable"
    }
  ]
```
Customize the library names, reasons to your project's requirements.

Commit and push the changes to your repository. The Dependabot Autoapproval GitHub Action will be triggered when a pull request is opened by Dependabot, automatically approving it if the library is present in the whitelist.

## Contributing
Contributions are welcome! If you find any issues or have suggestions for improvements, please feel free to open an issue or submit a pull request.

## License
This repository is licensed under the MIT License.
