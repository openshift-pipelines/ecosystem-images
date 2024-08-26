# ecosystem-images

All the images required for the ecosystem and built by Konflux. As of today it holds:
- [`opc`](./opc/) from [`openshift-pipelines/opc`](https://github.com/openshift-pipelines/opc)
- [`git-init`](./git-init) from [`tektoncd-catalog/git-clone`](https://github.com/tektoncd-catalog/git-clone)

The source code of each of these images are synced from their respective upstream repositories, and the `Dockerfile` are generated.

All of this is build using the Red Hat [`konflux-ci`][konfluxci] instance, and most of it's configuration is [here](./.konflux).
As [`konflux-ci`][konfluxci] uses [Tekton][tekton] and [`pipelines-as-code`][pac], the "CI" definition is available [here](./.tekton).

## Branch management

## rpms.lock.yaml

Update this file the [rpm-lockfile](https://github.com/konflux-ci/rpm-lockfile-prototype) tool, with:

```
$ rpm-lockfile-prototype --image $(grep ubi-minimal opc/Dockerfile | awk '{print [}') --outfile rpms.lock.yaml rpms.in.yaml']}')
```

[konfluxci]: https://konflux-ci.dev/
[tekton]: https://tekton.dev
[pac]: https://pipelinesascode.com/
