---
page_title: "tls_parameters.common_params.validation_params.trusted_ca"
subcategory: ""
description: "tls_parameters.common_params.validation_params.trusted_ca for xcsh_advertise_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1689, "body_sha256": "sha256:014d7c1a9ef5c47df262506e16dbeedb4a6368956ad4dde5368d84f949aa3416", "canonical_id": "xcsh-docs:data-sources:advertise_policy:properties:tls_parameters:common_params:validation_params:trusted_ca", "child_ids": ["xcsh-docs:data-sources:advertise_policy:properties:tls_parameters:common_params:validation_params:trusted_ca:trusted_ca_list"], "collection_id": "xcsh-docs:data-sources:advertise_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:advertise_policy:properties:tls_parameters:common_params:validation_params:trusted_ca", "parent_id": "xcsh-docs:data-sources:advertise_policy:properties:tls_parameters:common_params:validation_params", "path": "docs/guides/data-sources--advertise_policy--properties--tls_parameters--common_params--validation_params--trusted_ca.md", "provider_name": "advertise_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tls_parameters", "common_params", "validation_params", "trusted_ca"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/advertise_policy/properties/tls_parameters/common_params/validation_params/trusted_ca/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tls_parameters.common_params.validation_params.trusted_ca for xcsh_advertise_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["advertise_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_parameters.common_params.validation_params.trusted_ca

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md)
- [Property reference](data-sources--advertise_policy--reference.md)
- [tls_parameters](data-sources--advertise_policy--properties--tls_parameters.md)
- [tls_parameters.common_params](data-sources--advertise_policy--properties--tls_parameters--common_params.md)
- [tls_parameters.common_params.validation_params](data-sources--advertise_policy--properties--tls_parameters--common_params--validation_params.md)
- tls_parameters.common_params.validation_params.trusted_ca

<a id="section"></a>

Type: `"single"`. Computed.

Root CA Certificate Reference. Reference to Root CA Certificate.

Upstream description:

Reference to Root CA Certificate.

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

## Direct properties

- [trusted_ca_list](data-sources--advertise_policy--properties--tls_parameters--common_params--validation_params--trusted_ca--trusted_ca_list.md): complete subsection reference.

## Next pages

- [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list](data-sources--advertise_policy--properties--tls_parameters--common_params--validation_params--trusted_ca--trusted_ca_list.md)
- [tls_parameters.common_params.validation_params](data-sources--advertise_policy--properties--tls_parameters--common_params--validation_params.md)
- [xcsh_advertise_policy](../data-sources/advertise_policy.md)
