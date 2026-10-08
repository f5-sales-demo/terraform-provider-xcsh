---
page_title: "dns_ntp_config.custom_ntp"
subcategory: ""
description: "NTP Servers."
xcsh_docs: {"aliases": ["dns ntp config custom ntp"], "body_bytes": 1729, "body_sha256": "sha256:127d994de8cb53901f2891c31d097455f6f3deef91b2ccaaea95b31169e60b6e", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:dns_ntp_config:custom_ntp", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:dns_ntp_config", "path": "documentation/data-sources/securemesh_site_v2/properties/dns_ntp_config/custom_ntp/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-1311221211011010-0130031213233020-3003021213202022-1331131203313211-0132023222111032-1100011002003001-0300222032102230-0100223203113013", "registry_path": "docs/guides/data-sources--securemesh_site_v2--reference--group-007.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["dns_ntp_config", "custom_ntp"], "schema_version": 1, "sections": [{"aliases": ["dns ntp config custom ntp ntp servers"], "anchor": "schema-dns_ntp_config--custom_ntp--ntp_servers", "description": "NTP Servers.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:dns_ntp_config:custom_ntp", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dns_ntp_config", "custom_ntp", "ntp_servers"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/dns_ntp_config/custom_ntp/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "NTP Servers.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dns_ntp_config.custom_ntp

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
- [dns_ntp_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/dns_ntp_config/)
- dns_ntp_config.custom_ntp

<a id="section"></a>

Type: `"single"`. Computed.

NTP Servers. NTP Servers.

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

<a id="schema-dns_ntp_config--custom_ntp--ntp_servers"></a>

### ntp_servers property

Type: `["list", "string"]`. Computed.

NTP Servers. NTP Servers.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```
