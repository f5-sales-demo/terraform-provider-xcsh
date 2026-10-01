---
page_title: "discovery_consul.access_info.http_basic_auth_info"
subcategory: ""
description: "discovery_consul.access_info.http_basic_auth_info for xcsh_discovery."
xcsh_docs: {"aliases": [], "body_bytes": 2544, "body_sha256": "sha256:2d82c0c098407ab629a15fb9ecfceaa4352e9681692ead12ca71d4c26b571e15", "child_ids": ["xcsh-docs:data-sources:discovery:properties:discovery_consul:access_info:http_basic_auth_info:passwd_url"], "collection_id": "xcsh-docs:data-sources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:discovery:properties:discovery_consul:access_info:http_basic_auth_info", "parent_id": "xcsh-docs:data-sources:discovery:properties:discovery_consul:access_info", "path": "documentation/data-sources/discovery/properties/discovery_consul/access_info/http_basic_auth_info/index.md", "provider_name": "discovery", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["discovery_consul", "access_info", "http_basic_auth_info"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/discovery/properties/discovery_consul/access_info/http_basic_auth_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "discovery_consul.access_info.http_basic_auth_info for xcsh_discovery.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# discovery_consul.access_info.http_basic_auth_info

Breadcrumbs:

- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/)
- [discovery_consul](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_consul/)
- [discovery_consul.access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_consul/access_info/)
- discovery_consul.access_info.http_basic_auth_info

<a id="section"></a>

Type: `"single"`. Computed.

Authentication parameters to access Hashicorp Consul.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Direct properties

- [passwd_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_consul/access_info/http_basic_auth_info/passwd_url/): complete subsection reference.

<a id="schema-discovery_consul--access_info--http_basic_auth_info--user_name"></a>

### user_name property

Type: `"string"`. Computed.

User Name. Username in consul.

Upstream description:

Username in consul.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

## Next pages

- [discovery_consul.access_info.http_basic_auth_info.passwd_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_consul/access_info/http_basic_auth_info/passwd_url/)
- [discovery_consul.access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_consul/access_info/)
- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/)
