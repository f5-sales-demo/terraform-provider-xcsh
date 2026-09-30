---
page_title: "discovery_consul.access_info.http_basic_auth_info.passwd_url"
subcategory: ""
description: "discovery_consul.access_info.http_basic_auth_info.passwd_url for xcsh_discovery."
xcsh_docs: {"aliases": [], "body_bytes": 2001, "body_sha256": "sha256:e5d7d53a1f121800da719bcdd69dc16846226de933fd7bc4d1d7ca936b04d5fc", "canonical_id": "xcsh-docs:data-sources:discovery:properties:discovery_consul:access_info:http_basic_auth_info:passwd_url", "child_ids": ["xcsh-docs:data-sources:discovery:properties:discovery_consul:access_info:http_basic_auth_info:passwd_url:blindfold_secret_info", "xcsh-docs:data-sources:discovery:properties:discovery_consul:access_info:http_basic_auth_info:passwd_url:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:discovery:properties:discovery_consul:access_info:http_basic_auth_info:passwd_url", "parent_id": "xcsh-docs:data-sources:discovery:properties:discovery_consul:access_info:http_basic_auth_info", "path": "docs/guides/data-sources--discovery--properties--discovery_consul--access_info--http_basic_auth_info--passwd_url.md", "provider_name": "discovery", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["discovery_consul", "access_info", "http_basic_auth_info", "passwd_url"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/discovery/properties/discovery_consul/access_info/http_basic_auth_info/passwd_url/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "discovery_consul.access_info.http_basic_auth_info.passwd_url for xcsh_discovery.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# discovery_consul.access_info.http_basic_auth_info.passwd_url

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md)
- [Property reference](data-sources--discovery--reference.md)
- [discovery_consul](data-sources--discovery--properties--discovery_consul.md)
- [discovery_consul.access_info](data-sources--discovery--properties--discovery_consul--access_info.md)
- [discovery_consul.access_info.http_basic_auth_info](data-sources--discovery--properties--discovery_consul--access_info--http_basic_auth_info.md)
- discovery_consul.access_info.http_basic_auth_info.passwd_url

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

- [blindfold_secret_info](data-sources--discovery--properties--discovery_consul--access_info--http_basic_auth_info--passwd_url--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](data-sources--discovery--properties--discovery_consul--access_info--http_basic_auth_info--passwd_url--clear_secret_info.md): complete subsection reference.

## Next pages

- [discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info](data-sources--discovery--properties--discovery_consul--access_info--http_basic_auth_info--passwd_url--blindfold_secret_info.md)
- [discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info](data-sources--discovery--properties--discovery_consul--access_info--http_basic_auth_info--passwd_url--clear_secret_info.md)
- [discovery_consul.access_info.http_basic_auth_info](data-sources--discovery--properties--discovery_consul--access_info--http_basic_auth_info.md)
- [xcsh_discovery](../data-sources/discovery.md)
