---
page_title: "tls_parameters.cert_params.volterra_trusted_ca"
subcategory: ""
description: "tls_parameters.cert_params.volterra_trusted_ca for xcsh_cluster."
xcsh_docs: {"aliases": [], "body_bytes": 1050, "body_sha256": "sha256:667137d078aaa6e94ce722616e64c6d2e415e6f1a7b29a50d7fe2191445123fb", "canonical_id": "xcsh-docs:resources:cluster:properties:tls_parameters:cert_params:volterra_trusted_ca", "child_ids": [], "collection_id": "xcsh-docs:resources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:cluster:properties:tls_parameters:cert_params:volterra_trusted_ca", "parent_id": "xcsh-docs:resources:cluster:properties:tls_parameters:cert_params", "path": "docs/guides/resources--cluster--properties--tls_parameters--cert_params--volterra_trusted_ca.md", "provider_name": "cluster", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tls_parameters", "cert_params", "volterra_trusted_ca"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cluster/properties/tls_parameters/cert_params/volterra_trusted_ca/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tls_parameters.cert_params.volterra_trusted_ca for xcsh_cluster.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# tls_parameters.cert_params.volterra_trusted_ca

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md)
- [Property reference](resources--cluster--reference.md)
- [tls_parameters](resources--cluster--properties--tls_parameters.md)
- [tls_parameters.cert_params](resources--cluster--properties--tls_parameters--cert_params.md)
- tls_parameters.cert_params.volterra_trusted_ca

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for volterra trusted ca.

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
volterra_trusted_ca = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [tls_parameters.cert_params](resources--cluster--properties--tls_parameters--cert_params.md)
- [xcsh_cluster](../resources/cluster.md)
