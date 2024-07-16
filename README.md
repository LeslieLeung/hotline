# hotline

A client-side workflow engine/executor for AI applications.

## Demo

### Define a workflow(like GitHub Actions)

```yaml
workflows:
  - name: translate
    id: translate
    inputs:
      - name: text
        required: true
    steps:
      # Define a step that uses the `dify_workflow` action
      - name: translation
        id: translation
        uses: dify_workflow
        with:
          host: https://dify.ai
          api_key: app-ATotallyFakeApiKey
          # The input section defined in Dify workflow
          inputs:
            - name: text
              value: ${{ inputs.text }}
      # Define a step that prints the result
      - name: print
        uses: print
        with:
          # Reference the result from the previous step
          message: ${{ steps.translation.outputs.text }}
```

### Execute the workflow

```bash
hotline exec translate -f translate.yaml --text="Good morning, how can I help you?"
```