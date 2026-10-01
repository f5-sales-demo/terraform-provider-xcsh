---
page_title: "tls_parameters.cert_params.skip_server_verification"
subcategory: ""
description: "tls_parameters.cert_params.skip_server_verification for xcsh_cluster."
xcsh_docs: {"aliases": [], "body_bytes": 1134, "body_sha256": "sha256:f9253957a9f13e1ad90b8fd5e91a1eab4244d2a9da154858c9aef8ad1c105caa", "canonical_id": "xcsh-docs:resources:cluster:properties:tls_parameters:cert_params:skip_server_verification", "child_ids": [], "collection_id": "xcsh-docs:resources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:cluster:properties:tls_parameters:cert_params:skip_server_verification", "parent_id": "xcsh-docs:resources:cluster:properties:tls_parameters:cert_params", "path": "docs/guides/resources--cluster--properties--tls_parameters--cert_params--skip_server_verification.md", "provider_name": "cluster", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tls_parameters", "cert_params", "skip_server_verification"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cluster/properties/tls_parameters/cert_params/skip_server_verification/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tls_parameters.cert_params.skip_server_verification for xcsh_cluster.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_parameters.cert_params.skip_server_verification

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md)
- [Property reference](resources--cluster--reference.md)
- [tls_parameters](resources--cluster--properties--tls_parameters.md)
- [tls_parameters.cert_params](resources--cluster--properties--tls_parameters--cert_params.md)
- tls_parameters.cert_params.skip_server_verification

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
skip_server_verification = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [tls_parameters.cert_params](resources--cluster--properties--tls_parameters--cert_params.md)
- [xcsh_cluster](../resources/cluster.md)
