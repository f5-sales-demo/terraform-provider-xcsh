---
page_title: "discovery_consul.access_info.connection_info.tls_info.key_url"
subcategory: ""
description: "discovery_consul.access_info.connection_info.tls_info.key_url for xcsh_discovery."
xcsh_docs: {"aliases": [], "body_bytes": 2267, "body_sha256": "sha256:8750ccb8b50c1281261c52164293d5132507badc6e453bcd53e567aa388d8826", "canonical_id": "xcsh-docs:data-sources:discovery:properties:discovery_consul:access_info:connection_info:tls_info:key_url", "child_ids": ["xcsh-docs:data-sources:discovery:properties:discovery_consul:access_info:connection_info:tls_info:key_url:blindfold_secret_info", "xcsh-docs:data-sources:discovery:properties:discovery_consul:access_info:connection_info:tls_info:key_url:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:discovery:properties:discovery_consul:access_info:connection_info:tls_info:key_url", "parent_id": "xcsh-docs:data-sources:discovery:properties:discovery_consul:access_info:connection_info:tls_info", "path": "docs/guides/data-sources--discovery--properties--discovery_consul--access_info--connection_info--tls_info--key_url.md", "provider_name": "discovery", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["discovery_consul", "access_info", "connection_info", "tls_info", "key_url"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/discovery/properties/discovery_consul/access_info/connection_info/tls_info/key_url/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "discovery_consul.access_info.connection_info.tls_info.key_url for xcsh_discovery.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# discovery_consul.access_info.connection_info.tls_info.key_url

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md)
- [Property reference](data-sources--discovery--reference.md)
- [discovery_consul](data-sources--discovery--properties--discovery_consul.md)
- [discovery_consul.access_info](data-sources--discovery--properties--discovery_consul--access_info.md)
- [discovery_consul.access_info.connection_info](data-sources--discovery--properties--discovery_consul--access_info--connection_info.md)
- [discovery_consul.access_info.connection_info.tls_info](data-sources--discovery--properties--discovery_consul--access_info--connection_info--tls_info.md)
- discovery_consul.access_info.connection_info.tls_info.key_url

<a id="section"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

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

## Direct properties

- [blindfold_secret_info](data-sources--discovery--properties--discovery_consul--access_info--connection_info--tls_info--key_url--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](data-sources--discovery--properties--discovery_consul--access_info--connection_info--tls_info--key_url--clear_secret_info.md): complete subsection reference.

## Next pages

- [discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info](data-sources--discovery--properties--discovery_consul--access_info--connection_info--tls_info--key_url--blindfold_secret_info.md)
- [discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info](data-sources--discovery--properties--discovery_consul--access_info--connection_info--tls_info--key_url--clear_secret_info.md)
- [discovery_consul.access_info.connection_info.tls_info](data-sources--discovery--properties--discovery_consul--access_info--connection_info--tls_info.md)
- [xcsh_discovery](../data-sources/discovery.md)
