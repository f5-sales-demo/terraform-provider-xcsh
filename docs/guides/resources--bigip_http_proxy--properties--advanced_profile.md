---
page_title: "advanced_profile"
subcategory: ""
description: "advanced_profile for xcsh_bigip_http_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 1715, "body_sha256": "sha256:1e77fc2b0fa201458b62b105f0ac678af7387fc96c1e7e173a20005a047888b1", "canonical_id": "xcsh-docs:resources:bigip_http_proxy:properties:advanced_profile", "child_ids": ["xcsh-docs:resources:bigip_http_proxy:properties:advanced_profile:disable_spec", "xcsh-docs:resources:bigip_http_proxy:properties:advanced_profile:enable_default_profile"], "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bigip_http_proxy:properties:advanced_profile", "parent_id": "xcsh-docs:resources:bigip_http_proxy:reference", "path": "docs/guides/resources--bigip_http_proxy--properties--advanced_profile.md", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["advanced_profile"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/properties/advanced_profile/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "advanced_profile for xcsh_bigip_http_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# advanced_profile

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md)
- [Property reference](resources--bigip_http_proxy--reference.md)
- advanced_profile

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Defines various advanced Profile OPTIONS for a Loadbalancer.

Upstream description:

This defines various advanced Profile OPTIONS for a Loadbalancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_spec",
    "enable_default_profile")}
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
  "x-ves-oneof-field-choice": "[\"disable\",\"enable_default_profile\"]"
}
```

Terraform syntax:

```terraform
advanced_profile {
  # Configure direct properties listed below.
}
```

## Direct properties

- [disable_spec](resources--bigip_http_proxy--properties--advanced_profile--disable_spec.md): complete subsection reference.

- [enable_default_profile](resources--bigip_http_proxy--properties--advanced_profile--enable_default_profile.md): complete subsection reference.

## Next pages

- [advanced_profile.disable_spec](resources--bigip_http_proxy--properties--advanced_profile--disable_spec.md)
- [advanced_profile.enable_default_profile](resources--bigip_http_proxy--properties--advanced_profile--enable_default_profile.md)
- [Property reference](resources--bigip_http_proxy--reference.md)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md)
