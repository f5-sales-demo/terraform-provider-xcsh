---
page_title: "ingress_egress_gw.inside_subnet.existing_subnet"
subcategory: "Infrastructure"
description: "ingress_egress_gw.inside_subnet.existing_subnet for xcsh_gcp_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 2072, "body_sha256": "sha256:1038f4c9d9ac8c163a37fdf6c24e3d6091cb6c4694d0ad20f6f747b3c81b1d76", "canonical_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:inside_subnet:existing_subnet", "child_ids": [], "collection_id": "xcsh-docs:data-sources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:inside_subnet:existing_subnet", "parent_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:inside_subnet", "path": "docs/guides/data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_subnet--existing_subnet.md", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_egress_gw", "inside_subnet", "existing_subnet"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/gcp_vpc_site/properties/ingress_egress_gw/inside_subnet/existing_subnet/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw.inside_subnet.existing_subnet for xcsh_gcp_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# ingress_egress_gw.inside_subnet.existing_subnet

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md)
- [Property reference](data-sources--gcp_vpc_site--reference.md)
- [ingress_egress_gw](data-sources--gcp_vpc_site--properties--ingress_egress_gw.md)
- [ingress_egress_gw.inside_subnet](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_subnet.md)
- ingress_egress_gw.inside_subnet.existing_subnet

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for existing subnet.

Upstream description:

Name of existing GCP subnet.

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

<a id="schema-ingress_egress_gw--inside_subnet--existing_subnet--subnet_name"></a>

### subnet_name property

Type: `"string"`. Computed.

VPC Subnet Name. Name of your subnet in VPC network.

Upstream description:

Name of your subnet in VPC network.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
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
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

## Next pages

- [ingress_egress_gw.inside_subnet](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_subnet.md)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md)
