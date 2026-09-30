---
page_title: "f5_orchestrated_routing"
subcategory: "Infrastructure"
description: "f5_orchestrated_routing for xcsh_aws_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 1149, "body_sha256": "sha256:90ff4ed355e6aab46f40af174e8109c3dcbfe292769f8429506dbaf0adc52e03", "canonical_id": "xcsh-docs:resources:aws_vpc_site:properties:f5_orchestrated_routing", "child_ids": [], "collection_id": "xcsh-docs:resources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_vpc_site:properties:f5_orchestrated_routing", "parent_id": "xcsh-docs:resources:aws_vpc_site:reference", "path": "docs/guides/resources--aws_vpc_site--properties--f5_orchestrated_routing.md", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["f5_orchestrated_routing"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_vpc_site/properties/f5_orchestrated_routing/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "f5_orchestrated_routing for xcsh_aws_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# f5_orchestrated_routing

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md)
- [Property reference](resources--aws_vpc_site--reference.md)
- f5_orchestrated_routing

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: f5\_orchestrated\_routing, manual\_routing\] Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

OneOf alternatives in this subsection:

- [f5_orchestrated_routing](resources--aws_vpc_site--properties--f5_orchestrated_routing.md#section)
- [manual_routing](resources--aws_vpc_site--properties--manual_routing.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
f5_orchestrated_routing = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--aws_vpc_site--reference.md)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md)
