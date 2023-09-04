# ArgoCD ApplicationSet GitGen

This Github Action will create a YAML file residing in your targetted Github repository that will be used to generate the ArgoCD application with the ArgoCD application set Git File Generator coupled with Kustomize.

- More Info -> https://argocd-applicationset.readthedocs.io/en/stable/Generators-Git/#git-generator-files
- Kustomize -> https://kustomize.io/ 

## Example Usage
```

```

## Inputs
### Required Inputs
1. APPLICATION_NAME =

## Assumptions
- By default if namespace if not passed it will follow the name of the application.

## Expected Outcome from this Github Action
1. Create a `<filename>.yaml` file in the Github directory path you provided only when it doesn't exist.
2. The fields populated in the `<filename>.yaml` file will consists of
```
image: <Commonly pointing to your Container Repository>
imageTag: <Commonly pointing to the image store in your container registry>
kustomizePath: kustomize/template-api/dev
namespace: <Namespace where your application will be deployed in the K8s cluster>
```