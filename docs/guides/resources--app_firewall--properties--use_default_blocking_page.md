---
page_title: "use_default_blocking_page"
subcategory: "Security"
description: "use_default_blocking_page for xcsh_app_firewall."
xcsh_docs: {"aliases": [], "body_bytes": 884, "body_sha256": "sha256:90a285a02474c2b308ce23b883b20919e765d555ba88766aa908c9e58928a97e", "canonical_id": "xcsh-docs:resources:app_firewall:properties:use_default_blocking_page", "child_ids": [], "collection_id": "xcsh-docs:resources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_firewall:properties:use_default_blocking_page", "parent_id": "xcsh-docs:resources:app_firewall:reference", "path": "docs/guides/resources--app_firewall--properties--use_default_blocking_page.md", "provider_name": "app_firewall", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["use_default_blocking_page"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_firewall/properties/use_default_blocking_page/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "use_default_blocking_page for xcsh_app_firewall.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_firewallCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# use_default_blocking_page

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md)
- [Property reference](resources--app_firewall--reference.md)
- use_default_blocking_page

<a id="section"></a>

Type: `["object", {}]`. Optional, Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

Upstream description:

This can be used for messages where no values are needed.

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
use_default_blocking_page = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--app_firewall--reference.md)
- [xcsh_app_firewall](../resources/app_firewall.md)
