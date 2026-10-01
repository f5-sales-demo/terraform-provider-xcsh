---
page_title: "passport.default_sw_version"
subcategory: ""
description: "passport.default_sw_version for xcsh_registration."
xcsh_docs: {"aliases": [], "body_bytes": 971, "body_sha256": "sha256:083177a113502f9088f3728a15ab7730252a7ad0b78911d57631c36d4cfac502", "canonical_id": "xcsh-docs:resources:registration:properties:passport:default_sw_version", "child_ids": [], "collection_id": "xcsh-docs:resources:registration:collection", "completeness": "complete", "id": "xcsh-docs:resources:registration:properties:passport:default_sw_version", "parent_id": "xcsh-docs:resources:registration:properties:passport", "path": "docs/guides/resources--registration--properties--passport--default_sw_version.md", "provider_name": "registration", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["passport", "default_sw_version"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/registration/properties/passport/default_sw_version/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "passport.default_sw_version for xcsh_registration.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["registrationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# passport.default_sw_version

Breadcrumbs:

- [xcsh_registration](../resources/registration.md)
- [Property reference](resources--registration--reference.md)
- [passport](resources--registration--properties--passport.md)
- passport.default_sw_version

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
default_sw_version = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [passport](resources--registration--properties--passport.md)
- [xcsh_registration](../resources/registration.md)
