---
page_title: "blocked_services.dns"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["blocked services dns"], "body_bytes": 1182, "body_sha256": "sha256:462ad6d86bf2943566fb9a06ba0433d2ed8c17f0403ab0c866de724edf5ec7d8", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:blocked_services:dns", "parent_id": "xcsh-docs:resources:fleet:properties:blocked_services", "path": "documentation/resources/fleet/properties/blocked_services/dns/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3221332113111220-3010302110013032-3213220312323020-1000113322202101-2022231333313222-1303323113231210-0322133011123230-2012313233132100", "registry_path": "docs/guides/resources--fleet--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["blocked_services", "dns"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/blocked_services/dns/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# blocked_services.dns

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- [blocked_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/blocked_services/)
- blocked_services.dns

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
dns = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [blocked_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/blocked_services/)
- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
