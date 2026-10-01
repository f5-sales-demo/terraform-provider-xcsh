---
page_title: "ingress_egress_gw.az_nodes.workload_subnet"
subcategory: "Infrastructure"
description: "ingress_egress_gw.az_nodes.workload_subnet for xcsh_aws_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 2580, "body_sha256": "sha256:e4b8eb3564aa6e66ca2813bf28b93a01ce2be6b509b8b297e31493103fd99d59", "canonical_id": "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_egress_gw:az_nodes:workload_subnet", "child_ids": ["xcsh-docs:data-sources:aws_vpc_site:properties:ingress_egress_gw:az_nodes:workload_subnet:subnet_param"], "collection_id": "xcsh-docs:data-sources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_egress_gw:az_nodes:workload_subnet", "parent_id": "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_egress_gw:az_nodes", "path": "docs/guides/data-sources--aws_vpc_site--properties--ingress_egress_gw--az_nodes--workload_subnet.md", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_egress_gw", "az_nodes", "workload_subnet"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_vpc_site/properties/ingress_egress_gw/az_nodes/workload_subnet/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw.az_nodes.workload_subnet for xcsh_aws_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.az_nodes.workload_subnet

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md)
- [Property reference](data-sources--aws_vpc_site--reference.md)
- [ingress_egress_gw](data-sources--aws_vpc_site--properties--ingress_egress_gw.md)
- [ingress_egress_gw.az_nodes](data-sources--aws_vpc_site--properties--ingress_egress_gw--az_nodes.md)
- ingress_egress_gw.az_nodes.workload_subnet

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for workload subnet.

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

<a id="schema-ingress_egress_gw--az_nodes--workload_subnet--existing_subnet_id"></a>

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

- [subnet_param](data-sources--aws_vpc_site--properties--ingress_egress_gw--az_nodes--workload_subnet--subnet_param.md): complete subsection reference.

## Next pages

- [ingress_egress_gw.az_nodes.workload_subnet.subnet_param](data-sources--aws_vpc_site--properties--ingress_egress_gw--az_nodes--workload_subnet--subnet_param.md)
- [ingress_egress_gw.az_nodes](data-sources--aws_vpc_site--properties--ingress_egress_gw--az_nodes.md)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md)
