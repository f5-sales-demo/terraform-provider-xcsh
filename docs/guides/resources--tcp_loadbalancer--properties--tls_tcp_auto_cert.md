---
page_title: "tls_tcp_auto_cert"
subcategory: "Load Balancing"
description: "tls_tcp_auto_cert for xcsh_tcp_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1769, "body_sha256": "sha256:9171be2f7b52c354334a3768e144c0d492f98ebde97ddeda6068079391bb59a3", "canonical_id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp_auto_cert", "child_ids": ["xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp_auto_cert:no_mtls", "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp_auto_cert:tls_config", "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp_auto_cert:use_mtls"], "collection_id": "xcsh-docs:resources:tcp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp_auto_cert", "parent_id": "xcsh-docs:resources:tcp_loadbalancer:reference", "path": "docs/guides/resources--tcp_loadbalancer--properties--tls_tcp_auto_cert.md", "provider_name": "tcp_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tls_tcp_auto_cert"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tcp_loadbalancer/properties/tls_tcp_auto_cert/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tls_tcp_auto_cert for xcsh_tcp_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tcp_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_tcp_auto_cert

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md)
- [Property reference](resources--tcp_loadbalancer--reference.md)
- tls_tcp_auto_cert

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Choice for selecting TLS over TCP proxy with automatic certificates.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_mtls",
    "use_mtls")}
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
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

Terraform syntax:

```terraform
tls_tcp_auto_cert {
  # Configure direct properties listed below.
}
```

## Direct properties

- [no_mtls](resources--tcp_loadbalancer--properties--tls_tcp_auto_cert--no_mtls.md): complete subsection reference.

- [tls_config](resources--tcp_loadbalancer--properties--tls_tcp_auto_cert--tls_config.md): complete subsection reference.

- [use_mtls](resources--tcp_loadbalancer--properties--tls_tcp_auto_cert--use_mtls.md): complete subsection reference.

## Next pages

- [tls_tcp_auto_cert.no_mtls](resources--tcp_loadbalancer--properties--tls_tcp_auto_cert--no_mtls.md)
- [tls_tcp_auto_cert.tls_config](resources--tcp_loadbalancer--properties--tls_tcp_auto_cert--tls_config.md)
- [tls_tcp_auto_cert.use_mtls](resources--tcp_loadbalancer--properties--tls_tcp_auto_cert--use_mtls.md)
- [Property reference](resources--tcp_loadbalancer--reference.md)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md)
