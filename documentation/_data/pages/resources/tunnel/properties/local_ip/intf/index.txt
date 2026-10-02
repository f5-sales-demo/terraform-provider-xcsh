---
page_title: "local_ip.intf"
subcategory: ""
description: "Provides the local interface to pick up source IP and network for transporting encapsulated packet."
xcsh_docs: {"aliases": ["local ip intf"], "body_bytes": 1430, "body_sha256": "sha256:46d01543e5b792a6eb0398c83b85e2435778830714b628a1b4d6e500088d1031", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:tunnel:properties:local_ip:intf:local_intf"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:resources:tunnel:properties:local_ip:intf", "parent_id": "xcsh-docs:resources:tunnel:properties:local_ip", "path": "documentation/resources/tunnel/properties/local_ip/intf/index.md", "product": "distributed-cloud", "provider_name": "tunnel", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2223230302300101-0122121121132103-2321020002221033-3320003233333123-3313223101333111-2000011230222230-0300113212001320-1033032031000213", "registry_path": "docs/guides/resources--tunnel--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["local_ip", "intf"], "schema_version": 1, "sections": [{"aliases": ["local intf"], "anchor": "section", "description": "Local interface to be used for filling in source information of IP and network for transport.", "document_id": "xcsh-docs:resources:tunnel:properties:local_ip:intf:local_intf", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["local_ip", "intf", "local_intf"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tunnel/properties/local_ip/intf/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Provides the local interface to pick up source IP and network for transporting encapsulated packet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tunnelCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_ip.intf

Breadcrumbs:

- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/)
- [local_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/)
- local_ip.intf

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Provides the local interface to pick up source IP and network for transporting encapsulated packet.

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
intf {
  # Configure direct properties listed below.
}
```

## Direct properties

- [local_intf](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/intf/local_intf/): complete subsection reference.

## Next pages

- [local_ip.intf.local_intf](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/intf/local_intf/)
- [local_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/)
- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/)
