---
page_title: "waf_type"
subcategory: ""
description: "waf_type for xcsh_virtual_host."
xcsh_docs: {"aliases": [], "body_bytes": 1407, "body_sha256": "sha256:da52e398d2bfae85429c4658f6dc9c1e444d88ae84d982af105ebf903b41f6b9", "canonical_id": "xcsh-docs:data-sources:virtual_host:properties:waf_type", "child_ids": ["xcsh-docs:data-sources:virtual_host:properties:waf_type:app_firewall", "xcsh-docs:data-sources:virtual_host:properties:waf_type:disable_waf", "xcsh-docs:data-sources:virtual_host:properties:waf_type:inherit_waf"], "collection_id": "xcsh-docs:data-sources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:virtual_host:properties:waf_type", "parent_id": "xcsh-docs:data-sources:virtual_host:reference", "path": "docs/guides/data-sources--virtual_host--properties--waf_type.md", "provider_name": "virtual_host", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["waf_type"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/virtual_host/properties/waf_type/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "waf_type for xcsh_virtual_host.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# waf_type

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md)
- [Property reference](data-sources--virtual_host--reference.md)
- waf_type

<a id="section"></a>

Type: `"single"`. Computed.

WAF instance will be pointing to an app\_firewall object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ref_type": "[\"app_firewall\",\"disable_waf\",\"inherit_waf\"]"
}
```

## Direct properties

- [app_firewall](data-sources--virtual_host--properties--waf_type--app_firewall.md): complete subsection reference.

- [disable_waf](data-sources--virtual_host--properties--waf_type--disable_waf.md): complete subsection reference.

- [inherit_waf](data-sources--virtual_host--properties--waf_type--inherit_waf.md): complete subsection reference.

## Next pages

- [waf_type.app_firewall](data-sources--virtual_host--properties--waf_type--app_firewall.md)
- [waf_type.disable_waf](data-sources--virtual_host--properties--waf_type--disable_waf.md)
- [waf_type.inherit_waf](data-sources--virtual_host--properties--waf_type--inherit_waf.md)
- [Property reference](data-sources--virtual_host--reference.md)
- [xcsh_virtual_host](../data-sources/virtual_host.md)
