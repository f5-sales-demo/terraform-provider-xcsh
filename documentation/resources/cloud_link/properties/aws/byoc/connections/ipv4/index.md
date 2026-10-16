---
page_title: "aws.byoc.connections.ipv4"
subcategory: ""
description: "Configure BGP IPv4 peering for endpoints."
xcsh_docs: {"aliases": ["aws byoc connections ipv4"], "body_bytes": 3599, "body_sha256": "sha256:e42cc38f866b2f0a8d5141488992ed4e9e49c1cc158562a5d61f655224a91bcf", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cloud_link:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_link:properties:aws:byoc:connections:ipv4", "parent_id": "xcsh-docs:resources:cloud_link:properties:aws:byoc:connections", "path": "documentation/resources/cloud_link/properties/aws/byoc/connections/ipv4/index.md", "product": "distributed-cloud", "provider_name": "cloud_link", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "resources", "registry_anchor": "canonical-0302023103302130-0112101233231230-3203233032101133-1032133221323000-3100323202130322-3211310322230233-2233020131331030-3320031233333020", "registry_path": "docs/guides/resources--cloud_link--reference--group-001.md", "relationships": [{"anchor": "schema-aws--byoc--connections--ipv4--aws_router_peer_address", "enforcement": "provider-schema", "group": "aws.byoc.connections.ipv4:RequiredObjectAttributes:aws_router_peer_address,router_peer_address", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_link:properties:aws:byoc:connections:ipv4", "type": "requires"}, {"anchor": "schema-aws--byoc--connections--ipv4--router_peer_address", "enforcement": "provider-schema", "group": "aws.byoc.connections.ipv4:RequiredObjectAttributes:aws_router_peer_address,router_peer_address", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_link:properties:aws:byoc:connections:ipv4", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["aws", "byoc", "connections", "ipv4"], "schema_version": 1, "sections": [{"aliases": ["aws byoc connections ipv4 aws router peer address"], "anchor": "schema-aws--byoc--connections--ipv4--aws_router_peer_address", "description": "The BGP peer IP configured on the AWS endpoint.", "document_id": "xcsh-docs:resources:cloud_link:properties:aws:byoc:connections:ipv4", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws", "byoc", "connections", "ipv4", "aws_router_peer_address"], "syntax": "attribute", "type": "string"}, {"aliases": ["aws byoc connections ipv4 router peer address"], "anchor": "schema-aws--byoc--connections--ipv4--router_peer_address", "description": "The BGP peer IP configured on your (customer) endpoint.", "document_id": "xcsh-docs:resources:cloud_link:properties:aws:byoc:connections:ipv4", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws", "byoc", "connections", "ipv4", "router_peer_address"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_link/properties/aws/byoc/connections/ipv4/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "Configure BGP IPv4 peering for endpoints.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["cloud_linkCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws.byoc.connections.ipv4

Breadcrumbs:

- [xcsh_cloud_link](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/)
- [aws](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/aws/)
- [aws.byoc](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/aws/byoc/)
- [aws.byoc.connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/aws/byoc/connections/)
- aws.byoc.connections.ipv4

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configure BGP IPv4 peering for endpoints.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("aws_router_peer_address",
    "router_peer_address")}
```

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
ipv4 {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-aws--byoc--connections--ipv4--aws_router_peer_address"></a>

### aws_router_peer_address property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-15T10:49:30+00:00"
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

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-15T10:49:30+00:00"
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
