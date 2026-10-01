---
page_title: "waf_type.app_firewall"
subcategory: ""
description: "waf_type.app_firewall for xcsh_virtual_host."
xcsh_docs: {"aliases": [], "body_bytes": 1137, "body_sha256": "sha256:1d310fc37b99043a80d7de118bf393adeb5998327ca1f21d204dba518ece4b3b", "canonical_id": "xcsh-docs:data-sources:virtual_host:properties:waf_type:app_firewall", "child_ids": ["xcsh-docs:data-sources:virtual_host:properties:waf_type:app_firewall:app_firewall"], "collection_id": "xcsh-docs:data-sources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:virtual_host:properties:waf_type:app_firewall", "parent_id": "xcsh-docs:data-sources:virtual_host:properties:waf_type", "path": "docs/guides/data-sources--virtual_host--properties--waf_type--app_firewall.md", "provider_name": "virtual_host", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["waf_type", "app_firewall"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/virtual_host/properties/waf_type/app_firewall/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "waf_type.app_firewall for xcsh_virtual_host.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# waf_type.app_firewall

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md)
- [Property reference](data-sources--virtual_host--reference.md)
- [waf_type](data-sources--virtual_host--properties--waf_type.md)
- waf_type.app_firewall

<a id="section"></a>

Type: `"single"`. Computed.

List of references to the app\_firewall configuration objects.

Upstream description:

A list of references to the app\_firewall configuration objects.

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

## Direct properties

- [app_firewall](data-sources--virtual_host--properties--waf_type--app_firewall--app_firewall.md): complete subsection reference.

## Next pages

- [waf_type.app_firewall.app_firewall](data-sources--virtual_host--properties--waf_type--app_firewall--app_firewall.md)
- [waf_type](data-sources--virtual_host--properties--waf_type.md)
- [xcsh_virtual_host](../data-sources/virtual_host.md)
