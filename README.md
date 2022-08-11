# HubOrator

Austrian Open Science Cloud - Hub Operator

## What it does

It creates CRDs for Kubespawner Profiles to read in JupyterHub: [hub.austrianopensciencecloud.org_profiles.yaml](/config/crd/bases/hub.austrianopensciencecloud.org_profiles.yaml)

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

## Hot to use in JupyterHub

```python
import os
from kubespawner.clients import shared_client

## see: https://discourse.jupyter.org/t/tailoring-spawn-options-and-server-configuration-to-certain-users/8449#solution-to-problem-1-4
async def options_form(spawner):
  auth_state = await spawner.user.get_auth_state()

  group = 'hub.austrianopensciencecloud.org'
  version = 'v1alpha1'
  namespace = os.getenv('POD_NAMESPACE')
  plural = 'profiles'

  profile_list = []
  if len(spawner.profile_list) > 0:
    profile_list = spawner.profile_list

  api_client = shared_client('CustomObjectsApi')
  groups = auth_state.get("groups", [])
  #name = auth_state.get("name", "USER")
  #spawner.log.info(f"{name}s options_form groups: {groups}")

  ## see: https://github.com/kubernetes-client/python/blob/master/kubernetes/docs/CustomObjectsApi.md
  try:
    res = await api_client.list_namespaced_custom_object(group, version, namespace, plural)
    for item in res['items']:
      allowed_groups = item['spec'].get("allowed_groups", False)

      if not allowed_groups and item['spec'] not in profile_list:
        profile_list.append(item['spec'])
        continue

      if any(group in groups for group in allowed_groups) and item['spec'] not in profile_list:
        profile_list.append(item['spec'])
  
  except Exception as e:
    spawner.log.info('### Error in options_form')
    spawner.log.error(e)

  spawner.profile_list = profile_list
  return spawner._options_form_default()

c.KubeSpawner.options_form = options_form
```