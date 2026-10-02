---
page_title: "default_bot_setting"
subcategory: "Security"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["default bot setting"], "body_bytes": 1202, "body_sha256": "sha256:459188c4b3f6ca51bba1c84ff72b78dbd0d54511e7687528d66d82df650775e6", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_firewall:properties:default_bot_setting", "parent_id": "xcsh-docs:resources:app_firewall:reference", "path": "documentation/resources/app_firewall/properties/default_bot_setting/index.md", "product": "distributed-cloud", "provider_name": "app_firewall", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1010333202130203-1311300012320212-2220112222022320-0231332302121133-3213002003102111-1011013332021222-3323022012320133-0320311213330220", "registry_path": "docs/guides/resources--app_firewall--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["default_bot_setting"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_firewall/properties/default_bot_setting/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["app_firewallCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_bot_setting

Breadcrumbs:

- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/)
- default_bot_setting

<a id="section"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for default bot setting. Defaults to \`map\[\]\`. Server applies default
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
default_bot_setting = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/)
- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/)
