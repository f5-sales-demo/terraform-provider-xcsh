---
page_title: "default_anonymization"
subcategory: "Security"
description: "default_anonymization for xcsh_app_firewall."
xcsh_docs: {"aliases": [], "body_bytes": 903, "body_sha256": "sha256:0300acc97ae3bee9fc230e3805cc0594b52a6b8937d85425829ba0cc61d62149", "canonical_id": "xcsh-docs:resources:app_firewall:properties:default_anonymization", "child_ids": [], "collection_id": "xcsh-docs:resources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_firewall:properties:default_anonymization", "parent_id": "xcsh-docs:resources:app_firewall:reference", "path": "docs/guides/resources--app_firewall--properties--default_anonymization.md", "provider_name": "app_firewall", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["default_anonymization"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_firewall/properties/default_anonymization/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "default_anonymization for xcsh_app_firewall.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_firewallCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# default_anonymization

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md)
- [Property reference](resources--app_firewall--reference.md)
- default_anonymization

<a id="section"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for default anonymization. Defaults to \`map\[\]\`. Server applies default
when omitted.

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
default_anonymization = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--app_firewall--reference.md)
- [xcsh_app_firewall](../resources/app_firewall.md)
