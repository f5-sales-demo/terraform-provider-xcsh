---
page_title: "tls_tcp.tls_parameters.tls_certificates.private_key"
subcategory: "Load Balancing"
description: "tls_tcp.tls_parameters.tls_certificates.private_key for xcsh_tcp_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2324, "body_sha256": "sha256:7c19f2cc52390c1a7dfedd90b4427840d8f017c0eb1fe4ac0a25d5acfc62f4c8", "canonical_id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates:private_key", "child_ids": ["xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates:private_key:blindfold_secret_info", "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates:private_key:clear_secret_info"], "collection_id": "xcsh-docs:resources:tcp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates:private_key", "parent_id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates", "path": "docs/guides/resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--tls_certificates--private_key.md", "provider_name": "tcp_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tls_tcp", "tls_parameters", "tls_certificates", "private_key"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tcp_loadbalancer/properties/tls_tcp/tls_parameters/tls_certificates/private_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tls_tcp.tls_parameters.tls_certificates.private_key for xcsh_tcp_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tcp_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_tcp.tls_parameters.tls_certificates.private_key

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md)
- [Property reference](resources--tcp_loadbalancer--reference.md)
- [tls_tcp](resources--tcp_loadbalancer--properties--tls_tcp.md)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters.md)
- [tls_tcp.tls_parameters.tls_certificates](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--tls_certificates.md)
- tls_tcp.tls_parameters.tls_certificates.private_key

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
private_key {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--tls_certificates--private_key--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--tls_certificates--private_key--clear_secret_info.md): complete subsection reference.

## Next pages

- [tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--tls_certificates--private_key--blindfold_secret_info.md)
- [tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--tls_certificates--private_key--clear_secret_info.md)
- [tls_tcp.tls_parameters.tls_certificates](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--tls_certificates.md)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md)
