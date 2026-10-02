---
page_title: "ingress_egress_gw.az_nodes.inside_subnet"
subcategory: "Infrastructure"
description: "Parameters for AWS subnet."
xcsh_docs: {"aliases": ["ingress egress gw az nodes inside subnet"], "body_bytes": 2968, "body_sha256": "sha256:5861dfd6c0029e573cfd0d0bf6d800f06a4d30f996ff1ac9a2aab08924771050", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:aws_vpc_site:properties:ingress_egress_gw:az_nodes:inside_subnet:subnet_param"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_egress_gw:az_nodes:inside_subnet", "parent_id": "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_egress_gw:az_nodes", "path": "documentation/data-sources/aws_vpc_site/properties/ingress_egress_gw/az_nodes/inside_subnet/index.md", "product": "distributed-cloud", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0303122223100200-3001101120030203-2013330112331030-0323030001133312-2323011323130310-3310121032023322-0311122211221302-3322030330000232", "registry_path": "docs/guides/data-sources--aws_vpc_site--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_egress_gw", "az_nodes", "inside_subnet"], "schema_version": 1, "sections": [{"aliases": ["existing subnet id"], "anchor": "schema-ingress_egress_gw--az_nodes--inside_subnet--existing_subnet_id", "description": "Exclusive with Information about existing subnet ID.", "document_id": "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_egress_gw:az_nodes:inside_subnet", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "az_nodes", "inside_subnet", "existing_subnet_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["subnet param"], "anchor": "section", "description": "Parameters for creating a new cloud subnet.", "document_id": "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_egress_gw:az_nodes:inside_subnet:subnet_param", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_egress_gw", "az_nodes", "inside_subnet", "subnet_param"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_vpc_site/properties/ingress_egress_gw/az_nodes/inside_subnet/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Parameters for AWS subnet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.az_nodes.inside_subnet

Breadcrumbs:

- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/)
- [ingress_egress_gw.az_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/az_nodes/)
- ingress_egress_gw.az_nodes.inside_subnet

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for inside subnet.

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

<a id="schema-ingress_egress_gw--az_nodes--inside_subnet--existing_subnet_id"></a>

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

- [subnet_param](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/az_nodes/inside_subnet/subnet_param/): complete subsection reference.

## Next pages

- [ingress_egress_gw.az_nodes.inside_subnet.subnet_param](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/az_nodes/inside_subnet/subnet_param/)
- [ingress_egress_gw.az_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/az_nodes/)
- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/)
