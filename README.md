# HubOrator

Austrian Open Science Cloud - Hub Operator

## What it does

It creates CRDs for Kubespawner Profiles to read in JupyterHub.

- [KubeSpawner ProfileList](https://jupyterhub-kubespawner.readthedocs.io/en/latest/spawner.html#kubespawner.KubeSpawner.profile_list)

## Sample

```yaml
apiVersion: hub.austrianopensciencecloud.org/v1alpha1
kind: Profile
metadata:
  name: pokeball-notebook
spec: 
  slug: pokeball-notebook
  allowed_groups:
    - lect1:instructor
  display_name: "Pokeballs!!"
  description: "highly dangerous, use with care"
  profile_options:
    image:
      display_name: Image
      choices:
        v0.0.1:
          display_name: Pokeballs v0.0.1
          kubespawner_override:
            image: adlsregistrysbx.azurecr.io/pokeball/fortran:0.0.1
        v0.0.2:
          display_name: Pokeballs v0.0.2
          kubespawner_override:
            image: adlsregistrysbx.azurecr.io/pokeball/fortran:0.0.2
        v0.0.3:
          display_name: Pokeballs v0.0.3
          kubespawner_override:
            image: adlsregistrysbx.azurecr.io/pokeball/fortran:0.0.3
    ressources:
      display_name: Ressources
      choices:
        '2':
          display_name: 2 vCPU & 2G Memory
          kubespawner_override:
            cpu_limit: 2
            mem_limit: 2G
        '4':
          display_name: 4 vCPU & 4G Memory
          kubespawner_override:
            cpu_limit: 4
            mem_limit: 4G
        '6':
          display_name: 6 vCPU & 6G Memory
          kubespawner_override:
            cpu_limit: 6
            mem_limit: 6G

```