---
page_title: "local_ip.ip_address.auto"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["local ip ip address auto"], "body_bytes": 1318, "body_sha256": "sha256:6f07d138aed410c86a255df8f6e963c69633fe743153511187a3c1f142573a9d", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:auto", "parent_id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address", "path": "documentation/resources/tunnel/properties/local_ip/ip_address/auto/index.md", "product": "distributed-cloud", "provider_name": "tunnel", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2310113003120102-2123032023112022-1033021130133210-3202032303213133-0320330211322033-0121101031222013-0020231300333130-1032320033221121", "registry_path": "docs/guides/resources--tunnel--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["local_ip", "ip_address", "auto"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tunnel/properties/local_ip/ip_address/auto/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tunnelCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_ip.ip_address.auto

Breadcrumbs:

- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/)
- [local_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/)
- [local_ip.ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/ip_address/)
- local_ip.ip_address.auto

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
auto = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [local_ip.ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/ip_address/)
- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/)
