---
page_title: "aws_parameters.new_tgw"
subcategory: ""
description: "aws_parameters.new_tgw for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": [], "body_bytes": 1381, "body_sha256": "sha256:b3a009d763bf3d6d719c784b48b6d1f9fd916b4fb717f7b358f3d45dfb7e2de5", "canonical_id": "xcsh-docs:data-sources:aws_tgw_site:properties:aws_parameters:new_tgw", "child_ids": ["xcsh-docs:data-sources:aws_tgw_site:properties:aws_parameters:new_tgw:system_generated", "xcsh-docs:data-sources:aws_tgw_site:properties:aws_parameters:new_tgw:user_assigned"], "collection_id": "xcsh-docs:data-sources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_tgw_site:properties:aws_parameters:new_tgw", "parent_id": "xcsh-docs:data-sources:aws_tgw_site:properties:aws_parameters", "path": "docs/guides/data-sources--aws_tgw_site--properties--aws_parameters--new_tgw.md", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["aws_parameters", "new_tgw"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_tgw_site/properties/aws_parameters/new_tgw/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "aws_parameters.new_tgw for xcsh_aws_tgw_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_parameters.new_tgw

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md)
- [Property reference](data-sources--aws_tgw_site--reference.md)
- [aws_parameters](data-sources--aws_tgw_site--properties--aws_parameters.md)
- aws_parameters.new_tgw

<a id="section"></a>

Type: `"single"`. Computed.

TGWParamsType.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-asn_choice": "[\"system_generated\",\"user_assigned\"]"
}
```

## Direct properties

- [system_generated](data-sources--aws_tgw_site--properties--aws_parameters--new_tgw--system_generated.md): complete subsection reference.

- [user_assigned](data-sources--aws_tgw_site--properties--aws_parameters--new_tgw--user_assigned.md): complete subsection reference.

## Next pages

- [aws_parameters.new_tgw.system_generated](data-sources--aws_tgw_site--properties--aws_parameters--new_tgw--system_generated.md)
- [aws_parameters.new_tgw.user_assigned](data-sources--aws_tgw_site--properties--aws_parameters--new_tgw--user_assigned.md)
- [aws_parameters](data-sources--aws_tgw_site--properties--aws_parameters.md)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md)
