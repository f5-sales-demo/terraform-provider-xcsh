---
page_title: "custom_network_config.active_network_policies"
subcategory: ""
description: "custom_network_config.active_network_policies for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 1466, "body_sha256": "sha256:ba693f006a3abe38efda4a29fac969f14f8f72c1ebf90ab99ed168451ce15538", "canonical_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:active_network_policies", "child_ids": ["xcsh-docs:resources:voltstack_site:properties:custom_network_config:active_network_policies:network_policies"], "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:active_network_policies", "parent_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config", "path": "docs/guides/resources--voltstack_site--properties--custom_network_config--active_network_policies.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_network_config", "active_network_policies"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/custom_network_config/active_network_policies/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_network_config.active_network_policies for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# custom_network_config.active_network_policies

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md)
- [Property reference](resources--voltstack_site--reference.md)
- [custom_network_config](resources--voltstack_site--properties--custom_network_config.md)
- custom_network_config.active_network_policies

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for active network policies.

Upstream description:

List of firewall policy views.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("network_policies")}
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
active_network_policies {
  # Configure direct properties listed below.
}
```

## Direct properties

- [network_policies](resources--voltstack_site--properties--custom_network_config--active_network_policies--network_policies.md): complete subsection reference.

## Next pages

- [custom_network_config.active_network_policies.network_policies](resources--voltstack_site--properties--custom_network_config--active_network_policies--network_policies.md)
- [custom_network_config](resources--voltstack_site--properties--custom_network_config.md)
- [xcsh_voltstack_site](../resources/voltstack_site.md)
