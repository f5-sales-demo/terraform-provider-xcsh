---
page_title: "https.tls_parameters.tls_config"
subcategory: "Load Balancing"
description: "https.tls_parameters.tls_config for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 3063, "body_sha256": "sha256:ca1fc0d799f0e3fe32ca3d926b1603a82e246386375241c1a8f8448672995f9e", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:https:tls_parameters:tls_config", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:https:tls_parameters:tls_config:custom_security", "xcsh-docs:resources:http_loadbalancer:properties:https:tls_parameters:tls_config:default_security", "xcsh-docs:resources:http_loadbalancer:properties:https:tls_parameters:tls_config:low_security", "xcsh-docs:resources:http_loadbalancer:properties:https:tls_parameters:tls_config:medium_security"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:https:tls_parameters:tls_config", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:https:tls_parameters", "path": "docs/guides/resources--http_loadbalancer--properties--https--tls_parameters--tls_config.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["https", "tls_parameters", "tls_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/https/tls_parameters/tls_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "https.tls_parameters.tls_config for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https.tls_parameters.tls_config

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [https](resources--http_loadbalancer--properties--https.md)
- [https.tls_parameters](resources--http_loadbalancer--properties--https--tls_parameters.md)
- https.tls_parameters.tls_config

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_security",
    "default_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "low_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("default_security",
    "low_security"),
  validators.ConflictingObjectAttributes("default_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("low_security",
    "medium_security")}
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
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

Terraform syntax:

```terraform
tls_config {
  # Configure direct properties listed below.
}
```

## Direct properties

- [custom_security](resources--http_loadbalancer--properties--https--tls_parameters--tls_config--custom_security.md): complete subsection reference.

- [default_security](resources--http_loadbalancer--properties--https--tls_parameters--tls_config--default_security.md): complete subsection reference.

- [low_security](resources--http_loadbalancer--properties--https--tls_parameters--tls_config--low_security.md): complete subsection reference.

- [medium_security](resources--http_loadbalancer--properties--https--tls_parameters--tls_config--medium_security.md): complete subsection reference.

## Next pages

- [https.tls_parameters.tls_config.custom_security](resources--http_loadbalancer--properties--https--tls_parameters--tls_config--custom_security.md)
- [https.tls_parameters.tls_config.default_security](resources--http_loadbalancer--properties--https--tls_parameters--tls_config--default_security.md)
- [https.tls_parameters.tls_config.low_security](resources--http_loadbalancer--properties--https--tls_parameters--tls_config--low_security.md)
- [https.tls_parameters.tls_config.medium_security](resources--http_loadbalancer--properties--https--tls_parameters--tls_config--medium_security.md)
- [https.tls_parameters](resources--http_loadbalancer--properties--https--tls_parameters.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
