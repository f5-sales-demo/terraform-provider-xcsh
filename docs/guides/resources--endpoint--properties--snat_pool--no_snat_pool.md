---
page_title: "snat_pool.no_snat_pool"
subcategory: "Networking"
description: "snat_pool.no_snat_pool for xcsh_endpoint."
xcsh_docs: {"aliases": [], "body_bytes": 855, "body_sha256": "sha256:63680c65d9ebbbe4205183815914309db29c8cff9d70857eb3410fb9f6219dfe", "canonical_id": "xcsh-docs:resources:endpoint:properties:snat_pool:no_snat_pool", "child_ids": [], "collection_id": "xcsh-docs:resources:endpoint:collection", "completeness": "complete", "id": "xcsh-docs:resources:endpoint:properties:snat_pool:no_snat_pool", "parent_id": "xcsh-docs:resources:endpoint:properties:snat_pool", "path": "docs/guides/resources--endpoint--properties--snat_pool--no_snat_pool.md", "provider_name": "endpoint", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["snat_pool", "no_snat_pool"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/endpoint/properties/snat_pool/no_snat_pool/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "snat_pool.no_snat_pool for xcsh_endpoint.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["endpointCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# snat_pool.no_snat_pool

Breadcrumbs:

- [xcsh_endpoint](../resources/endpoint.md)
- [Property reference](resources--endpoint--reference.md)
- [snat_pool](resources--endpoint--properties--snat_pool.md)
- snat_pool.no_snat_pool

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no snat pool.

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
no_snat_pool = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [snat_pool](resources--endpoint--properties--snat_pool.md)
- [xcsh_endpoint](../resources/endpoint.md)
