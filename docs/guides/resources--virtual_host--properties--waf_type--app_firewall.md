---
page_title: "waf_type.app_firewall"
subcategory: ""
description: "waf_type.app_firewall for xcsh_virtual_host."
xcsh_docs: {"aliases": [], "body_bytes": 1388, "body_sha256": "sha256:4f4fde9e80611cbc2a60d1fd5cb758ea68f926ef7d659553d3b199406c6509f1", "canonical_id": "xcsh-docs:resources:virtual_host:properties:waf_type:app_firewall", "child_ids": ["xcsh-docs:resources:virtual_host:properties:waf_type:app_firewall:app_firewall"], "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:waf_type:app_firewall", "parent_id": "xcsh-docs:resources:virtual_host:properties:waf_type", "path": "docs/guides/resources--virtual_host--properties--waf_type--app_firewall.md", "provider_name": "virtual_host", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["waf_type", "app_firewall"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/waf_type/app_firewall/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "waf_type.app_firewall for xcsh_virtual_host.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# waf_type.app_firewall

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md)
- [Property reference](resources--virtual_host--reference.md)
- [waf_type](resources--virtual_host--properties--waf_type.md)
- waf_type.app_firewall

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

List of references to the app\_firewall configuration objects.

Upstream description:

A list of references to the app\_firewall configuration objects.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("app_firewall")}
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
app_firewall {
  # Configure direct properties listed below.
}
```

## Direct properties

- [app_firewall](resources--virtual_host--properties--waf_type--app_firewall--app_firewall.md): complete subsection reference.

## Next pages

- [waf_type.app_firewall.app_firewall](resources--virtual_host--properties--waf_type--app_firewall--app_firewall.md)
- [waf_type](resources--virtual_host--properties--waf_type.md)
- [xcsh_virtual_host](../resources/virtual_host.md)
