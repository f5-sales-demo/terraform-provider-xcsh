---
page_title: "custom_network_config.slo_config.static_routes"
subcategory: ""
description: "custom_network_config.slo_config.static_routes for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 1969, "body_sha256": "sha256:48c31972698ad9f7ff491720cb4b15c3eeb58d37c7359807979c0ada6165c3cd", "child_ids": ["xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config:static_routes:static_routes"], "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config:static_routes", "parent_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config", "path": "documentation/resources/voltstack_site/properties/custom_network_config/slo_config/static_routes/index.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["custom_network_config", "slo_config", "static_routes"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/custom_network_config/slo_config/static_routes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_network_config.slo_config.static_routes for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# custom_network_config.slo_config.static_routes

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/)
- [custom_network_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/)
- [custom_network_config.slo_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/slo_config/)
- custom_network_config.slo_config.static_routes

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for static routes.

Upstream description:

List of static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("static_routes")}
```

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
static_routes {
  # Configure direct properties listed below.
}
```

## Direct properties

- [static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/slo_config/static_routes/static_routes/): complete subsection reference.

## Next pages

- [custom_network_config.slo_config.static_routes.static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/slo_config/static_routes/static_routes/)
- [custom_network_config.slo_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/slo_config/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
