---
page_title: "discovery_consul.access_info"
subcategory: ""
description: "discovery_consul.access_info for xcsh_discovery."
xcsh_docs: {"aliases": [], "body_bytes": 1922, "body_sha256": "sha256:67c7d4d750faa8d51de6785126eac107a08f7e9d12741ee8552b091ffffe9627", "child_ids": ["xcsh-docs:resources:discovery:properties:discovery_consul:access_info:connection_info", "xcsh-docs:resources:discovery:properties:discovery_consul:access_info:http_basic_auth_info"], "collection_id": "xcsh-docs:resources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:discovery:properties:discovery_consul:access_info", "parent_id": "xcsh-docs:resources:discovery:properties:discovery_consul", "path": "documentation/resources/discovery/properties/discovery_consul/access_info/index.md", "provider_name": "discovery", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["discovery_consul", "access_info"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/discovery/properties/discovery_consul/access_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "discovery_consul.access_info for xcsh_discovery.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# discovery_consul.access_info

Breadcrumbs:

- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/)
- [discovery_consul](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/)
- discovery_consul.access_info

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Hashicorp Consul API server information.

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

Terraform syntax:

```terraform
access_info {
  # Configure direct properties listed below.
}
```

## Direct properties

- [connection_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/access_info/connection_info/): complete subsection reference.

- [http_basic_auth_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/access_info/http_basic_auth_info/): complete subsection reference.

## Next pages

- [discovery_consul.access_info.connection_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/access_info/connection_info/)
- [discovery_consul.access_info.http_basic_auth_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/access_info/http_basic_auth_info/)
- [discovery_consul](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/)
- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/)
