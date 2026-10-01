---
page_title: "Resource"
subcategory: "Security"
description: "Resource for xcsh_forward_proxy_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1160, "body_sha256": "sha256:fe9132fed545eba17b1b7c887c7d760b7f397a6671e18673a0437add4a1bc816", "canonical_id": "xcsh-docs:resources:forward_proxy_policy:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:forward_proxy_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:bcc594a880ce1beadc720b2bd1c76066016eb6d5489e15cf1b1295e1e8d40584", "source_path": "examples/resources/xcsh_forward_proxy_policy/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:forward_proxy_policy:example:resource", "parent_id": "xcsh-docs:resources:forward_proxy_policy:examples", "path": "docs/guides/resources--forward_proxy_policy--example--resource.md", "provider_name": "forward_proxy_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/forward_proxy_policy/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_forward_proxy_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["forward_proxy_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md)
- [Examples](resources--forward_proxy_policy--examples.md)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_forward_proxy_policy/resource.tf`; digest `sha256:bcc594a880ce1beadc720b2bd1c76066016eb6d5489e15cf1b1295e1e8d40584`.

```terraform
# ForwardProxyPolicy Resource Example
# Manages a Forward Proxy Policy resource in F5 Distributed Cloud for forward proxy policy specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ForwardProxyPolicy configuration
resource "xcsh_forward_proxy_policy" "example" {
  name      = "example-forward-proxy-policy"
  namespace = "staging"
}
```

## Next pages

- [Examples](resources--forward_proxy_policy--examples.md)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md)
