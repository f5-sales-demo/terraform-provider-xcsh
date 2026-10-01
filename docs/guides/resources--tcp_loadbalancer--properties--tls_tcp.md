---
page_title: "tls_tcp"
subcategory: "Load Balancing"
description: "tls_tcp for xcsh_tcp_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1549, "body_sha256": "sha256:e47183f0eb3000f537075184dd7e2784545f0f3924f44045a7bf00229ca6d22e", "canonical_id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp", "child_ids": ["xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_cert_params", "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_parameters"], "collection_id": "xcsh-docs:resources:tcp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp", "parent_id": "xcsh-docs:resources:tcp_loadbalancer:reference", "path": "docs/guides/resources--tcp_loadbalancer--properties--tls_tcp.md", "provider_name": "tcp_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tls_tcp"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tcp_loadbalancer/properties/tls_tcp/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tls_tcp for xcsh_tcp_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tcp_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_tcp

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md)
- [Property reference](resources--tcp_loadbalancer--reference.md)
- tls_tcp

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Choice for selecting TLS over TCP proxy with bring your own certificates.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("tls_cert_params",
    "tls_parameters")}
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
  "x-ves-oneof-field-tls_certificates_choice": "[\"tls_cert_params\",\"tls_parameters\"]"
}
```

Terraform syntax:

```terraform
tls_tcp {
  # Configure direct properties listed below.
}
```

## Direct properties

- [tls_cert_params](resources--tcp_loadbalancer--properties--tls_tcp--tls_cert_params.md): complete subsection reference.

- [tls_parameters](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters.md): complete subsection reference.

## Next pages

- [tls_tcp.tls_cert_params](resources--tcp_loadbalancer--properties--tls_tcp--tls_cert_params.md)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters.md)
- [Property reference](resources--tcp_loadbalancer--reference.md)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md)
