---
page_title: "discovery_consul.access_info.connection_info.tls_info.key_url"
subcategory: ""
description: "discovery_consul.access_info.connection_info.tls_info.key_url for xcsh_discovery."
xcsh_docs: {"aliases": [], "body_bytes": 2535, "body_sha256": "sha256:0a87389aa43ab286a308556265aae203fb7463ee47043262e8091e5667b68f41", "canonical_id": "xcsh-docs:resources:discovery:properties:discovery_consul:access_info:connection_info:tls_info:key_url", "child_ids": ["xcsh-docs:resources:discovery:properties:discovery_consul:access_info:connection_info:tls_info:key_url:blindfold_secret_info", "xcsh-docs:resources:discovery:properties:discovery_consul:access_info:connection_info:tls_info:key_url:clear_secret_info"], "collection_id": "xcsh-docs:resources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:discovery:properties:discovery_consul:access_info:connection_info:tls_info:key_url", "parent_id": "xcsh-docs:resources:discovery:properties:discovery_consul:access_info:connection_info:tls_info", "path": "docs/guides/resources--discovery--properties--discovery_consul--access_info--connection_info--tls_info--key_url.md", "provider_name": "discovery", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["discovery_consul", "access_info", "connection_info", "tls_info", "key_url"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/discovery/properties/discovery_consul/access_info/connection_info/tls_info/key_url/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "discovery_consul.access_info.connection_info.tls_info.key_url for xcsh_discovery.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# discovery_consul.access_info.connection_info.tls_info.key_url

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md)
- [Property reference](resources--discovery--reference.md)
- [discovery_consul](resources--discovery--properties--discovery_consul.md)
- [discovery_consul.access_info](resources--discovery--properties--discovery_consul--access_info.md)
- [discovery_consul.access_info.connection_info](resources--discovery--properties--discovery_consul--access_info--connection_info.md)
- [discovery_consul.access_info.connection_info.tls_info](resources--discovery--properties--discovery_consul--access_info--connection_info--tls_info.md)
- discovery_consul.access_info.connection_info.tls_info.key_url

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
key_url {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](resources--discovery--properties--discovery_consul--access_info--connection_info--tls_info--key_url--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](resources--discovery--properties--discovery_consul--access_info--connection_info--tls_info--key_url--clear_secret_info.md): complete subsection reference.

## Next pages

- [discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info](resources--discovery--properties--discovery_consul--access_info--connection_info--tls_info--key_url--blindfold_secret_info.md)
- [discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info](resources--discovery--properties--discovery_consul--access_info--connection_info--tls_info--key_url--clear_secret_info.md)
- [discovery_consul.access_info.connection_info.tls_info](resources--discovery--properties--discovery_consul--access_info--connection_info--tls_info.md)
- [xcsh_discovery](../resources/discovery.md)
