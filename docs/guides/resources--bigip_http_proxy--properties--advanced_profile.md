---
page_title: "advanced_profile"
subcategory: ""
description: "advanced_profile for xcsh_bigip_http_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 1616, "body_sha256": "sha256:c98acec7cdf5e83191ab199b0d1df12290ad4986e250b18db474532c586b9f83", "canonical_id": "xcsh-docs:resources:bigip_http_proxy:properties:advanced_profile", "child_ids": ["xcsh-docs:resources:bigip_http_proxy:properties:advanced_profile:disable_spec", "xcsh-docs:resources:bigip_http_proxy:properties:advanced_profile:enable_default_profile"], "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bigip_http_proxy:properties:advanced_profile", "parent_id": "xcsh-docs:resources:bigip_http_proxy:reference", "path": "docs/guides/resources--bigip_http_proxy--properties--advanced_profile.md", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["advanced_profile"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/properties/advanced_profile/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "advanced_profile for xcsh_bigip_http_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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
