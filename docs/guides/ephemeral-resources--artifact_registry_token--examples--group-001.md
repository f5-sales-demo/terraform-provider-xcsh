---
page_title: "xcsh_artifact_registry_token examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_artifact_registry_token examples."
---

# xcsh_artifact_registry_token examples

<a id="canonical-0011313231110131-0332120013220332-0111300222332010-0023102120300323-0123110301313201-3222121230031110-3232221322111321-3103110203301113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_artifact_registry_token](../ephemeral-resources/artifact_registry_token.md#canonical-1323301000023223-3012100013003333-1031230331233031-2332031031030022-0033001320123233-0303101023323201-3331031322333331-2330103332212301)
- Examples

<a id="canonical-0001122022233211-1332133330001333-0210022123200300-0312322032112123-1313211322302301-3301201333332110-0223021322021131-2301022121231322"></a>

### Complete configurations for `xcsh_artifact_registry_token`

- [Ephemeral](ephemeral-resources--artifact_registry_token--examples--group-001.md#canonical-0331000100300321-1112003032322321-3321000221031110-0311121332232201-2031031331021200-3333230213330010-2031301221220022-0123102213131101): valid configuration.

<a id="canonical-0331000100300321-1112003032322321-3321000221031110-0311121332232201-2031031331021200-3333230213330010-2031301221220022-0123102213131101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Ephemeral example

Breadcrumbs:

- [xcsh_artifact_registry_token](../ephemeral-resources/artifact_registry_token.md#canonical-1323301000023223-3012100013003333-1031230331233031-2332031031030022-0033001320123233-0303101023323201-3331031322333331-2330103332212301)
- [Examples](ephemeral-resources--artifact_registry_token--examples--group-001.md#canonical-0011313231110131-0332120013220332-0111300222332010-0023102120300323-0123110301313201-3222121230031110-3232221322111321-3103110203301113)
- Ephemeral

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/ephemeral-resources/xcsh_artifact_registry_token/ephemeral.tf`; digest `sha256:f91e2c9b88b3ed39e41f1e34edb4d5addb5289872f1ac18a74ddd0ff3f1aec6f`.

```terraform
# ArtifactRegistryToken EphemeralResource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

ephemeral "xcsh_artifact_registry_token" "example" {
  namespace = "example-value"
}
```
