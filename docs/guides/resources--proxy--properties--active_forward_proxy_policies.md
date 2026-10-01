---
page_title: "active_forward_proxy_policies"
subcategory: ""
description: "active_forward_proxy_policies for xcsh_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 1754, "body_sha256": "sha256:6b9f175a093ee26549a1371d3fc95e4f14f095def89a06f5645e59aa7ad57fd0", "canonical_id": "xcsh-docs:resources:proxy:properties:active_forward_proxy_policies", "child_ids": ["xcsh-docs:resources:proxy:properties:active_forward_proxy_policies:forward_proxy_policies"], "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:active_forward_proxy_policies", "parent_id": "xcsh-docs:resources:proxy:reference", "path": "docs/guides/resources--proxy--properties--active_forward_proxy_policies.md", "provider_name": "proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["active_forward_proxy_policies"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/active_forward_proxy_policies/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "active_forward_proxy_policies for xcsh_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# active_forward_proxy_policies

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md)
- [Property reference](resources--proxy--reference.md)
- active_forward_proxy_policies

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: active\_forward\_proxy\_policies, no\_forward\_proxy\_policy; Default:
no\_forward\_proxy\_policy\] Ordered List of Forward Proxy Policies active.

Upstream description:

Ordered List of Forward Proxy Policies active.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("forward_proxy_policies")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

OneOf alternatives in this subsection:

- [active_forward_proxy_policies](resources--proxy--properties--active_forward_proxy_policies.md#section)
- [no_forward_proxy_policy](resources--proxy--properties--no_forward_proxy_policy.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
active_forward_proxy_policies {
  # Configure direct properties listed below.
}
```

## Direct properties

- [forward_proxy_policies](resources--proxy--properties--active_forward_proxy_policies--forward_proxy_policies.md): complete subsection reference.

## Next pages

- [active_forward_proxy_policies.forward_proxy_policies](resources--proxy--properties--active_forward_proxy_policies--forward_proxy_policies.md)
- [Property reference](resources--proxy--reference.md)
- [xcsh_proxy](../resources/proxy.md)
