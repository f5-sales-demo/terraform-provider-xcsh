---
page_title: "site_subnet_params.subnet_dhcp_server_params.dhcp_networks"
subcategory: ""
description: "List of networks from which DHCP server can allocate IP addresses."
xcsh_docs: {"aliases": ["site subnet params subnet dhcp server params dhcp networks"], "body_bytes": 2546, "body_sha256": "sha256:9f70ec923b54c6d5f4dc09868cae29603cd8424a39124b3b5ffa09dde84d4ea4", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:subnet:collection", "completeness": "complete", "id": "xcsh-docs:resources:subnet:properties:site_subnet_params:subnet_dhcp_server_params:dhcp_networks", "parent_id": "xcsh-docs:resources:subnet:properties:site_subnet_params:subnet_dhcp_server_params", "path": "documentation/resources/subnet/properties/site_subnet_params/subnet_dhcp_server_params/dhcp_networks/index.md", "product": "distributed-cloud", "provider_name": "subnet", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-1212322310001120-0113213123323222-3120021021013102-1202231121203200-0221330101002121-3330223233322031-0210133223212210-3100033011133231", "registry_path": "docs/guides/resources--subnet--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["site_subnet_params", "subnet_dhcp_server_params", "dhcp_networks"], "schema_version": 1, "sections": [{"aliases": ["site subnet params subnet dhcp server params dhcp networks network prefix"], "anchor": "schema-site_subnet_params--subnet_dhcp_server_params--dhcp_networks--network_prefix", "description": "Exclusive with Network prefix for subnet.", "document_id": "xcsh-docs:resources:subnet:properties:site_subnet_params:subnet_dhcp_server_params:dhcp_networks", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["site_subnet_params", "subnet_dhcp_server_params", "dhcp_networks", "network_prefix"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/subnet/properties/site_subnet_params/subnet_dhcp_server_params/dhcp_networks/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "List of networks from which DHCP server can allocate IP addresses.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["subnetCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# site_subnet_params.subnet_dhcp_server_params.dhcp_networks

Breadcrumbs:

- [xcsh_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/subnet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/subnet/properties/)
- [site_subnet_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/subnet/properties/site_subnet_params/)
- [site_subnet_params.subnet_dhcp_server_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/subnet/properties/site_subnet_params/subnet_dhcp_server_params/)
- site_subnet_params.subnet_dhcp_server_params.dhcp_networks

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-20T00:10:01+00:00"
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

Terraform syntax:

```terraform
dhcp_networks {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-site_subnet_params--subnet_dhcp_server_params--dhcp_networks--network_prefix"></a>

### network_prefix property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-20T00:10:01+00:00"
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
