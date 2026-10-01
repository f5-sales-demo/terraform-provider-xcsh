---
page_title: "where.virtual_network"
subcategory: ""
description: "where.virtual_network for xcsh_advertise_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1375, "body_sha256": "sha256:5c415b7e738a3627238ad4967f79a5d4529138e5b096b0418a3e810a3271fd96", "canonical_id": "xcsh-docs:resources:advertise_policy:properties:where:virtual_network", "child_ids": ["xcsh-docs:resources:advertise_policy:properties:where:virtual_network:ref"], "collection_id": "xcsh-docs:resources:advertise_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:advertise_policy:properties:where:virtual_network", "parent_id": "xcsh-docs:resources:advertise_policy:properties:where", "path": "docs/guides/resources--advertise_policy--properties--where--virtual_network.md", "provider_name": "advertise_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["where", "virtual_network"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/advertise_policy/properties/where/virtual_network/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "where.virtual_network for xcsh_advertise_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["advertise_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# where.virtual_network

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md)
- [Property reference](resources--advertise_policy--reference.md)
- [where](resources--advertise_policy--properties--where.md)
- where.virtual_network

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Specifies a direct reference to a network configuration object.

Upstream description:

This specifies a direct reference to a network configuration object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ref")}
```

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
virtual_network {
  # Configure direct properties listed below.
}
```

## Direct properties

- [ref](resources--advertise_policy--properties--where--virtual_network--ref.md): complete subsection reference.

## Next pages

- [where.virtual_network.ref](resources--advertise_policy--properties--where--virtual_network--ref.md)
- [where](resources--advertise_policy--properties--where.md)
- [xcsh_advertise_policy](../resources/advertise_policy.md)
