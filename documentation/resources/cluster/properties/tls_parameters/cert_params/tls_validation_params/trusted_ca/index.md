---
page_title: "tls_parameters.cert_params.tls_validation_params.trusted_ca"
subcategory: ""
description: "tls_parameters.cert_params.tls_validation_params.trusted_ca for xcsh_cluster."
xcsh_docs: {"aliases": [], "body_bytes": 2144, "body_sha256": "sha256:814046876ebdf832cb94971a143e455c74207cbb010cdc4565f2f49dd2b75b07", "child_ids": ["xcsh-docs:resources:cluster:properties:tls_parameters:cert_params:tls_validation_params:trusted_ca:trusted_ca_list"], "collection_id": "xcsh-docs:resources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:cluster:properties:tls_parameters:cert_params:tls_validation_params:trusted_ca", "parent_id": "xcsh-docs:resources:cluster:properties:tls_parameters:cert_params:tls_validation_params", "path": "documentation/resources/cluster/properties/tls_parameters/cert_params/tls_validation_params/trusted_ca/index.md", "provider_name": "cluster", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["tls_parameters", "cert_params", "tls_validation_params", "trusted_ca"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cluster/properties/tls_parameters/cert_params/tls_validation_params/trusted_ca/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tls_parameters.cert_params.tls_validation_params.trusted_ca for xcsh_cluster.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_parameters.cert_params.tls_validation_params.trusted_ca

Breadcrumbs:

- [xcsh_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/)
- [tls_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/)
- [tls_parameters.cert_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/cert_params/)
- [tls_parameters.cert_params.tls_validation_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/cert_params/tls_validation_params/)
- tls_parameters.cert_params.tls_validation_params.trusted_ca

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
trusted_ca {
  # Configure direct properties listed below.
}
```

## Direct properties

- [trusted_ca_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/cert_params/tls_validation_params/trusted_ca/trusted_ca_list/): complete subsection reference.

## Next pages

- [tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/cert_params/tls_validation_params/trusted_ca/trusted_ca_list/)
- [tls_parameters.cert_params.tls_validation_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/cert_params/tls_validation_params/)
- [xcsh_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/)
