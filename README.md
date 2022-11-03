# version-action
GitHub Action to generate versions with the date

## Outputs

## `version`

human-readable version tag containing the creation date.

Example: `v2022-11-03-878e`

## Example usage
```
jobs:
  test:
    runs-on: ubuntu-latest
    name: Test version action
    steps:
      - name: Generate new tag version
        id: version
        uses: flynshue/version-action

      - name: Print version from action
        run: |
          echo "${{ steps.version.outputs.version }}"
```