---
page_title: "snat_pool"
subcategory: "Networking"
description: "snat_pool for xcsh_endpoint."
xcsh_docs: {"aliases": [], "body_bytes": 1443, "body_sha256": "sha256:7af7038aff3c437cdefbd52be92acd738870bcf713d11a6dc4c98fea3f6ebcfa", "canonical_id": "xcsh-docs:resources:endpoint:properties:snat_pool", "child_ids": ["xcsh-docs:resources:endpoint:properties:snat_pool:no_snat_pool", "xcsh-docs:resources:endpoint:properties:snat_pool:snat_pool"], "collection_id": "xcsh-docs:resources:endpoint:collection", "completeness": "complete", "id": "xcsh-docs:resources:endpoint:properties:snat_pool", "parent_id": "xcsh-docs:resources:endpoint:reference", "path": "docs/guides/resources--endpoint--properties--snat_pool.md", "provider_name": "endpoint", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["snat_pool"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/endpoint/properties/snat_pool/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "snat_pool for xcsh_endpoint.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["endpointCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# snat_pool

Breadcrumbs:

- [xcsh_endpoint](../resources/endpoint.md)
- [Property reference](resources--endpoint--reference.md)
- snat_pool

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

SNAT Pool. SNAT Pool configuration.

Upstream description:

SNAT Pool configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_snat_pool",
    "snat_pool")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-snat_pool_choice": "[\"no_snat_pool\",\"snat_pool\"]"
}
```

Terraform syntax:

```terraform
snat_pool {
  # Configure direct properties listed below.
}
```

## Direct properties

- [no_snat_pool](resources--endpoint--properties--snat_pool--no_snat_pool.md): complete subsection reference.

- [snat_pool](resources--endpoint--properties--snat_pool--snat_pool.md): complete subsection reference.

## Next pages

- [snat_pool.no_snat_pool](resources--endpoint--properties--snat_pool--no_snat_pool.md)
- [snat_pool.snat_pool](resources--endpoint--properties--snat_pool--snat_pool.md)
- [Property reference](resources--endpoint--reference.md)
- [xcsh_endpoint](../resources/endpoint.md)
