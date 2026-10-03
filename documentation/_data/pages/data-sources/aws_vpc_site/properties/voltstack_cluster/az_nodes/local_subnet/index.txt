---
page_title: "voltstack_cluster.az_nodes.local_subnet"
subcategory: "Infrastructure"
description: "Parameters for AWS subnet."
xcsh_docs: {"aliases": ["voltstack cluster az nodes local subnet"], "body_bytes": 2961, "body_sha256": "sha256:f48fce1ce03c80e9e752053935c458c57083d7452056790d2c47bd8418f5ea80", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:az_nodes:local_subnet:subnet_param"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:az_nodes:local_subnet", "parent_id": "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:az_nodes", "path": "documentation/data-sources/aws_vpc_site/properties/voltstack_cluster/az_nodes/local_subnet/index.md", "product": "distributed-cloud", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3003203303132311-2313110332111333-1232002233212232-0030202302031230-1133312200303322-2223320010032132-1132130121300001-0033302031213002", "registry_path": "docs/guides/data-sources--aws_vpc_site--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["voltstack_cluster", "az_nodes", "local_subnet"], "schema_version": 1, "sections": [{"aliases": ["voltstack cluster az nodes local subnet existing subnet id"], "anchor": "schema-voltstack_cluster--az_nodes--local_subnet--existing_subnet_id", "description": "Exclusive with Information about existing subnet ID.", "document_id": "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:az_nodes:local_subnet", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["voltstack_cluster", "az_nodes", "local_subnet", "existing_subnet_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["voltstack cluster az nodes local subnet subnet param"], "anchor": "section", "description": "Parameters for creating a new cloud subnet.", "document_id": "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:az_nodes:local_subnet:subnet_param", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["voltstack_cluster", "az_nodes", "local_subnet", "subnet_param"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_vpc_site/properties/voltstack_cluster/az_nodes/local_subnet/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Parameters for AWS subnet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# voltstack_cluster.az_nodes.local_subnet

Breadcrumbs:

- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/)
- [voltstack_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/)
- [voltstack_cluster.az_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/az_nodes/)
- voltstack_cluster.az_nodes.local_subnet

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for local subnet.

Upstream description:

Parameters for AWS subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"existing_subnet_id\",\"subnet_param\"]"
}
```

## Direct properties

<a id="schema-voltstack_cluster--az_nodes--local_subnet--existing_subnet_id"></a>

### existing_subnet_id property

Type: `"string"`. Computed.

Exclusive with \[subnet\_param\] Information about existing subnet ID.

Upstream description:

Exclusive with \[subnet\_param\] Information about existing subnet ID.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

- [subnet_param](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/az_nodes/local_subnet/subnet_param/): complete subsection reference.

## Next pages

- [voltstack_cluster.az_nodes.local_subnet.subnet_param](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/az_nodes/local_subnet/subnet_param/)
- [voltstack_cluster.az_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/az_nodes/)
- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/)
