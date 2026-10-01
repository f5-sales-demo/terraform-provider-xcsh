---
page_title: "https.tls_parameters"
subcategory: "Load Balancing"
description: "https.tls_parameters for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2223, "body_sha256": "sha256:6cc1e8b16693a7c8e9738f5871eb0c4ee1b8cd25b4216cdd3dc7825714cf8bba", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:https:tls_parameters", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:https:tls_parameters:no_mtls", "xcsh-docs:resources:http_loadbalancer:properties:https:tls_parameters:tls_certificates", "xcsh-docs:resources:http_loadbalancer:properties:https:tls_parameters:tls_config", "xcsh-docs:resources:http_loadbalancer:properties:https:tls_parameters:use_mtls"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:https:tls_parameters", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:https", "path": "docs/guides/resources--http_loadbalancer--properties--https--tls_parameters.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["https", "tls_parameters"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/https/tls_parameters/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "https.tls_parameters for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https.tls_parameters

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [https](resources--http_loadbalancer--properties--https.md)
- https.tls_parameters

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

- [no_mtls](resources--http_loadbalancer--properties--https--tls_parameters--no_mtls.md): complete subsection reference.

- [tls_certificates](resources--http_loadbalancer--properties--https--tls_parameters--tls_certificates.md): complete subsection reference.

- [tls_config](resources--http_loadbalancer--properties--https--tls_parameters--tls_config.md): complete subsection reference.

- [use_mtls](resources--http_loadbalancer--properties--https--tls_parameters--use_mtls.md): complete subsection reference.

## Next pages

- [https.tls_parameters.no_mtls](resources--http_loadbalancer--properties--https--tls_parameters--no_mtls.md)
- [https.tls_parameters.tls_certificates](resources--http_loadbalancer--properties--https--tls_parameters--tls_certificates.md)
- [https.tls_parameters.tls_config](resources--http_loadbalancer--properties--https--tls_parameters--tls_config.md)
- [https.tls_parameters.use_mtls](resources--http_loadbalancer--properties--https--tls_parameters--use_mtls.md)
- [https](resources--http_loadbalancer--properties--https.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
