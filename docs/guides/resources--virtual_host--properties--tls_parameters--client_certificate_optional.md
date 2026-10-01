---
page_title: "tls_parameters.client_certificate_optional"
subcategory: ""
description: "tls_parameters.client_certificate_optional for xcsh_virtual_host."
xcsh_docs: {"aliases": [], "body_bytes": 1034, "body_sha256": "sha256:92fd1681fc991d7fc8fd11f32017d7b2f7049c991ca25f61586144244d9a64ef", "canonical_id": "xcsh-docs:resources:virtual_host:properties:tls_parameters:client_certificate_optional", "child_ids": [], "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:tls_parameters:client_certificate_optional", "parent_id": "xcsh-docs:resources:virtual_host:properties:tls_parameters", "path": "docs/guides/resources--virtual_host--properties--tls_parameters--client_certificate_optional.md", "provider_name": "virtual_host", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tls_parameters", "client_certificate_optional"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/tls_parameters/client_certificate_optional/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tls_parameters.client_certificate_optional for xcsh_virtual_host.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_parameters.client_certificate_optional

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md)
- [Property reference](resources--virtual_host--reference.md)
- [tls_parameters](resources--virtual_host--properties--tls_parameters.md)
- tls_parameters.client_certificate_optional

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
client_certificate_optional = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [tls_parameters](resources--virtual_host--properties--tls_parameters.md)
- [xcsh_virtual_host](../resources/virtual_host.md)
