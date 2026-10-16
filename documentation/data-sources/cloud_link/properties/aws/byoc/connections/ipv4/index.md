---
page_title: "aws.byoc.connections.ipv4"
subcategory: ""
description: "Configure BGP IPv4 peering for endpoints."
xcsh_docs: {"aliases": ["aws byoc connections ipv4"], "body_bytes": 3282, "body_sha256": "sha256:007a2f3c989f4e113c487dc4688b4ed3b883cb5fc0566ce63944b795a0263c37", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cloud_link:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_link:properties:aws:byoc:connections:ipv4", "parent_id": "xcsh-docs:data-sources:cloud_link:properties:aws:byoc:connections", "path": "documentation/data-sources/cloud_link/properties/aws/byoc/connections/ipv4/index.md", "product": "distributed-cloud", "provider_name": "cloud_link", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "data-sources", "registry_anchor": "canonical-1223231323121200-0331331200303102-3033122032332332-2330012220000111-2313000001020112-1202033210133001-3132303113020302-1300321221311220", "registry_path": "docs/guides/data-sources--cloud_link--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["aws", "byoc", "connections", "ipv4"], "schema_version": 1, "sections": [{"aliases": ["aws byoc connections ipv4 aws router peer address"], "anchor": "schema-aws--byoc--connections--ipv4--aws_router_peer_address", "description": "The BGP peer IP configured on the AWS endpoint.", "document_id": "xcsh-docs:data-sources:cloud_link:properties:aws:byoc:connections:ipv4", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws", "byoc", "connections", "ipv4", "aws_router_peer_address"], "syntax": "attribute", "type": "string"}, {"aliases": ["aws byoc connections ipv4 router peer address"], "anchor": "schema-aws--byoc--connections--ipv4--router_peer_address", "description": "The BGP peer IP configured on your (customer) endpoint.", "document_id": "xcsh-docs:data-sources:cloud_link:properties:aws:byoc:connections:ipv4", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws", "byoc", "connections", "ipv4", "router_peer_address"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_link/properties/aws/byoc/connections/ipv4/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "Configure BGP IPv4 peering for endpoints.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["cloud_linkCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws.byoc.connections.ipv4

Breadcrumbs:

- [xcsh_cloud_link](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/properties/)
- [aws](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/properties/aws/)
- [aws.byoc](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/properties/aws/byoc/)
- [aws.byoc.connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/properties/aws/byoc/connections/)
- aws.byoc.connections.ipv4

<a id="section"></a>

Type: `"single"`. Computed.

Configure BGP IPv4 peering for endpoints.

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

<a id="schema-aws--byoc--connections--ipv4--aws_router_peer_address"></a>

### aws_router_peer_address property

Type: `"string"`. Computed.

The BGP peer IP configured on the AWS endpoint.

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
      "validatedAt": "2026-10-16T10:31:02+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "32",
    "ves.io.schema.rules.string.min_ip_prefix_length": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "32",
    "ves.io.schema.rules.string.min_ip_prefix_length": "1"
  }
}
```

<a id="schema-aws--byoc--connections--ipv4--router_peer_address"></a>

### router_peer_address property

Type: `"string"`. Computed.

The BGP peer IP configured on your (customer) endpoint.

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
      "validatedAt": "2026-10-16T10:31:02+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "32",
    "ves.io.schema.rules.string.min_ip_prefix_length": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "32",
    "ves.io.schema.rules.string.min_ip_prefix_length": "1"
  }
}
```
