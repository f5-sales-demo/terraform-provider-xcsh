---
page_title: "custom_network_config.active_forward_proxy_policies"
subcategory: ""
description: "custom_network_config.active_forward_proxy_policies for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 1570, "body_sha256": "sha256:d0fde3efeb6e04e3ca302cd6997468ff3e52090bdc2b10afeae1b7c4696fb4cf", "canonical_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:active_forward_proxy_policies", "child_ids": ["xcsh-docs:resources:voltstack_site:properties:custom_network_config:active_forward_proxy_policies:forward_proxy_policies"], "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:active_forward_proxy_policies", "parent_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config", "path": "docs/guides/resources--voltstack_site--properties--custom_network_config--active_forward_proxy_policies.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_network_config", "active_forward_proxy_policies"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/custom_network_config/active_forward_proxy_policies/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_network_config.active_forward_proxy_policies for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config.active_forward_proxy_policies

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md)
- [Property reference](resources--voltstack_site--reference.md)
- [custom_network_config](resources--voltstack_site--properties--custom_network_config.md)
- custom_network_config.active_forward_proxy_policies

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
active_forward_proxy_policies {
  # Configure direct properties listed below.
}
```

## Direct properties

- [forward_proxy_policies](resources--voltstack_site--properties--custom_network_config--active_forward_proxy_policies--forward_proxy_policies.md): complete subsection reference.

## Next pages

- [custom_network_config.active_forward_proxy_policies.forward_proxy_policies](resources--voltstack_site--properties--custom_network_config--active_forward_proxy_policies--forward_proxy_policies.md)
- [custom_network_config](resources--voltstack_site--properties--custom_network_config.md)
- [xcsh_voltstack_site](../resources/voltstack_site.md)
