---
page_title: "virtual_server.source_port"
subcategory: ""
description: "virtual_server.source_port for xcsh_application_profiles."
xcsh_docs: {"aliases": [], "body_bytes": 1913, "body_sha256": "sha256:2fe5fd6940244cdf1812f6e28952c2b626f926a9aa4e5e35c99fd9e282d1487a", "canonical_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:source_port", "child_ids": ["xcsh-docs:data-sources:application_profiles:properties:virtual_server:source_port:source_port_change", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:source_port:source_port_preserve", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:source_port:source_port_preserve_strict"], "collection_id": "xcsh-docs:data-sources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:source_port", "parent_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server", "path": "docs/guides/data-sources--application_profiles--properties--virtual_server--source_port.md", "provider_name": "application_profiles", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["virtual_server", "source_port"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/application_profiles/properties/virtual_server/source_port/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "virtual_server.source_port for xcsh_application_profiles.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["application_profilesCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# virtual_server.source_port

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md)
- [Property reference](data-sources--application_profiles--reference.md)
- [virtual_server](data-sources--application_profiles--properties--virtual_server.md)
- virtual_server.source_port

<a id="section"></a>

Type: `"single"`. Computed.

Specifies whether the system preserves the source port of the connection. The default is Preserve.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-source_port_choice": "[\"source_port_change\",\"source_port_preserve\",\"source_port_preserve_strict\"]"
}
```

## Direct properties

- [source_port_change](data-sources--application_profiles--properties--virtual_server--source_port--source_port_change.md): complete subsection reference.

- [source_port_preserve](data-sources--application_profiles--properties--virtual_server--source_port--source_port_preserve.md): complete subsection reference.

- [source_port_preserve_strict](data-sources--application_profiles--properties--virtual_server--source_port--source_port_preserve_strict.md): complete subsection reference.

## Next pages

- [virtual_server.source_port.source_port_change](data-sources--application_profiles--properties--virtual_server--source_port--source_port_change.md)
- [virtual_server.source_port.source_port_preserve](data-sources--application_profiles--properties--virtual_server--source_port--source_port_preserve.md)
- [virtual_server.source_port.source_port_preserve_strict](data-sources--application_profiles--properties--virtual_server--source_port--source_port_preserve_strict.md)
- [virtual_server](data-sources--application_profiles--properties--virtual_server.md)
- [xcsh_application_profiles](../data-sources/application_profiles.md)
