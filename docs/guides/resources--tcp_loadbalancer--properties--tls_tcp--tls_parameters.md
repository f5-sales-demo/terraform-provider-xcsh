---
page_title: "tls_tcp.tls_parameters"
subcategory: "Load Balancing"
description: "tls_tcp.tls_parameters for xcsh_tcp_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2145, "body_sha256": "sha256:842e3928f21969c9bd329c8d0cd024b7774c3f3fb800090fca4185922daaca8c", "canonical_id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_parameters", "child_ids": ["xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:no_mtls", "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates", "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_config", "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:use_mtls"], "collection_id": "xcsh-docs:resources:tcp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_parameters", "parent_id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp", "path": "docs/guides/resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters.md", "provider_name": "tcp_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tls_tcp", "tls_parameters"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tcp_loadbalancer/properties/tls_tcp/tls_parameters/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tls_tcp.tls_parameters for xcsh_tcp_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tcp_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# tls_tcp.tls_parameters

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md)
- [Property reference](resources--tcp_loadbalancer--reference.md)
- [tls_tcp](resources--tcp_loadbalancer--properties--tls_tcp.md)
- tls_tcp.tls_parameters

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tls parameters.

Upstream description:

Inline TLS parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("tls_certificates"),
  validators.ConflictingObjectAttributes("no_mtls",
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
tls_parameters {
  # Configure direct properties listed below.
}
```

## Direct properties

- [no_mtls](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--no_mtls.md): complete subsection reference.

- [tls_certificates](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--tls_certificates.md): complete subsection reference.

- [tls_config](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--tls_config.md): complete subsection reference.

- [use_mtls](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--use_mtls.md): complete subsection reference.

## Next pages

- [tls_tcp.tls_parameters.no_mtls](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--no_mtls.md)
- [tls_tcp.tls_parameters.tls_certificates](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--tls_certificates.md)
- [tls_tcp.tls_parameters.tls_config](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--tls_config.md)
- [tls_tcp.tls_parameters.use_mtls](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--use_mtls.md)
- [tls_tcp](resources--tcp_loadbalancer--properties--tls_tcp.md)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md)
