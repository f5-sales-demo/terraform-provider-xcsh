---
page_title: "tls_parameters.common_params.validation_params.trusted_ca"
subcategory: ""
description: "tls_parameters.common_params.validation_params.trusted_ca for xcsh_cluster."
xcsh_docs: {"aliases": [], "body_bytes": 1583, "body_sha256": "sha256:4889b077b321ac7275c21a30b5e04385093cdb68a4c2e7db86a63c82e85990aa", "canonical_id": "xcsh-docs:resources:cluster:properties:tls_parameters:common_params:validation_params:trusted_ca", "child_ids": ["xcsh-docs:resources:cluster:properties:tls_parameters:common_params:validation_params:trusted_ca:trusted_ca_list"], "collection_id": "xcsh-docs:resources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:cluster:properties:tls_parameters:common_params:validation_params:trusted_ca", "parent_id": "xcsh-docs:resources:cluster:properties:tls_parameters:common_params:validation_params", "path": "docs/guides/resources--cluster--properties--tls_parameters--common_params--validation_params--trusted_ca.md", "provider_name": "cluster", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tls_parameters", "common_params", "validation_params", "trusted_ca"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cluster/properties/tls_parameters/common_params/validation_params/trusted_ca/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tls_parameters.common_params.validation_params.trusted_ca for xcsh_cluster.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# tls_parameters.common_params.validation_params.trusted_ca

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md)
- [Property reference](resources--cluster--reference.md)
- [tls_parameters](resources--cluster--properties--tls_parameters.md)
- [tls_parameters.common_params](resources--cluster--properties--tls_parameters--common_params.md)
- [tls_parameters.common_params.validation_params](resources--cluster--properties--tls_parameters--common_params--validation_params.md)
- tls_parameters.common_params.validation_params.trusted_ca

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

- [trusted_ca_list](resources--cluster--properties--tls_parameters--common_params--validation_params--trusted_ca--trusted_ca_list.md): complete subsection reference.

## Next pages

- [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list](resources--cluster--properties--tls_parameters--common_params--validation_params--trusted_ca--trusted_ca_list.md)
- [tls_parameters.common_params.validation_params](resources--cluster--properties--tls_parameters--common_params--validation_params.md)
- [xcsh_cluster](../resources/cluster.md)
