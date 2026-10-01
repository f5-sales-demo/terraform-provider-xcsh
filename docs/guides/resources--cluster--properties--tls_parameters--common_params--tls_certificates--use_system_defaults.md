---
page_title: "tls_parameters.common_params.tls_certificates.use_system_defaults"
subcategory: ""
description: "tls_parameters.common_params.tls_certificates.use_system_defaults for xcsh_cluster."
xcsh_docs: {"aliases": [], "body_bytes": 1364, "body_sha256": "sha256:bca75b3e46f4afce8c8f1e1be36812acb6be4cda36fab60169ee0c6ef9efe27a", "canonical_id": "xcsh-docs:resources:cluster:properties:tls_parameters:common_params:tls_certificates:use_system_defaults", "child_ids": [], "collection_id": "xcsh-docs:resources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:cluster:properties:tls_parameters:common_params:tls_certificates:use_system_defaults", "parent_id": "xcsh-docs:resources:cluster:properties:tls_parameters:common_params:tls_certificates", "path": "docs/guides/resources--cluster--properties--tls_parameters--common_params--tls_certificates--use_system_defaults.md", "provider_name": "cluster", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tls_parameters", "common_params", "tls_certificates", "use_system_defaults"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cluster/properties/tls_parameters/common_params/tls_certificates/use_system_defaults/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tls_parameters.common_params.tls_certificates.use_system_defaults for xcsh_cluster.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_parameters.common_params.tls_certificates.use_system_defaults

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md)
- [Property reference](resources--cluster--reference.md)
- [tls_parameters](resources--cluster--properties--tls_parameters.md)
- [tls_parameters.common_params](resources--cluster--properties--tls_parameters--common_params.md)
- [tls_parameters.common_params.tls_certificates](resources--cluster--properties--tls_parameters--common_params--tls_certificates.md)
- tls_parameters.common_params.tls_certificates.use_system_defaults

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for use system defaults.

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
use_system_defaults = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [tls_parameters.common_params.tls_certificates](resources--cluster--properties--tls_parameters--common_params--tls_certificates.md)
- [xcsh_cluster](../resources/cluster.md)
