---
page_title: "ingress_gw"
subcategory: "Infrastructure"
description: "ingress_gw for xcsh_aws_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 2577, "body_sha256": "sha256:331165bd46a4fcdf01c79d7abae2536cb70ac3263de93eb63f5abc02d55b775d", "canonical_id": "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_gw", "child_ids": ["xcsh-docs:data-sources:aws_vpc_site:properties:ingress_gw:allowed_vip_port", "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_gw:az_nodes", "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_gw:performance_enhancement_mode"], "collection_id": "xcsh-docs:data-sources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_gw", "parent_id": "xcsh-docs:data-sources:aws_vpc_site:reference", "path": "docs/guides/data-sources--aws_vpc_site--properties--ingress_gw.md", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_gw"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_vpc_site/properties/ingress_gw/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_gw for xcsh_aws_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# ingress_gw

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md)
- [Property reference](data-sources--aws_vpc_site--reference.md)
- ingress_gw

<a id="section"></a>

Type: `"single"`. Computed.

AWS Ingress Gateway. Single interface AWS ingress site.

Upstream description:

Single interface AWS ingress site.

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

- [allowed_vip_port](data-sources--aws_vpc_site--properties--ingress_gw--allowed_vip_port.md): complete subsection reference.

<a id="schema-ingress_gw--aws_certified_hw"></a>

### aws_certified_hw property

Type: `"string"`. Computed.

\[Enum: aws-byol-voltmesh\] AWS Certified Hardware. Name for AWS certified hardware. The only
possible value is \`aws-byol-voltmesh\`.

Upstream description:

Name for AWS certified hardware.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "aws-byol-voltmesh"
  ],
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
    "ves.io.schema.rules.string.in": "[\\\"aws-byol-voltmesh\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"aws-byol-voltmesh\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [az_nodes](data-sources--aws_vpc_site--properties--ingress_gw--az_nodes.md): complete subsection reference.

- [performance_enhancement_mode](data-sources--aws_vpc_site--properties--ingress_gw--performance_enhancement_mode.md): complete subsection reference.

## Next pages

- [ingress_gw.allowed_vip_port](data-sources--aws_vpc_site--properties--ingress_gw--allowed_vip_port.md)
- [ingress_gw.az_nodes](data-sources--aws_vpc_site--properties--ingress_gw--az_nodes.md)
- [ingress_gw.performance_enhancement_mode](data-sources--aws_vpc_site--properties--ingress_gw--performance_enhancement_mode.md)
- [Property reference](data-sources--aws_vpc_site--reference.md)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md)
