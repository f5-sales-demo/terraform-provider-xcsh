---
page_title: "proxy_config.https.tls_parameters.tls_certificates.private_key"
subcategory: ""
description: "proxy_config.https.tls_parameters.tls_certificates.private_key for xcsh_bigip_http_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 2484, "body_sha256": "sha256:05051ec2716ee094b49e67c38fb6f5006e24e22aad89b360b84ef17e935cd4f8", "canonical_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_parameters:tls_certificates:private_key", "child_ids": ["xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_parameters:tls_certificates:private_key:blindfold_secret_info", "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_parameters:tls_certificates:private_key:clear_secret_info"], "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_parameters:tls_certificates:private_key", "parent_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_parameters:tls_certificates", "path": "docs/guides/resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--tls_certificates--private_key.md", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["proxy_config", "https", "tls_parameters", "tls_certificates", "private_key"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/properties/proxy_config/https/tls_parameters/tls_certificates/private_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "proxy_config.https.tls_parameters.tls_certificates.private_key for xcsh_bigip_http_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# proxy_config.https.tls_parameters.tls_certificates.private_key

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md)
- [Property reference](resources--bigip_http_proxy--reference.md)
- [proxy_config](resources--bigip_http_proxy--properties--proxy_config.md)
- [proxy_config.https](resources--bigip_http_proxy--properties--proxy_config--https.md)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters.md)
- [proxy_config.https.tls_parameters.tls_certificates](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--tls_certificates.md)
- proxy_config.https.tls_parameters.tls_certificates.private_key

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

- [blindfold_secret_info](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--tls_certificates--private_key--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--tls_certificates--private_key--clear_secret_info.md): complete subsection reference.

## Next pages

- [proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--tls_certificates--private_key--blindfold_secret_info.md)
- [proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--tls_certificates--private_key--clear_secret_info.md)
- [proxy_config.https.tls_parameters.tls_certificates](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--tls_certificates.md)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md)
