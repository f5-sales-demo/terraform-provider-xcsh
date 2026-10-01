---
page_title: "https.tls_parameters.tls_certificates.private_key"
subcategory: "Load Balancing"
description: "https.tls_parameters.tls_certificates.private_key for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2305, "body_sha256": "sha256:e7d8e1a5d3a852aa1fe06d3fe785e0ff5361262767d431f54fbf64bdc6acc3d1", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:https:tls_parameters:tls_certificates:private_key", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:https:tls_parameters:tls_certificates:private_key:blindfold_secret_info", "xcsh-docs:resources:http_loadbalancer:properties:https:tls_parameters:tls_certificates:private_key:clear_secret_info"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:https:tls_parameters:tls_certificates:private_key", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:https:tls_parameters:tls_certificates", "path": "docs/guides/resources--http_loadbalancer--properties--https--tls_parameters--tls_certificates--private_key.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["https", "tls_parameters", "tls_certificates", "private_key"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/https/tls_parameters/tls_certificates/private_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "https.tls_parameters.tls_certificates.private_key for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https.tls_parameters.tls_certificates.private_key

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [https](resources--http_loadbalancer--properties--https.md)
- [https.tls_parameters](resources--http_loadbalancer--properties--https--tls_parameters.md)
- [https.tls_parameters.tls_certificates](resources--http_loadbalancer--properties--https--tls_parameters--tls_certificates.md)
- https.tls_parameters.tls_certificates.private_key

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

- [blindfold_secret_info](resources--http_loadbalancer--properties--https--tls_parameters--tls_certificates--private_key--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](resources--http_loadbalancer--properties--https--tls_parameters--tls_certificates--private_key--clear_secret_info.md): complete subsection reference.

## Next pages

- [https.tls_parameters.tls_certificates.private_key.blindfold_secret_info](resources--http_loadbalancer--properties--https--tls_parameters--tls_certificates--private_key--blindfold_secret_info.md)
- [https.tls_parameters.tls_certificates.private_key.clear_secret_info](resources--http_loadbalancer--properties--https--tls_parameters--tls_certificates--private_key--clear_secret_info.md)
- [https.tls_parameters.tls_certificates](resources--http_loadbalancer--properties--https--tls_parameters--tls_certificates.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
