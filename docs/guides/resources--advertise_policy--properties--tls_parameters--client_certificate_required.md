---
page_title: "tls_parameters.client_certificate_required"
subcategory: ""
description: "tls_parameters.client_certificate_required for xcsh_advertise_policy."
xcsh_docs: {"aliases": [], "body_bytes": 963, "body_sha256": "sha256:a08f0e4bccc508f8d881b9718e6009a7e9a7128ee703fe18db69cc5b53c74da0", "canonical_id": "xcsh-docs:resources:advertise_policy:properties:tls_parameters:client_certificate_required", "child_ids": [], "collection_id": "xcsh-docs:resources:advertise_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:advertise_policy:properties:tls_parameters:client_certificate_required", "parent_id": "xcsh-docs:resources:advertise_policy:properties:tls_parameters", "path": "docs/guides/resources--advertise_policy--properties--tls_parameters--client_certificate_required.md", "provider_name": "advertise_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tls_parameters", "client_certificate_required"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/advertise_policy/properties/tls_parameters/client_certificate_required/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tls_parameters.client_certificate_required for xcsh_advertise_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["advertise_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# tls_parameters.client_certificate_required

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md)
- [Property reference](resources--advertise_policy--reference.md)
- [tls_parameters](resources--advertise_policy--properties--tls_parameters.md)
- tls_parameters.client_certificate_required

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
client_certificate_required = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [tls_parameters](resources--advertise_policy--properties--tls_parameters.md)
- [xcsh_advertise_policy](../resources/advertise_policy.md)
