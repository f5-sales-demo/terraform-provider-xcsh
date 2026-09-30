---
page_title: "authentication.cookie_params.kms_key_hmac"
subcategory: ""
description: "authentication.cookie_params.kms_key_hmac for xcsh_virtual_host."
xcsh_docs: {"aliases": [], "body_bytes": 1045, "body_sha256": "sha256:0f7fef885558d6f5106269593e7800e60780aa397953f9daa31c818e0bbaf54c", "canonical_id": "xcsh-docs:resources:virtual_host:properties:authentication:cookie_params:kms_key_hmac", "child_ids": [], "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:authentication:cookie_params:kms_key_hmac", "parent_id": "xcsh-docs:resources:virtual_host:properties:authentication:cookie_params", "path": "docs/guides/resources--virtual_host--properties--authentication--cookie_params--kms_key_hmac.md", "provider_name": "virtual_host", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["authentication", "cookie_params", "kms_key_hmac"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/authentication/cookie_params/kms_key_hmac/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "authentication.cookie_params.kms_key_hmac for xcsh_virtual_host.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# authentication.cookie_params.kms_key_hmac

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md)
- [Property reference](resources--virtual_host--reference.md)
- [authentication](resources--virtual_host--properties--authentication.md)
- [authentication.cookie_params](resources--virtual_host--properties--authentication--cookie_params.md)
- authentication.cookie_params.kms_key_hmac

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for kms key hmac.

Upstream description:

Reference to KMS Key Object.

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
kms_key_hmac = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [authentication.cookie_params](resources--virtual_host--properties--authentication--cookie_params.md)
- [xcsh_virtual_host](../resources/virtual_host.md)
