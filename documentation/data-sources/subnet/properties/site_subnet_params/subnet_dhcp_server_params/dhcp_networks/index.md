---
page_title: "site_subnet_params.subnet_dhcp_server_params.dhcp_networks"
subcategory: ""
description: "List of networks from which DHCP server can allocate IP addresses."
xcsh_docs: {"aliases": ["site subnet params subnet dhcp server params dhcp networks"], "body_bytes": 2436, "body_sha256": "sha256:a5268a8b68613f948d2450010f08f4a13e20bee414c302c3045ec88f88a4c5a5", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:subnet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:subnet:properties:site_subnet_params:subnet_dhcp_server_params:dhcp_networks", "parent_id": "xcsh-docs:data-sources:subnet:properties:site_subnet_params:subnet_dhcp_server_params", "path": "documentation/data-sources/subnet/properties/site_subnet_params/subnet_dhcp_server_params/dhcp_networks/index.md", "product": "distributed-cloud", "provider_name": "subnet", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-2031201330311312-2001113032012130-3001233233212132-0133221111022111-0102222310323013-1103012220321303-0103012321103020-1323030333030211", "registry_path": "docs/guides/data-sources--subnet--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["site_subnet_params", "subnet_dhcp_server_params", "dhcp_networks"], "schema_version": 1, "sections": [{"aliases": ["site subnet params subnet dhcp server params dhcp networks network prefix"], "anchor": "schema-site_subnet_params--subnet_dhcp_server_params--dhcp_networks--network_prefix", "description": "Exclusive with Network prefix for subnet.", "document_id": "xcsh-docs:data-sources:subnet:properties:site_subnet_params:subnet_dhcp_server_params:dhcp_networks", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["site_subnet_params", "subnet_dhcp_server_params", "dhcp_networks", "network_prefix"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/subnet/properties/site_subnet_params/subnet_dhcp_server_params/dhcp_networks/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "List of networks from which DHCP server can allocate IP addresses.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["subnetCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# site_subnet_params.subnet_dhcp_server_params.dhcp_networks

Breadcrumbs:

- [xcsh_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/subnet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/subnet/properties/)
- [site_subnet_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/subnet/properties/site_subnet_params/)
- [site_subnet_params.subnet_dhcp_server_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/subnet/properties/site_subnet_params/subnet_dhcp_server_params/)
- site_subnet_params.subnet_dhcp_server_params.dhcp_networks

<a id="section"></a>

Type: `"list"`. Computed.

List of networks from which DHCP server can allocate IP addresses.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

<a id="schema-site_subnet_params--subnet_dhcp_server_params--dhcp_networks--network_prefix"></a>

### network_prefix property

Type: `"string"`. Computed.

Exclusive with \[\] Network prefix for subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```
