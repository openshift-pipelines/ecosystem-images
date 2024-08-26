# `.tekton` definitions

These are the `Pipeline` (and other Tekton Objects) used to build the different ecosystem images.
They are initially coming from the template of konflux-ci *but* we maintain them ourselves. This means that we need to make sure they are up-to-date (renovate bot from konflux does this) and valid to be able to release through the Red Hat konflux instance.

- [`image-pipeline.yaml`](./image-pipeline.yaml) is the main pipeline that all images should use.
- Each image has 2 `PipelineRun`, one for pull request and one on push.
