---
page_title: "primary.rr_set_group"
subcategory: "DNS"
description: "Create and manage set groups, and resource record sets within them, x-VES-I/O-managed set is managed by F5."
xcsh_docs: {"aliases": ["primary rr set group"], "body_bytes": 2281, "body_sha256": "sha256:365cd17cc3c1228743a193167b74e26a1b2f7ddd5b05583282419c79ccc0e2fc", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:data-sources:dns_zone:properties:primary:rr_set_group:metadata", "xcsh-docs:data-sources:dns_zone:properties:primary:rr_set_group:rr_set"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_zone:properties:primary:rr_set_group", "parent_id": "xcsh-docs:data-sources:dns_zone:properties:primary", "path": "documentation/data-sources/dns_zone/properties/primary/rr_set_group/index.md", "product": "distributed-cloud", "provider_name": "dns_zone", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202", "registry_path": "docs/guides/data-sources--dns_zone--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["primary", "rr_set_group"], "schema_version": 1, "sections": [{"aliases": ["metadata"], "anchor": "section", "description": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:data-sources:dns_zone:properties:primary:rr_set_group:metadata", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["primary", "rr_set_group", "metadata"], "syntax": "attribute", "type": "object"}, {"aliases": ["rr set"], "anchor": "section", "description": "Collection of DNS resource record sets.", "document_id": "xcsh-docs:data-sources:dns_zone:properties:primary:rr_set_group:rr_set", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["primary", "rr_set_group", "rr_set"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_zone/properties/primary/rr_set_group/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Create and manage set groups, and resource record sets within them, x-VES-I/O-managed set is managed by F5.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# primary.rr_set_group

Breadcrumbs:

- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/)
- [primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/)
- primary.rr_set_group

<a id="section"></a>

Type: `"list"`. Computed.

Create and manage set groups, and resource record sets within them, x-VES-I/O-managed set is managed
by F5.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 50,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 50,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "50",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "50",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

## Direct properties

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/metadata/): complete subsection reference.

- [rr_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/): complete subsection reference.

## Next pages

- [primary.rr_set_group.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/metadata/)
- [primary.rr_set_group.rr_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/)
- [primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/)
- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/)
