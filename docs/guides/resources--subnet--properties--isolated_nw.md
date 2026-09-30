---
page_title: "isolated_nw"
subcategory: ""
description: "isolated_nw for xcsh_subnet."
xcsh_docs: {"aliases": [], "body_bytes": 755, "body_sha256": "sha256:19a53708b27db892428b64995f88c7a8e64b1f49461369c93479020e1dc963ab", "canonical_id": "xcsh-docs:resources:subnet:properties:isolated_nw", "child_ids": [], "collection_id": "xcsh-docs:resources:subnet:collection", "completeness": "complete", "id": "xcsh-docs:resources:subnet:properties:isolated_nw", "parent_id": "xcsh-docs:resources:subnet:reference", "path": "docs/guides/resources--subnet--properties--isolated_nw.md", "provider_name": "subnet", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["isolated_nw"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/subnet/properties/isolated_nw/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "isolated_nw for xcsh_subnet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["subnetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# isolated_nw

Breadcrumbs:

- [xcsh_subnet](../resources/subnet.md)
- [Property reference](resources--subnet--reference.md)
- isolated_nw

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for isolated nw.

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
isolated_nw = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--subnet--reference.md)
- [xcsh_subnet](../resources/subnet.md)
