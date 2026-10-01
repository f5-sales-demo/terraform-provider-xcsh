---
page_title: "connect_to_slo"
subcategory: ""
description: "connect_to_slo for xcsh_subnet."
xcsh_docs: {"aliases": [], "body_bytes": 866, "body_sha256": "sha256:9ee9f3cf0e132a45a5f23ad53a4c54802ff1d4b85bdef56e83b14b2dd7ecb4b3", "canonical_id": "xcsh-docs:resources:subnet:properties:connect_to_slo", "child_ids": [], "collection_id": "xcsh-docs:resources:subnet:collection", "completeness": "complete", "id": "xcsh-docs:resources:subnet:properties:connect_to_slo", "parent_id": "xcsh-docs:resources:subnet:reference", "path": "docs/guides/resources--subnet--properties--connect_to_slo.md", "provider_name": "subnet", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["connect_to_slo"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/subnet/properties/connect_to_slo/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "connect_to_slo for xcsh_subnet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["subnetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# connect_to_slo

Breadcrumbs:

- [xcsh_subnet](../resources/subnet.md)
- [Property reference](resources--subnet--reference.md)
- connect_to_slo

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for connect to slo.

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

Terraform syntax:

```terraform
connect_to_slo = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--subnet--reference.md)
- [xcsh_subnet](../resources/subnet.md)
